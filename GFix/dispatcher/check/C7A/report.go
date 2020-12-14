package C7A

import (
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/output"
)

type bug_report struct {
	ch   *channel
	send *ch_send
	recv *ch_receive
}

func report(bug bug_report) {

	for _, reported := range C7A_reported {
		if reported.ch == bug.ch { // only report one bug for one channel
			return
		}
	}

	C7A_reported = append(C7A_reported, bug)
	if bug.send == nil && bug.recv == nil {
		return
	}
	global.Bug_index_mu.Lock()
	global.Bug_index++
	/*fmt.Print("----------Bug[")
	fmt.Print(global.Bug_index)*/
	global.Bug_index_mu.Unlock()
	if bug.send != nil { //A local channel's send is definitely executed, but receive is not.
		//print file, make line no, send line no, recv line no
		/*if isGL1_new(bug) {
			fmt.Printf("1, path=%s, make_lineno=%d, send_lineno=%d\n", // recv_lineno=%d\n",
				getFileName(global.Prog.Fset, bug.ch.make_inst),
				getLineNo(global.Prog.Fset, bug.ch.make_inst),
				getLineNo(global.Prog.Fset, bug.send.inst),
			)
		} else {
			fmt.Printf("0, path=%s, make_lineno=%d, send_lineno=%d\n", // recv_lineno=%d\n",
				getFileName(global.Prog.Fset, bug.ch.make_inst),
				getLineNo(global.Prog.Fset, bug.ch.make_inst),
				getLineNo(global.Prog.Fset, bug.send.inst),
			)
		}*/
		/*
			lineno := getGL2PatchLineNo(bug)
			if lineno != -1 {
				fmt.Printf("2, path=%s, make_lineno=%d, defer_insert_lineno=%d, defer_remove_lineno=\n",
					getFileName(global.Prog.Fset, bug.ch.make_inst),
					getLineNo(global.Prog.Fset, bug.ch.make_inst),
					lineno,
				)
			} else {
			}
		*/
	} else {
		linenoIns, linenosRemove := getGL2PatchLineNo(bug)
		fmt.Printf("2, path=%s, make_lineno=%d, defer_insert_lineno=%d, defer_remove_lineno=", //%d
			getFileName(global.Prog.Fset, bug.ch.make_inst),
			getLineNo(global.Prog.Fset, bug.ch.make_inst),
			linenoIns,
		)
		for _, x := range linenosRemove {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
		/*	fmt.Print("]----------\n\tType: Goroutine Leak \tReason: A local channel's receive is definitely executed, but send is not.\n")
			fmt.Print("\tLocation of make of channel:\n")
			output.Print_inst_and_location(bug.ch.make_inst)
			fmt.Print("\tLocation of receive of channel:\n")
			output.Print_inst_and_location(bug.recv.inst)*/
	}
	return
}

func report_direct_usage(bug bug_report) {

	for _, reported := range C7A_reported {
		if reported.ch == bug.ch { // only report one bug for one channel
			return
		}
	}

	C7A_reported = append(C7A_reported, bug)
	if bug.send == nil && bug.recv == nil {
		return
	}
	global.Bug_index_mu.Lock()
	global.Bug_index++
	fmt.Print("----------Bug[")
	fmt.Print(global.Bug_index)
	global.Bug_index_mu.Unlock()
	if bug.send != nil {
		fmt.Print("]----------\n\tType: Goroutine Leak \tReason: A local unbuffered channel's send is directly used after creation.\n")
		fmt.Print("\tLocation of make of channel:\n")
		output.Print_inst_and_location(bug.ch.make_inst)
		fmt.Print("\tLocation of send of channel:\n")
		output.Print_inst_and_location(bug.send.inst)
	} else {
		fmt.Print("]----------\n\tType: Goroutine Leak \tReason: A local unbuffered channel's recv is directly used after creation.\n")
		fmt.Print("\tLocation of make of channel:\n")
		output.Print_inst_and_location(bug.ch.make_inst)
		fmt.Print("\tLocation of receive of channel:\n")
		output.Print_inst_and_location(bug.recv.inst)
	}
	return
}
