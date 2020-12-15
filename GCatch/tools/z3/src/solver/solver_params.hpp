// Automatically generated file
#ifndef __SOLVER_PARAMS_HPP_
#define __SOLVER_PARAMS_HPP_
#include "util/params.h"
#include "util/gparams.h"
struct solver_params {
  params_ref const & p;
  params_ref g;
  solver_params(params_ref const & _p = params_ref::get_empty()):
     p(_p), g(gparams::get_module("solver")) {}
  static void collect_param_descrs(param_descrs & d) {
    d.insert("enforce_model_conversion", CPK_BOOL, "apply model transformation on new assertions", "false","solver");
    d.insert("smtlib2_log", CPK_SYMBOL, "file to save solver interaction", "","solver");
    d.insert("cancel_backup_file", CPK_SYMBOL, "file to save partial search state if search is canceled", "","solver");
    d.insert("timeout", CPK_UINT, "timeout on the solver object; overwrites a global timeout", "4294967295","solver");
  }
  /*
     REG_MODULE_PARAMS('solver', 'solver_params::collect_param_descrs')
     REG_MODULE_DESCRIPTION('solver', 'solver parameters')
  */
  bool enforce_model_conversion() const { return p.get_bool("enforce_model_conversion", g, false); }
  symbol smtlib2_log() const { return p.get_sym("smtlib2_log", g, symbol("")); }
  symbol cancel_backup_file() const { return p.get_sym("cancel_backup_file", g, symbol("")); }
  unsigned timeout() const { return p.get_uint("timeout", g, 4294967295u); }
};
#endif
