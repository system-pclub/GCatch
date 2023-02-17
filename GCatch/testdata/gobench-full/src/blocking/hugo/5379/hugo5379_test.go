package hugo5379

import (
	"context"

	"log"
	"sync"
	"testing"
	"time"
)

type shortcodeHandler struct {
	p                      *PageWithoutContent
	contentShortcodes      map[int]func() error
	contentShortcodesDelta map[int]func() error
	init                   sync.Once
}

func (s *shortcodeHandler) executeShortcodesForDelta(p *PageWithoutContent) error {
	for k, _ := range s.contentShortcodesDelta {
		render := s.contentShortcodesDelta[k]
		if err := render(); err != nil {
			continue
		}
	}
	return nil
}

func (s *shortcodeHandler) updateDelta() {
	s.init.Do(func() {
		s.contentShortcodes = createShortcodeRenderers(s.p.withoutContent())
	})

	delta := make(map[int]func() error)

	for k, v := range s.contentShortcodes {
		if _, ok := delta[k]; !ok {
			delta[k] = v
		}
	}

	s.contentShortcodesDelta = delta
}

type Page struct {
	*pageInit
	*pageContentInit
	pageWithoutContent *PageWithoutContent
	contentInit        sync.Once
	contentInitMu      sync.Mutex
	shortcodeState     *shortcodeHandler
}

func (p *Page) WordCount() {
	p.initContentPlainAndMeta()
}

func (p *Page) initContentPlainAndMeta() {
	p.initContent()
	p.initPlain(true)
}

func (p *Page) initPlain(lock bool) {
	p.plainInit.Do(func() {
		if lock {
			p.contentInitMu.Lock() /// Double locking here.
			defer p.contentInitMu.Unlock()
		}
	})
}

func (p *Page) withoutContent() *PageWithoutContent {
	p.pageInit.withoutContentInit.Do(func() {
		p.pageWithoutContent = &PageWithoutContent{Page: p}
	})
	return p.pageWithoutContent
}

func (p *Page) prepareForRender() error {
	var err error
	if err = handleShortcodes(p.withoutContent()); err != nil {
		return err
	}
	return nil
}

func (p *Page) setContentInit() {
	p.shortcodeState.updateDelta()
}

func (p *Page) initContent() {
	p.contentInit.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		c := make(chan error, 1)

		go func() {
			var err error
			p.contentInitMu.Lock() // first lock here
			defer p.contentInitMu.Unlock()

			err = p.prepareForRender()
			if err != nil {
				c <- err
				return
			}
			c <- err
		}()

		select {
		case <-ctx.Done():
		case <-c:
		}
	})
}

type PageWithoutContent struct {
	*Page
}

type pageInit struct {
	withoutContentInit sync.Once
}

type pageContentInit struct {
	contentInit sync.Once
	plainInit   sync.Once
}

type HugoSites struct {
	Sites []*Site
}

func (h *HugoSites) render() {
	for _, s := range h.Sites {
		for _, s2 := range h.Sites {
			s2.preparePagesForRender()
		}
		s.renderPages()
	}
}

func (h *HugoSites) Build() {
	h.render()
}

type Pages []*Page

type PageCollections struct {
	Pages Pages
}

type Site struct {
	*PageCollections
}

func (s *Site) preparePagesForRender() {
	for _, p := range s.Pages {
		p.setContentInit()
	}
}

func (s *Site) renderForLayouts() {
	/// Omit reflections
	for _, p := range s.Pages {
		p.WordCount()
	}
}

func (s *Site) renderAndWritePage() {
	s.renderForLayouts()
}

func (s *Site) renderPages() {
	numWorkers := 2
	wg := &sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go pageRenderer(s, wg)
	}

	wg.Wait()
}

type sitesBuilder struct {
	H *HugoSites
}

func (s *sitesBuilder) Build() *sitesBuilder {
	return s.build()
}

func (s *sitesBuilder) build() *sitesBuilder {
	s.H.Build()
	return s
}

func (s *sitesBuilder) CreateSitesE() error {
	sites, err := NewHugoSites()
	if err != nil {
		return err
	}
	s.H = sites
	return nil
}

func (s *sitesBuilder) CreateSites() *sitesBuilder {
	if err := s.CreateSitesE(); err != nil {
		log.Fatalf("Failed to create sites: %s", err)
	}
	return s
}

func newHugoSites(sites ...*Site) (*HugoSites, error) {
	h := &HugoSites{Sites: sites}
	return h, nil
}

func newSite() *Site {
	c := &PageCollections{}
	s := &Site{
		PageCollections: c,
	}
	return s
}

func createSitesFromConfig() []*Site {
	var (
		sites []*Site
	)

	var s *Site
	s = newSite()
	sites = append(sites, s)
	return sites
}

func NewHugoSites() (*HugoSites, error) {
	sites := createSitesFromConfig()
	return newHugoSites(sites...)
}

func prepareShortcodeForPage(p *PageWithoutContent) map[int]func() error {
	m := make(map[int]func() error)
	m[0] = func() error {
		return renderShortcode(p)
	}
	return m
}

func renderShortcode(p *PageWithoutContent) error {
	return renderShortcodeWithPage(p)
}

func renderShortcodeWithPage(p *PageWithoutContent) error {
	/// Omit reflections
	p.WordCount()
	return nil
}

func createShortcodeRenderers(p *PageWithoutContent) map[int]func() error {
	return prepareShortcodeForPage(p)
}

func newShortcodeHandler(p *Page) *shortcodeHandler {
	return &shortcodeHandler{
		p:                      p.withoutContent(),
		contentShortcodes:      make(map[int]func() error),
		contentShortcodesDelta: make(map[int]func() error),
	}
}

func handleShortcodes(p *PageWithoutContent) error {
	return p.shortcodeState.executeShortcodesForDelta(p)
}

func pageRenderer(s *Site, wg *sync.WaitGroup) {
	defer wg.Done()
	s.renderAndWritePage()
}
func TestHugo5379(t *testing.T) {
	b := &sitesBuilder{}
	s := b.CreateSites()
	for _, site := range s.H.Sites {
		p := &Page{
			pageInit:           &pageInit{},
			pageContentInit:    &pageContentInit{},
			pageWithoutContent: &PageWithoutContent{},
			contentInit:        sync.Once{},
			contentInitMu:      sync.Mutex{},
			shortcodeState:     nil,
		}
		p.shortcodeState = newShortcodeHandler(p)
		site.Pages = append(site.Pages, p)
	}
	s.Build()
}
