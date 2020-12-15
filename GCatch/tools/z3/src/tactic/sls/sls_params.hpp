// Automatically generated file
#ifndef __SLS_PARAMS_HPP_
#define __SLS_PARAMS_HPP_
#include "util/params.h"
#include "util/gparams.h"
struct sls_params {
  params_ref const & p;
  params_ref g;
  sls_params(params_ref const & _p = params_ref::get_empty()):
     p(_p), g(gparams::get_module("sls")) {}
  static void collect_param_descrs(param_descrs & d) {
    d.insert("max_memory", CPK_UINT, "maximum amount of memory in megabytes", "4294967295","sls");
    d.insert("max_restarts", CPK_UINT, "maximum number of restarts", "4294967295","sls");
    d.insert("walksat", CPK_BOOL, "use walksat assertion selection (instead of gsat)", "true","sls");
    d.insert("walksat_ucb", CPK_BOOL, "use bandit heuristic for walksat assertion selection (instead of random)", "true","sls");
    d.insert("walksat_ucb_constant", CPK_DOUBLE, "the ucb constant c in the term score + c * f(touched)", "20.0","sls");
    d.insert("walksat_ucb_init", CPK_BOOL, "initialize total ucb touched to formula size", "false","sls");
    d.insert("walksat_ucb_forget", CPK_DOUBLE, "scale touched by this factor every base restart interval", "1.0","sls");
    d.insert("walksat_ucb_noise", CPK_DOUBLE, "add noise 0 <= 256 * ucb_noise to ucb score for assertion selection", "0.0002","sls");
    d.insert("walksat_repick", CPK_BOOL, "repick assertion if randomizing in local minima", "true","sls");
    d.insert("scale_unsat", CPK_DOUBLE, "scale score of unsat expressions by this factor", "0.5","sls");
    d.insert("paws_init", CPK_UINT, "initial/minimum assertion weights", "40","sls");
    d.insert("paws_sp", CPK_UINT, "smooth assertion weights with probability paws_sp / 1024", "52","sls");
    d.insert("wp", CPK_UINT, "random walk with probability wp / 1024", "100","sls");
    d.insert("vns_mc", CPK_UINT, "in local minima, try Monte Carlo sampling vns_mc many 2-bit-flips per bit", "0","sls");
    d.insert("vns_repick", CPK_BOOL, "in local minima, try picking a different assertion (only for walksat)", "false","sls");
    d.insert("restart_base", CPK_UINT, "base restart interval given by moves per run", "100","sls");
    d.insert("restart_init", CPK_BOOL, "initialize to 0 or random value (= 1) after restart", "false","sls");
    d.insert("early_prune", CPK_BOOL, "use early pruning for score prediction", "true","sls");
    d.insert("random_offset", CPK_BOOL, "use random offset for candidate evaluation", "true","sls");
    d.insert("rescore", CPK_BOOL, "rescore/normalize top-level score every base restart interval", "true","sls");
    d.insert("track_unsat", CPK_BOOL, "keep a list of unsat assertions as done in SAT - currently disabled internally", "false","sls");
    d.insert("random_seed", CPK_UINT, "random seed", "0","sls");
  }
  /*
     REG_MODULE_PARAMS('sls', 'sls_params::collect_param_descrs')
     REG_MODULE_DESCRIPTION('sls', 'Experimental Stochastic Local Search Solver (for QFBV only).')
  */
  unsigned max_memory() const { return p.get_uint("max_memory", g, 4294967295u); }
  unsigned max_restarts() const { return p.get_uint("max_restarts", g, 4294967295u); }
  bool walksat() const { return p.get_bool("walksat", g, true); }
  bool walksat_ucb() const { return p.get_bool("walksat_ucb", g, true); }
  double walksat_ucb_constant() const { return p.get_double("walksat_ucb_constant", g, 20.0); }
  bool walksat_ucb_init() const { return p.get_bool("walksat_ucb_init", g, false); }
  double walksat_ucb_forget() const { return p.get_double("walksat_ucb_forget", g, 1.0); }
  double walksat_ucb_noise() const { return p.get_double("walksat_ucb_noise", g, 0.0002); }
  bool walksat_repick() const { return p.get_bool("walksat_repick", g, true); }
  double scale_unsat() const { return p.get_double("scale_unsat", g, 0.5); }
  unsigned paws_init() const { return p.get_uint("paws_init", g, 40u); }
  unsigned paws_sp() const { return p.get_uint("paws_sp", g, 52u); }
  unsigned wp() const { return p.get_uint("wp", g, 100u); }
  unsigned vns_mc() const { return p.get_uint("vns_mc", g, 0u); }
  bool vns_repick() const { return p.get_bool("vns_repick", g, false); }
  unsigned restart_base() const { return p.get_uint("restart_base", g, 100u); }
  bool restart_init() const { return p.get_bool("restart_init", g, false); }
  bool early_prune() const { return p.get_bool("early_prune", g, true); }
  bool random_offset() const { return p.get_bool("random_offset", g, true); }
  bool rescore() const { return p.get_bool("rescore", g, true); }
  bool track_unsat() const { return p.get_bool("track_unsat", g, false); }
  unsigned random_seed() const { return p.get_uint("random_seed", g, 0u); }
};
#endif
