// Automatically generated file
#ifndef __NLSAT_PARAMS_HPP_
#define __NLSAT_PARAMS_HPP_
#include "util/params.h"
#include "util/gparams.h"
struct nlsat_params {
  params_ref const & p;
  params_ref g;
  nlsat_params(params_ref const & _p = params_ref::get_empty()):
     p(_p), g(gparams::get_module("nlsat")) {}
  static void collect_param_descrs(param_descrs & d) {
    d.insert("max_memory", CPK_UINT, "maximum amount of memory in megabytes", "4294967295","nlsat");
    d.insert("lazy", CPK_UINT, "how lazy the solver is.", "0","nlsat");
    d.insert("reorder", CPK_BOOL, "reorder variables.", "true","nlsat");
    d.insert("log_lemmas", CPK_BOOL, "display lemmas as self-contained SMT formulas", "false","nlsat");
    d.insert("check_lemmas", CPK_BOOL, "check lemmas on the fly using an independent nlsat solver", "false","nlsat");
    d.insert("simplify_conflicts", CPK_BOOL, "simplify conflicts using equalities before resolving them in nlsat solver.", "true","nlsat");
    d.insert("minimize_conflicts", CPK_BOOL, "minimize conflicts", "false","nlsat");
    d.insert("randomize", CPK_BOOL, "randomize selection of a witness in nlsat.", "true","nlsat");
    d.insert("max_conflicts", CPK_UINT, "maximum number of conflicts.", "4294967295","nlsat");
    d.insert("shuffle_vars", CPK_BOOL, "use a random variable order.", "false","nlsat");
    d.insert("inline_vars", CPK_BOOL, "inline variables that can be isolated from equations (not supported in incremental mode)", "false","nlsat");
    d.insert("seed", CPK_UINT, "random seed.", "0","nlsat");
    d.insert("factor", CPK_BOOL, "factor polynomials produced during conflict resolution.", "true","nlsat");
  }
  /*
     REG_MODULE_PARAMS('nlsat', 'nlsat_params::collect_param_descrs')
     REG_MODULE_DESCRIPTION('nlsat', 'nonlinear solver')
  */
  unsigned max_memory() const { return p.get_uint("max_memory", g, 4294967295u); }
  unsigned lazy() const { return p.get_uint("lazy", g, 0u); }
  bool reorder() const { return p.get_bool("reorder", g, true); }
  bool log_lemmas() const { return p.get_bool("log_lemmas", g, false); }
  bool check_lemmas() const { return p.get_bool("check_lemmas", g, false); }
  bool simplify_conflicts() const { return p.get_bool("simplify_conflicts", g, true); }
  bool minimize_conflicts() const { return p.get_bool("minimize_conflicts", g, false); }
  bool randomize() const { return p.get_bool("randomize", g, true); }
  unsigned max_conflicts() const { return p.get_uint("max_conflicts", g, 4294967295u); }
  bool shuffle_vars() const { return p.get_bool("shuffle_vars", g, false); }
  bool inline_vars() const { return p.get_bool("inline_vars", g, false); }
  unsigned seed() const { return p.get_uint("seed", g, 0u); }
  bool factor() const { return p.get_bool("factor", g, true); }
};
#endif
