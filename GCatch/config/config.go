package config

const MAX_SECOND = 1800000
const POINTER_CONSIDER_REFLECTION = false
const MAX_LCA_LAYER = 5 // The maximum caller-callee layers when updating dependency map and finding LCA (Lowest Common Ancester)
const MAX_INST_IN_SYNCGRAPH = 10000
const DISABLE_OPTIMIZATION_CALLEES = false // If set to false, we won't enter every callee while building syncgraph

// flag constants
const ConstPrintDeferMap = "print-defer-map"
