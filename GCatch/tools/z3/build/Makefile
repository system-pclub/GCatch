# Automatically generated file.
include config.mk
all: shell api_dll
	@echo Z3 was successfully built.
	@echo "Z3Py scripts can already be executed in the 'build/python' directory."
	@echo "Z3Py scripts stored in arbitrary directories can be executed if the 'build/python' directory is added to the PYTHONPATH environment variable and the 'build' directory is added to the LD_LIBRARY_PATH environment variable."
	@echo Use the following command to install Z3 at prefix $(PREFIX).
	@echo "    sudo make install"

examples: _ex_cpp_example _ex_z3_tptp _ex_c_example _ex_maxsat _ex_py_example
	@echo Z3 examples were successfully built.
includes_0 = -I../src
util/mutex.h.node: ../src/util/mutex.h
	@echo done > util/mutex.h.node
util/error_codes.h.node: ../src/util/error_codes.h
	@echo done > util/error_codes.h.node
util/warning.h.node: ../src/util/warning.h
	@echo done > util/warning.h.node
util/debug.h.node: ../src/util/debug.h util/error_codes.h.node util/warning.h.node
	@echo done > util/debug.h.node
util/z3_exception.h.node: ../src/util/z3_exception.h
	@echo done > util/z3_exception.h.node
util/memory_manager.h.node: ../src/util/memory_manager.h util/z3_exception.h.node
	@echo done > util/memory_manager.h.node
util/util.h.node: ../src/util/util.h util/debug.h.node util/memory_manager.h.node
	@echo done > util/util.h.node
util/machine.h.node: ../src/util/machine.h
	@echo done > util/machine.h.node
util/small_object_allocator.h.node: ../src/util/small_object_allocator.h util/machine.h.node util/debug.h.node
	@echo done > util/small_object_allocator.h.node
util/trace.h.node: ../src/util/trace.h
	@echo done > util/trace.h.node
util/scoped_numeral.h.node: ../src/util/scoped_numeral.h
	@echo done > util/scoped_numeral.h.node
util/hash.h.node: ../src/util/hash.h util/util.h.node
	@echo done > util/hash.h.node
util/vector.h.node: ../src/util/vector.h util/debug.h.node util/memory_manager.h.node util/hash.h.node util/z3_exception.h.node
	@echo done > util/vector.h.node
util/scoped_numeral_vector.h.node: ../src/util/scoped_numeral_vector.h util/vector.h.node
	@echo done > util/scoped_numeral_vector.h.node
util/buffer.h.node: ../src/util/buffer.h util/memory_manager.h.node
	@echo done > util/buffer.h.node
util/mpn.h.node: ../src/util/mpn.h util/util.h.node util/buffer.h.node
	@echo done > util/mpn.h.node
util/mpz.h.node: ../src/util/mpz.h util/mutex.h.node util/util.h.node util/small_object_allocator.h.node util/trace.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node util/mpn.h.node
	@echo done > util/mpz.h.node
util/mpq.h.node: ../src/util/mpq.h util/mpz.h.node util/trace.h.node
	@echo done > util/mpq.h.node
util/rational.h.node: ../src/util/rational.h util/mpq.h.node
	@echo done > util/rational.h.node
util/s_integer.h.node: ../src/util/s_integer.h util/rational.h.node
	@echo done > util/s_integer.h.node
util/s_integer$(OBJ_EXT): ../src/util/s_integer.cpp util/s_integer.h.node
	@echo src/util/s_integer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/s_integer$(OBJ_EXT) ../src/util/s_integer.cpp
util/rlimit.h.node: ../src/util/rlimit.h util/vector.h.node
	@echo done > util/rlimit.h.node
util/common_msgs.h.node: ../src/util/common_msgs.h
	@echo done > util/common_msgs.h.node
util/rlimit$(OBJ_EXT): ../src/util/rlimit.cpp util/rlimit.h.node util/common_msgs.h.node util/mutex.h.node
	@echo src/util/rlimit.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/rlimit$(OBJ_EXT) ../src/util/rlimit.cpp
util/event_handler.h.node: ../src/util/event_handler.h
	@echo done > util/event_handler.h.node
util/scoped_timer.h.node: ../src/util/scoped_timer.h util/event_handler.h.node
	@echo done > util/scoped_timer.h.node
util/memory_manager$(OBJ_EXT): ../src/util/memory_manager.cpp util/mutex.h.node util/trace.h.node util/memory_manager.h.node util/error_codes.h.node util/debug.h.node util/scoped_timer.h.node
	@echo src/util/memory_manager.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/memory_manager$(OBJ_EXT) ../src/util/memory_manager.cpp
util/mpbq.h.node: ../src/util/mpbq.h util/mpq.h.node util/rational.h.node util/vector.h.node
	@echo done > util/mpbq.h.node
util/mpbq$(OBJ_EXT): ../src/util/mpbq.cpp util/mpbq.h.node
	@echo src/util/mpbq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpbq$(OBJ_EXT) ../src/util/mpbq.cpp
util/bit_vector.h.node: ../src/util/bit_vector.h util/debug.h.node util/vector.h.node util/memory_manager.h.node
	@echo done > util/bit_vector.h.node
util/bit_vector$(OBJ_EXT): ../src/util/bit_vector.cpp util/bit_vector.h.node util/trace.h.node
	@echo src/util/bit_vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/bit_vector$(OBJ_EXT) ../src/util/bit_vector.cpp
util/mpq_inf.h.node: ../src/util/mpq_inf.h util/mpq.h.node util/hash.h.node
	@echo done > util/mpq_inf.h.node
util/mpq_inf$(OBJ_EXT): ../src/util/mpq_inf.cpp util/mpq_inf.h.node
	@echo src/util/mpq_inf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpq_inf$(OBJ_EXT) ../src/util/mpq_inf.cpp
util/inf_int_rational.h.node: ../src/util/inf_int_rational.h util/debug.h.node util/vector.h.node util/rational.h.node
	@echo done > util/inf_int_rational.h.node
util/inf_int_rational$(OBJ_EXT): ../src/util/inf_int_rational.cpp util/inf_int_rational.h.node
	@echo src/util/inf_int_rational.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/inf_int_rational$(OBJ_EXT) ../src/util/inf_int_rational.cpp
util/env_params.h.node: ../src/util/env_params.h
	@echo done > util/env_params.h.node
util/tptr.h.node: ../src/util/tptr.h util/machine.h.node
	@echo done > util/tptr.h.node
util/string_buffer.h.node: ../src/util/string_buffer.h util/util.h.node util/memory_manager.h.node
	@echo done > util/string_buffer.h.node
util/symbol.h.node: ../src/util/symbol.h util/util.h.node util/tptr.h.node util/string_buffer.h.node
	@echo done > util/symbol.h.node
util/cmd_context_types.h.node: ../src/util/cmd_context_types.h util/symbol.h.node util/z3_exception.h.node
	@echo done > util/cmd_context_types.h.node
util/params.h.node: ../src/util/params.h util/cmd_context_types.h.node util/vector.h.node
	@echo done > util/params.h.node
util/gparams.h.node: ../src/util/gparams.h util/params.h.node
	@echo done > util/gparams.h.node
util/env_params$(OBJ_EXT): ../src/util/env_params.cpp util/env_params.h.node util/params.h.node util/gparams.h.node util/util.h.node util/memory_manager.h.node
	@echo src/util/env_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/env_params$(OBJ_EXT) ../src/util/env_params.cpp
util/hashtable.h.node: ../src/util/hashtable.h util/debug.h.node util/util.h.node util/memory_manager.h.node util/hash.h.node util/vector.h.node
	@echo done > util/hashtable.h.node
util/str_hashtable.h.node: ../src/util/str_hashtable.h util/hashtable.h.node util/hash.h.node
	@echo done > util/str_hashtable.h.node
util/z3_version.h.node: ../src/util/z3_version.h
	@echo done > util/z3_version.h.node
util/debug$(OBJ_EXT): ../src/util/debug.cpp util/mutex.h.node util/str_hashtable.h.node util/z3_exception.h.node util/z3_version.h.node
	@echo src/util/debug.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/debug$(OBJ_EXT) ../src/util/debug.cpp
util/mpq$(OBJ_EXT): ../src/util/mpq.cpp util/mpq.h.node util/warning.h.node util/z3_exception.h.node
	@echo src/util/mpq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpq$(OBJ_EXT) ../src/util/mpq.cpp
util/scoped_ctrl_c.h.node: ../src/util/scoped_ctrl_c.h util/event_handler.h.node util/util.h.node
	@echo done > util/scoped_ctrl_c.h.node
util/scoped_ctrl_c$(OBJ_EXT): ../src/util/scoped_ctrl_c.cpp util/scoped_ctrl_c.h.node
	@echo src/util/scoped_ctrl_c.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/scoped_ctrl_c$(OBJ_EXT) ../src/util/scoped_ctrl_c.cpp
util/page.h.node: ../src/util/page.h util/memory_manager.h.node
	@echo done > util/page.h.node
util/stack.h.node: ../src/util/stack.h util/page.h.node util/debug.h.node
	@echo done > util/stack.h.node
util/stack$(OBJ_EXT): ../src/util/stack.cpp util/stack.h.node util/page.h.node util/tptr.h.node
	@echo src/util/stack.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/stack$(OBJ_EXT) ../src/util/stack.cpp
util/bit_util.h.node: ../src/util/bit_util.h
	@echo done > util/bit_util.h.node
util/mpz$(OBJ_EXT): ../src/util/mpz.cpp util/mpz.h.node util/buffer.h.node util/trace.h.node util/hash.h.node util/bit_util.h.node util/mpn.h.node
	@echo src/util/mpz.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpz$(OBJ_EXT) ../src/util/mpz.cpp
util/warning$(OBJ_EXT): ../src/util/warning.cpp util/error_codes.h.node util/util.h.node util/buffer.h.node util/vector.h.node
	@echo src/util/warning.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/warning$(OBJ_EXT) ../src/util/warning.cpp
util/map.h.node: ../src/util/map.h util/hashtable.h.node
	@echo done > util/map.h.node
util/dictionary.h.node: ../src/util/dictionary.h util/map.h.node util/symbol.h.node
	@echo done > util/dictionary.h.node
util/params$(OBJ_EXT): ../src/util/params.cpp util/params.h.node util/rational.h.node util/symbol.h.node util/dictionary.h.node
	@echo src/util/params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/params$(OBJ_EXT) ../src/util/params.cpp
util/luby$(OBJ_EXT): ../src/util/luby.cpp
	@echo src/util/luby.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/luby$(OBJ_EXT) ../src/util/luby.cpp
util/mpf.h.node: ../src/util/mpf.h util/mpz.h.node util/mpq.h.node util/map.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node util/hash.h.node
	@echo done > util/mpf.h.node
util/mpf$(OBJ_EXT): ../src/util/mpf.cpp util/mpf.h.node
	@echo src/util/mpf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpf$(OBJ_EXT) ../src/util/mpf.cpp
util/prime_generator.h.node: ../src/util/prime_generator.h util/vector.h.node util/z3_exception.h.node util/util.h.node
	@echo done > util/prime_generator.h.node
util/prime_generator$(OBJ_EXT): ../src/util/prime_generator.cpp util/prime_generator.h.node util/mutex.h.node
	@echo src/util/prime_generator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/prime_generator$(OBJ_EXT) ../src/util/prime_generator.cpp
util/statistics.h.node: ../src/util/statistics.h util/vector.h.node util/rlimit.h.node
	@echo done > util/statistics.h.node
util/smt2_util.h.node: ../src/util/smt2_util.h util/symbol.h.node
	@echo done > util/smt2_util.h.node
util/statistics$(OBJ_EXT): ../src/util/statistics.cpp util/statistics.h.node util/map.h.node util/str_hashtable.h.node util/buffer.h.node util/smt2_util.h.node
	@echo src/util/statistics.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/statistics$(OBJ_EXT) ../src/util/statistics.cpp
util/hash$(OBJ_EXT): ../src/util/hash.cpp util/debug.h.node util/hash.h.node
	@echo src/util/hash.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/hash$(OBJ_EXT) ../src/util/hash.cpp
util/timeout.h.node: ../src/util/timeout.h
	@echo done > util/timeout.h.node
util/timeout$(OBJ_EXT): ../src/util/timeout.cpp util/util.h.node util/timeout.h.node util/error_codes.h.node util/event_handler.h.node util/scoped_timer.h.node
	@echo src/util/timeout.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/timeout$(OBJ_EXT) ../src/util/timeout.cpp
util/small_object_allocator$(OBJ_EXT): ../src/util/small_object_allocator.cpp util/memory_manager.h.node util/small_object_allocator.h.node util/debug.h.node util/util.h.node util/vector.h.node
	@echo src/util/small_object_allocator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/small_object_allocator$(OBJ_EXT) ../src/util/small_object_allocator.cpp
util/approx_set.h.node: ../src/util/approx_set.h util/debug.h.node
	@echo done > util/approx_set.h.node
util/approx_set$(OBJ_EXT): ../src/util/approx_set.cpp util/approx_set.h.node
	@echo src/util/approx_set.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/approx_set$(OBJ_EXT) ../src/util/approx_set.cpp
util/uint_set.h.node: ../src/util/uint_set.h util/util.h.node util/vector.h.node
	@echo done > util/uint_set.h.node
util/obj_hashtable.h.node: ../src/util/obj_hashtable.h util/hash.h.node util/hashtable.h.node
	@echo done > util/obj_hashtable.h.node
util/region.h.node: ../src/util/region.h util/vector.h.node
	@echo done > util/region.h.node
util/obj_ref.h.node: ../src/util/obj_ref.h
	@echo done > util/obj_ref.h.node
util/trail.h.node: ../src/util/trail.h util/obj_hashtable.h.node util/region.h.node util/obj_ref.h.node util/vector.h.node
	@echo done > util/trail.h.node
util/union_find.h.node: ../src/util/union_find.h util/trail.h.node util/trace.h.node
	@echo done > util/union_find.h.node
util/state_graph.h.node: ../src/util/state_graph.h util/map.h.node util/uint_set.h.node util/union_find.h.node util/vector.h.node
	@echo done > util/state_graph.h.node
util/state_graph$(OBJ_EXT): ../src/util/state_graph.cpp util/state_graph.h.node
	@echo src/util/state_graph.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/state_graph$(OBJ_EXT) ../src/util/state_graph.cpp
util/inf_s_integer.h.node: ../src/util/inf_s_integer.h util/s_integer.h.node util/rational.h.node
	@echo done > util/inf_s_integer.h.node
util/inf_s_integer$(OBJ_EXT): ../src/util/inf_s_integer.cpp util/inf_s_integer.h.node
	@echo src/util/inf_s_integer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/inf_s_integer$(OBJ_EXT) ../src/util/inf_s_integer.cpp
util/id_gen.h.node: ../src/util/id_gen.h util/vector.h.node util/util.h.node
	@echo done > util/id_gen.h.node
util/mpfx.h.node: ../src/util/mpfx.h util/id_gen.h.node util/util.h.node util/vector.h.node util/z3_exception.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node util/mpn.h.node
	@echo done > util/mpfx.h.node
util/mpfx$(OBJ_EXT): ../src/util/mpfx.cpp util/mpfx.h.node util/mpn.h.node util/mpz.h.node util/mpq.h.node util/bit_util.h.node util/trace.h.node
	@echo src/util/mpfx.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpfx$(OBJ_EXT) ../src/util/mpfx.cpp
util/mpff.h.node: ../src/util/mpff.h util/id_gen.h.node util/util.h.node util/vector.h.node util/z3_exception.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node util/mpn.h.node
	@echo done > util/mpff.h.node
util/mpff$(OBJ_EXT): ../src/util/mpff.cpp util/mpff.h.node util/mpn.h.node util/mpz.h.node util/mpq.h.node util/bit_util.h.node util/trace.h.node
	@echo src/util/mpff.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpff$(OBJ_EXT) ../src/util/mpff.cpp
util/common_msgs$(OBJ_EXT): ../src/util/common_msgs.cpp util/common_msgs.h.node
	@echo src/util/common_msgs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/common_msgs$(OBJ_EXT) ../src/util/common_msgs.cpp
util/symbol$(OBJ_EXT): ../src/util/symbol.cpp util/symbol.h.node util/mutex.h.node util/str_hashtable.h.node util/region.h.node util/string_buffer.h.node
	@echo src/util/symbol.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/symbol$(OBJ_EXT) ../src/util/symbol.cpp
util/fixed_bit_vector.h.node: ../src/util/fixed_bit_vector.h util/debug.h.node util/small_object_allocator.h.node
	@echo done > util/fixed_bit_vector.h.node
util/fixed_bit_vector$(OBJ_EXT): ../src/util/fixed_bit_vector.cpp util/fixed_bit_vector.h.node util/trace.h.node util/hash.h.node
	@echo src/util/fixed_bit_vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/fixed_bit_vector$(OBJ_EXT) ../src/util/fixed_bit_vector.cpp
util/inf_rational.h.node: ../src/util/inf_rational.h util/debug.h.node util/vector.h.node util/rational.h.node
	@echo done > util/inf_rational.h.node
util/inf_rational$(OBJ_EXT): ../src/util/inf_rational.cpp util/inf_rational.h.node
	@echo src/util/inf_rational.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/inf_rational$(OBJ_EXT) ../src/util/inf_rational.cpp
util/page$(OBJ_EXT): ../src/util/page.cpp util/page.h.node util/debug.h.node
	@echo src/util/page.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/page$(OBJ_EXT) ../src/util/page.cpp
util/ref_vector.h.node: ../src/util/ref_vector.h util/vector.h.node util/obj_ref.h.node
	@echo done > util/ref_vector.h.node
util/sexpr.h.node: ../src/util/sexpr.h util/rational.h.node util/symbol.h.node util/obj_ref.h.node util/ref_vector.h.node
	@echo done > util/sexpr.h.node
util/sexpr$(OBJ_EXT): ../src/util/sexpr.cpp util/sexpr.h.node util/vector.h.node util/buffer.h.node
	@echo src/util/sexpr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/sexpr$(OBJ_EXT) ../src/util/sexpr.cpp
util/z3_exception$(OBJ_EXT): ../src/util/z3_exception.cpp util/z3_exception.h.node util/warning.h.node util/error_codes.h.node util/debug.h.node
	@echo src/util/z3_exception.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/z3_exception$(OBJ_EXT) ../src/util/z3_exception.cpp
util/cmd_context_types$(OBJ_EXT): ../src/util/cmd_context_types.cpp util/cmd_context_types.h.node
	@echo src/util/cmd_context_types.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/cmd_context_types$(OBJ_EXT) ../src/util/cmd_context_types.cpp
util/smt2_util$(OBJ_EXT): ../src/util/smt2_util.cpp util/smt2_util.h.node
	@echo src/util/smt2_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/smt2_util$(OBJ_EXT) ../src/util/smt2_util.cpp
util/timeit.h.node: ../src/util/timeit.h
	@echo done > util/timeit.h.node
util/stopwatch.h.node: ../src/util/stopwatch.h util/debug.h.node
	@echo done > util/stopwatch.h.node
util/timeit$(OBJ_EXT): ../src/util/timeit.cpp util/timeit.h.node util/memory_manager.h.node util/stopwatch.h.node
	@echo src/util/timeit.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/timeit$(OBJ_EXT) ../src/util/timeit.cpp
util/trace$(OBJ_EXT): ../src/util/trace.cpp util/trace.h.node util/str_hashtable.h.node
	@echo src/util/trace.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/trace$(OBJ_EXT) ../src/util/trace.cpp
util/util$(OBJ_EXT): ../src/util/util.cpp util/util.h.node
	@echo src/util/util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/util$(OBJ_EXT) ../src/util/util.cpp
util/mpn$(OBJ_EXT): ../src/util/mpn.cpp util/debug.h.node util/trace.h.node util/buffer.h.node util/mpn.h.node
	@echo src/util/mpn.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/mpn$(OBJ_EXT) ../src/util/mpn.cpp
util/scoped_timer$(OBJ_EXT): ../src/util/scoped_timer.cpp util/scoped_timer.h.node util/mutex.h.node util/util.h.node
	@echo src/util/scoped_timer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/scoped_timer$(OBJ_EXT) ../src/util/scoped_timer.cpp
util/gparams$(OBJ_EXT): ../src/util/gparams.cpp util/gparams.h.node util/str_hashtable.h.node util/trace.h.node util/mutex.h.node util/region.h.node util/map.h.node
	@echo src/util/gparams.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/gparams$(OBJ_EXT) ../src/util/gparams.cpp
util/min_cut.h.node: ../src/util/min_cut.h util/vector.h.node
	@echo done > util/min_cut.h.node
util/min_cut$(OBJ_EXT): ../src/util/min_cut.cpp util/min_cut.h.node util/trace.h.node
	@echo src/util/min_cut.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/min_cut$(OBJ_EXT) ../src/util/min_cut.cpp
util/lbool.h.node: ../src/util/lbool.h util/util.h.node
	@echo done > util/lbool.h.node
util/lbool$(OBJ_EXT): ../src/util/lbool.cpp util/lbool.h.node
	@echo src/util/lbool.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/lbool$(OBJ_EXT) ../src/util/lbool.cpp
util/permutation.h.node: ../src/util/permutation.h util/vector.h.node
	@echo done > util/permutation.h.node
util/permutation$(OBJ_EXT): ../src/util/permutation.cpp util/permutation.h.node
	@echo src/util/permutation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/permutation$(OBJ_EXT) ../src/util/permutation.cpp
util/hwf.h.node: ../src/util/hwf.h util/mpz.h.node util/mpq.h.node util/mpf.h.node
	@echo done > util/hwf.h.node
util/hwf$(OBJ_EXT): ../src/util/hwf.cpp util/hwf.h.node
	@echo src/util/hwf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/hwf$(OBJ_EXT) ../src/util/hwf.cpp
util/bit_util$(OBJ_EXT): ../src/util/bit_util.cpp util/bit_util.h.node util/util.h.node util/debug.h.node
	@echo src/util/bit_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/bit_util$(OBJ_EXT) ../src/util/bit_util.cpp
util/rational$(OBJ_EXT): ../src/util/rational.cpp util/mutex.h.node util/util.h.node util/rational.h.node
	@echo src/util/rational.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/rational$(OBJ_EXT) ../src/util/rational.cpp
util/region$(OBJ_EXT): ../src/util/region.cpp util/region.h.node util/tptr.h.node util/debug.h.node util/memory_manager.h.node util/page.h.node
	@echo src/util/region.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/region$(OBJ_EXT) ../src/util/region.cpp
util/approx_nat.h.node: ../src/util/approx_nat.h
	@echo done > util/approx_nat.h.node
util/approx_nat$(OBJ_EXT): ../src/util/approx_nat.cpp util/approx_nat.h.node
	@echo src/util/approx_nat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_0) $(CXX_OUT_FLAG)util/approx_nat$(OBJ_EXT) ../src/util/approx_nat.cpp
util/util$(LIB_EXT): util/s_integer$(OBJ_EXT) util/rlimit$(OBJ_EXT) util/memory_manager$(OBJ_EXT) util/mpbq$(OBJ_EXT) util/bit_vector$(OBJ_EXT) util/mpq_inf$(OBJ_EXT) util/inf_int_rational$(OBJ_EXT) util/env_params$(OBJ_EXT) util/debug$(OBJ_EXT) util/mpq$(OBJ_EXT) util/scoped_ctrl_c$(OBJ_EXT) util/stack$(OBJ_EXT) util/mpz$(OBJ_EXT) util/warning$(OBJ_EXT) util/params$(OBJ_EXT) util/luby$(OBJ_EXT) util/mpf$(OBJ_EXT) util/prime_generator$(OBJ_EXT) util/statistics$(OBJ_EXT) util/hash$(OBJ_EXT) util/timeout$(OBJ_EXT) util/small_object_allocator$(OBJ_EXT) util/approx_set$(OBJ_EXT) util/state_graph$(OBJ_EXT) util/inf_s_integer$(OBJ_EXT) util/mpfx$(OBJ_EXT) util/mpff$(OBJ_EXT) util/common_msgs$(OBJ_EXT) util/symbol$(OBJ_EXT) util/fixed_bit_vector$(OBJ_EXT) util/inf_rational$(OBJ_EXT) util/page$(OBJ_EXT) util/sexpr$(OBJ_EXT) util/z3_exception$(OBJ_EXT) util/cmd_context_types$(OBJ_EXT) util/smt2_util$(OBJ_EXT) util/timeit$(OBJ_EXT) util/trace$(OBJ_EXT) util/util$(OBJ_EXT) util/mpn$(OBJ_EXT) util/scoped_timer$(OBJ_EXT) util/gparams$(OBJ_EXT) util/min_cut$(OBJ_EXT) util/lbool$(OBJ_EXT) util/permutation$(OBJ_EXT) util/hwf$(OBJ_EXT) util/bit_util$(OBJ_EXT) util/rational$(OBJ_EXT) util/region$(OBJ_EXT) util/approx_nat$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)util/util$(LIB_EXT) util/s_integer$(OBJ_EXT) util/rlimit$(OBJ_EXT) util/memory_manager$(OBJ_EXT) util/mpbq$(OBJ_EXT) util/bit_vector$(OBJ_EXT) util/mpq_inf$(OBJ_EXT) util/inf_int_rational$(OBJ_EXT) util/env_params$(OBJ_EXT) util/debug$(OBJ_EXT) util/mpq$(OBJ_EXT) util/scoped_ctrl_c$(OBJ_EXT) util/stack$(OBJ_EXT) util/mpz$(OBJ_EXT) util/warning$(OBJ_EXT) util/params$(OBJ_EXT) util/luby$(OBJ_EXT) util/mpf$(OBJ_EXT) util/prime_generator$(OBJ_EXT) util/statistics$(OBJ_EXT) util/hash$(OBJ_EXT) util/timeout$(OBJ_EXT) util/small_object_allocator$(OBJ_EXT) util/approx_set$(OBJ_EXT) util/state_graph$(OBJ_EXT) util/inf_s_integer$(OBJ_EXT) util/mpfx$(OBJ_EXT) util/mpff$(OBJ_EXT) util/common_msgs$(OBJ_EXT) util/symbol$(OBJ_EXT) util/fixed_bit_vector$(OBJ_EXT) util/inf_rational$(OBJ_EXT) util/page$(OBJ_EXT) util/sexpr$(OBJ_EXT) util/z3_exception$(OBJ_EXT) util/cmd_context_types$(OBJ_EXT) util/smt2_util$(OBJ_EXT) util/timeit$(OBJ_EXT) util/trace$(OBJ_EXT) util/util$(OBJ_EXT) util/mpn$(OBJ_EXT) util/scoped_timer$(OBJ_EXT) util/gparams$(OBJ_EXT) util/min_cut$(OBJ_EXT) util/lbool$(OBJ_EXT) util/permutation$(OBJ_EXT) util/hwf$(OBJ_EXT) util/bit_util$(OBJ_EXT) util/rational$(OBJ_EXT) util/region$(OBJ_EXT) util/approx_nat$(OBJ_EXT)
util: util/util$(LIB_EXT)

includes_1 = -I../src/util -I../src
util/mpzzp.h.node: ../src/util/mpzzp.h util/mpz.h.node
	@echo done > util/mpzzp.h.node
util/basic_interval.h.node: ../src/util/basic_interval.h
	@echo done > util/basic_interval.h.node
util/mpbqi.h.node: ../src/util/mpbqi.h util/mpbq.h.node util/basic_interval.h.node
	@echo done > util/mpbqi.h.node
util/sign.h.node: ../src/util/sign.h
	@echo done > util/sign.h.node
math/polynomial/polynomial.h.node: ../src/math/polynomial/polynomial.h util/mpz.h.node util/rational.h.node util/obj_ref.h.node util/ref_vector.h.node util/z3_exception.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node util/params.h.node util/mpbqi.h.node util/rlimit.h.node util/lbool.h.node util/sign.h.node
	@echo done > math/polynomial/polynomial.h.node
math/polynomial/upolynomial.h.node: ../src/math/polynomial/upolynomial.h util/mpzzp.h.node util/rational.h.node math/polynomial/polynomial.h.node util/z3_exception.h.node util/mpbq.h.node util/rlimit.h.node
	@echo done > math/polynomial/upolynomial.h.node
math/polynomial/upolynomial_factorization.h.node: ../src/math/polynomial/upolynomial_factorization.h math/polynomial/upolynomial.h.node math/polynomial/polynomial.h.node util/bit_vector.h.node util/z3_exception.h.node
	@echo done > math/polynomial/upolynomial_factorization.h.node
math/polynomial/upolynomial_factorization_int.h.node: ../src/math/polynomial/upolynomial_factorization_int.h math/polynomial/upolynomial_factorization.h.node
	@echo done > math/polynomial/upolynomial_factorization_int.h.node
math/polynomial/upolynomial_factorization$(OBJ_EXT): ../src/math/polynomial/upolynomial_factorization.cpp util/trace.h.node util/util.h.node math/polynomial/upolynomial_factorization_int.h.node util/prime_generator.h.node
	@echo src/math/polynomial/upolynomial_factorization.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/upolynomial_factorization$(OBJ_EXT) ../src/math/polynomial/upolynomial_factorization.cpp
math/polynomial/polynomial_cache.h.node: ../src/math/polynomial/polynomial_cache.h math/polynomial/polynomial.h.node
	@echo done > math/polynomial/polynomial_cache.h.node
util/chashtable.h.node: ../src/util/chashtable.h util/memory_manager.h.node util/debug.h.node util/trace.h.node util/tptr.h.node util/util.h.node util/hashtable.h.node
	@echo done > util/chashtable.h.node
math/polynomial/polynomial_cache$(OBJ_EXT): ../src/math/polynomial/polynomial_cache.cpp math/polynomial/polynomial_cache.h.node util/chashtable.h.node
	@echo src/math/polynomial/polynomial_cache.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/polynomial_cache$(OBJ_EXT) ../src/math/polynomial/polynomial_cache.cpp
math/polynomial/sexpr2upolynomial.h.node: ../src/math/polynomial/sexpr2upolynomial.h math/polynomial/upolynomial.h.node util/cmd_context_types.h.node
	@echo done > math/polynomial/sexpr2upolynomial.h.node
math/polynomial/sexpr2upolynomial$(OBJ_EXT): ../src/math/polynomial/sexpr2upolynomial.cpp math/polynomial/sexpr2upolynomial.h.node util/sexpr.h.node
	@echo src/math/polynomial/sexpr2upolynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/sexpr2upolynomial$(OBJ_EXT) ../src/math/polynomial/sexpr2upolynomial.cpp
util/scoped_ptr_vector.h.node: ../src/util/scoped_ptr_vector.h util/vector.h.node util/util.h.node
	@echo done > util/scoped_ptr_vector.h.node
math/polynomial/polynomial_primes.h.node: ../src/math/polynomial/polynomial_primes.h
	@echo done > math/polynomial/polynomial_primes.h.node
math/polynomial/algebraic_numbers.h.node: ../src/math/polynomial/algebraic_numbers.h util/rational.h.node util/mpq.h.node math/polynomial/polynomial.h.node util/z3_exception.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node util/tptr.h.node util/statistics.h.node util/params.h.node util/rlimit.h.node
	@echo done > math/polynomial/algebraic_numbers.h.node
math/polynomial/linear_eq_solver.h.node: ../src/math/polynomial/linear_eq_solver.h
	@echo done > math/polynomial/linear_eq_solver.h.node
util/scoped_numeral_buffer.h.node: ../src/util/scoped_numeral_buffer.h util/buffer.h.node
	@echo done > util/scoped_numeral_buffer.h.node
util/ref_buffer.h.node: ../src/util/ref_buffer.h util/buffer.h.node util/obj_ref.h.node util/ref_vector.h.node
	@echo done > util/ref_buffer.h.node
math/polynomial/polynomial$(OBJ_EXT): ../src/math/polynomial/polynomial.cpp math/polynomial/polynomial.h.node util/vector.h.node util/chashtable.h.node util/small_object_allocator.h.node util/id_gen.h.node util/buffer.h.node util/scoped_ptr_vector.h.node math/polynomial/upolynomial_factorization.h.node math/polynomial/polynomial_primes.h.node util/permutation.h.node math/polynomial/algebraic_numbers.h.node util/mpzzp.h.node util/timeit.h.node math/polynomial/linear_eq_solver.h.node util/scoped_numeral_buffer.h.node util/ref_buffer.h.node util/common_msgs.h.node
	@echo src/math/polynomial/polynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/polynomial$(OBJ_EXT) ../src/math/polynomial/polynomial.cpp
math/polynomial/algebraic_params.hpp.node: ../src/math/polynomial/algebraic_params.hpp util/params.h.node util/gparams.h.node
	@echo done > math/polynomial/algebraic_params.hpp.node
math/polynomial/algebraic_numbers$(OBJ_EXT): ../src/math/polynomial/algebraic_numbers.cpp util/mpbq.h.node util/basic_interval.h.node util/scoped_ptr_vector.h.node util/mpbqi.h.node util/timeit.h.node util/common_msgs.h.node math/polynomial/algebraic_numbers.h.node math/polynomial/upolynomial.h.node math/polynomial/sexpr2upolynomial.h.node math/polynomial/algebraic_params.hpp.node
	@echo src/math/polynomial/algebraic_numbers.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/algebraic_numbers$(OBJ_EXT) ../src/math/polynomial/algebraic_numbers.cpp
math/polynomial/rpolynomial.h.node: ../src/math/polynomial/rpolynomial.h util/mpz.h.node util/rational.h.node util/obj_ref.h.node util/ref_vector.h.node util/z3_exception.h.node math/polynomial/polynomial.h.node
	@echo done > math/polynomial/rpolynomial.h.node
math/polynomial/rpolynomial$(OBJ_EXT): ../src/math/polynomial/rpolynomial.cpp math/polynomial/rpolynomial.h.node util/tptr.h.node util/buffer.h.node
	@echo src/math/polynomial/rpolynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/rpolynomial$(OBJ_EXT) ../src/math/polynomial/rpolynomial.cpp
math/polynomial/upolynomial$(OBJ_EXT): ../src/math/polynomial/upolynomial.cpp math/polynomial/upolynomial.h.node math/polynomial/upolynomial_factorization.h.node math/polynomial/polynomial_primes.h.node util/buffer.h.node util/common_msgs.h.node
	@echo src/math/polynomial/upolynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_1) $(CXX_OUT_FLAG)math/polynomial/upolynomial$(OBJ_EXT) ../src/math/polynomial/upolynomial.cpp
math/polynomial/polynomial$(LIB_EXT): math/polynomial/upolynomial_factorization$(OBJ_EXT) math/polynomial/polynomial_cache$(OBJ_EXT) math/polynomial/sexpr2upolynomial$(OBJ_EXT) math/polynomial/polynomial$(OBJ_EXT) math/polynomial/algebraic_numbers$(OBJ_EXT) math/polynomial/rpolynomial$(OBJ_EXT) math/polynomial/upolynomial$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/polynomial/polynomial$(LIB_EXT) math/polynomial/upolynomial_factorization$(OBJ_EXT) math/polynomial/polynomial_cache$(OBJ_EXT) math/polynomial/sexpr2upolynomial$(OBJ_EXT) math/polynomial/polynomial$(OBJ_EXT) math/polynomial/algebraic_numbers$(OBJ_EXT) math/polynomial/rpolynomial$(OBJ_EXT) math/polynomial/upolynomial$(OBJ_EXT)
polynomial: math/polynomial/polynomial$(LIB_EXT)

includes_2 = -I../src/util -I../src
util/ext_numeral.h.node: ../src/util/ext_numeral.h util/debug.h.node
	@echo done > util/ext_numeral.h.node
math/interval/interval.h.node: ../src/math/interval/interval.h util/mpq.h.node util/ext_numeral.h.node util/rlimit.h.node
	@echo done > math/interval/interval.h.node
math/interval/interval_def.h.node: ../src/math/interval/interval_def.h math/interval/interval.h.node util/debug.h.node util/trace.h.node util/scoped_numeral.h.node util/common_msgs.h.node
	@echo done > math/interval/interval_def.h.node
math/interval/interval_mpq$(OBJ_EXT): ../src/math/interval/interval_mpq.cpp math/interval/interval_def.h.node
	@echo src/math/interval/interval_mpq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_2) $(CXX_OUT_FLAG)math/interval/interval_mpq$(OBJ_EXT) ../src/math/interval/interval_mpq.cpp
util/dependency.h.node: ../src/util/dependency.h util/vector.h.node util/region.h.node
	@echo done > util/dependency.h.node
math/interval/dep_intervals.h.node: ../src/math/interval/dep_intervals.h util/dependency.h.node util/region.h.node util/rational.h.node util/rlimit.h.node math/interval/interval.h.node
	@echo done > math/interval/dep_intervals.h.node
math/interval/dep_intervals$(OBJ_EXT): ../src/math/interval/dep_intervals.cpp math/interval/interval_def.h.node math/interval/dep_intervals.h.node util/mpq.h.node
	@echo src/math/interval/dep_intervals.cpp
	@$(CXX) $(CXXFLAGS) $(includes_2) $(CXX_OUT_FLAG)math/interval/dep_intervals$(OBJ_EXT) ../src/math/interval/dep_intervals.cpp
math/interval/interval$(LIB_EXT): math/interval/interval_mpq$(OBJ_EXT) math/interval/dep_intervals$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/interval/interval$(LIB_EXT) math/interval/interval_mpq$(OBJ_EXT) math/interval/dep_intervals$(OBJ_EXT)
interval: math/interval/interval$(LIB_EXT)

includes_3 = -I../src/util -I../src/math/interval -I../src
math/dd/dd_bdd.h.node: ../src/math/dd/dd_bdd.h util/vector.h.node util/map.h.node util/small_object_allocator.h.node
	@echo done > math/dd/dd_bdd.h.node
math/dd/dd_bdd$(OBJ_EXT): ../src/math/dd/dd_bdd.cpp util/trace.h.node util/stopwatch.h.node math/dd/dd_bdd.h.node
	@echo src/math/dd/dd_bdd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_3) $(CXX_OUT_FLAG)math/dd/dd_bdd$(OBJ_EXT) ../src/math/dd/dd_bdd.cpp
math/dd/dd_pdd.h.node: ../src/math/dd/dd_pdd.h util/vector.h.node util/map.h.node util/small_object_allocator.h.node util/rational.h.node
	@echo done > math/dd/dd_pdd.h.node
math/dd/dd_pdd$(OBJ_EXT): ../src/math/dd/dd_pdd.cpp util/trace.h.node util/stopwatch.h.node math/dd/dd_pdd.h.node
	@echo src/math/dd/dd_pdd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_3) $(CXX_OUT_FLAG)math/dd/dd_pdd$(OBJ_EXT) ../src/math/dd/dd_pdd.cpp
math/dd/dd$(LIB_EXT): math/dd/dd_bdd$(OBJ_EXT) math/dd/dd_pdd$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/dd/dd$(LIB_EXT) math/dd/dd_bdd$(OBJ_EXT) math/dd/dd_pdd$(OBJ_EXT)
dd: math/dd/dd$(LIB_EXT)

includes_4 = -I../src/util -I../src
math/simplex/sparse_matrix.h.node: ../src/math/simplex/sparse_matrix.h util/mpq_inf.h.node util/statistics.h.node
	@echo done > math/simplex/sparse_matrix.h.node
util/heap.h.node: ../src/util/heap.h util/vector.h.node util/debug.h.node
	@echo done > util/heap.h.node
math/simplex/simplex.h.node: ../src/math/simplex/simplex.h math/simplex/sparse_matrix.h.node util/mpq_inf.h.node util/heap.h.node util/lbool.h.node util/uint_set.h.node
	@echo done > math/simplex/simplex.h.node
math/simplex/sparse_matrix_def.h.node: ../src/math/simplex/sparse_matrix_def.h math/simplex/sparse_matrix.h.node util/uint_set.h.node
	@echo done > math/simplex/sparse_matrix_def.h.node
math/simplex/simplex_def.h.node: ../src/math/simplex/simplex_def.h
	@echo done > math/simplex/simplex_def.h.node
math/simplex/simplex$(OBJ_EXT): ../src/math/simplex/simplex.cpp math/simplex/simplex.h.node math/simplex/sparse_matrix_def.h.node math/simplex/simplex_def.h.node util/rational.h.node util/inf_rational.h.node
	@echo src/math/simplex/simplex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_4) $(CXX_OUT_FLAG)math/simplex/simplex$(OBJ_EXT) ../src/math/simplex/simplex.cpp
util/inf_eps_rational.h.node: ../src/util/inf_eps_rational.h util/debug.h.node util/vector.h.node util/rational.h.node util/inf_rational.h.node
	@echo done > util/inf_eps_rational.h.node
math/simplex/model_based_opt.h.node: ../src/math/simplex/model_based_opt.h util/util.h.node util/rational.h.node util/inf_eps_rational.h.node
	@echo done > math/simplex/model_based_opt.h.node
math/simplex/model_based_opt$(OBJ_EXT): ../src/math/simplex/model_based_opt.cpp math/simplex/model_based_opt.h.node util/uint_set.h.node util/z3_exception.h.node
	@echo src/math/simplex/model_based_opt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_4) $(CXX_OUT_FLAG)math/simplex/model_based_opt$(OBJ_EXT) ../src/math/simplex/model_based_opt.cpp
math/simplex/bit_matrix.h.node: ../src/math/simplex/bit_matrix.h util/region.h.node util/vector.h.node
	@echo done > math/simplex/bit_matrix.h.node
math/simplex/bit_matrix$(OBJ_EXT): ../src/math/simplex/bit_matrix.cpp math/simplex/bit_matrix.h.node util/stopwatch.h.node
	@echo src/math/simplex/bit_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_4) $(CXX_OUT_FLAG)math/simplex/bit_matrix$(OBJ_EXT) ../src/math/simplex/bit_matrix.cpp
math/simplex/simplex$(LIB_EXT): math/simplex/simplex$(OBJ_EXT) math/simplex/model_based_opt$(OBJ_EXT) math/simplex/bit_matrix$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/simplex/simplex$(LIB_EXT) math/simplex/simplex$(OBJ_EXT) math/simplex/model_based_opt$(OBJ_EXT) math/simplex/bit_matrix$(OBJ_EXT)
simplex: math/simplex/simplex$(LIB_EXT)

includes_5 = -I../src/util -I../src
util/checked_int64.h.node: ../src/util/checked_int64.h util/z3_exception.h.node util/rational.h.node
	@echo done > util/checked_int64.h.node
math/hilbert/hilbert_basis.h.node: ../src/math/hilbert/hilbert_basis.h util/rational.h.node util/lbool.h.node util/statistics.h.node util/checked_int64.h.node util/rlimit.h.node
	@echo done > math/hilbert/hilbert_basis.h.node
math/hilbert/heap_trie.h.node: ../src/math/hilbert/heap_trie.h util/map.h.node util/vector.h.node util/buffer.h.node util/statistics.h.node util/small_object_allocator.h.node
	@echo done > math/hilbert/heap_trie.h.node
math/hilbert/hilbert_basis$(OBJ_EXT): ../src/math/hilbert/hilbert_basis.cpp math/hilbert/hilbert_basis.h.node util/heap.h.node util/map.h.node math/hilbert/heap_trie.h.node util/stopwatch.h.node
	@echo src/math/hilbert/hilbert_basis.cpp
	@$(CXX) $(CXXFLAGS) $(includes_5) $(CXX_OUT_FLAG)math/hilbert/hilbert_basis$(OBJ_EXT) ../src/math/hilbert/hilbert_basis.cpp
math/hilbert/hilbert$(LIB_EXT): math/hilbert/hilbert_basis$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/hilbert/hilbert$(LIB_EXT) math/hilbert/hilbert_basis$(OBJ_EXT)
hilbert: math/hilbert/hilbert$(LIB_EXT)

includes_6 = -I../src/util -I../src
math/automata/automaton.h.node: ../src/math/automata/automaton.h util/util.h.node util/vector.h.node util/uint_set.h.node util/trace.h.node
	@echo done > math/automata/automaton.h.node
math/automata/automaton$(OBJ_EXT): ../src/math/automata/automaton.cpp math/automata/automaton.h.node
	@echo src/math/automata/automaton.cpp
	@$(CXX) $(CXXFLAGS) $(includes_6) $(CXX_OUT_FLAG)math/automata/automaton$(OBJ_EXT) ../src/math/automata/automaton.cpp
math/automata/automata$(LIB_EXT): math/automata/automaton$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/automata/automata$(LIB_EXT) math/automata/automaton$(OBJ_EXT)
automata: math/automata/automata$(LIB_EXT)

includes_7 = -I../src/math/interval -I../src/util -I../src
math/realclosure/mpz_matrix.h.node: ../src/math/realclosure/mpz_matrix.h util/mpz.h.node
	@echo done > math/realclosure/mpz_matrix.h.node
math/realclosure/mpz_matrix$(OBJ_EXT): ../src/math/realclosure/mpz_matrix.cpp math/realclosure/mpz_matrix.h.node util/buffer.h.node
	@echo src/math/realclosure/mpz_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_7) $(CXX_OUT_FLAG)math/realclosure/mpz_matrix$(OBJ_EXT) ../src/math/realclosure/mpz_matrix.cpp
math/realclosure/realclosure.h.node: ../src/math/realclosure/realclosure.h util/mpq.h.node util/params.h.node util/scoped_numeral.h.node util/scoped_numeral_vector.h.node math/interval/interval.h.node util/z3_exception.h.node util/rlimit.h.node
	@echo done > math/realclosure/realclosure.h.node
math/realclosure/rcf_params.hpp.node: ../src/math/realclosure/rcf_params.hpp util/params.h.node util/gparams.h.node
	@echo done > math/realclosure/rcf_params.hpp.node
util/array.h.node: ../src/util/array.h
	@echo done > util/array.h.node
math/realclosure/realclosure$(OBJ_EXT): ../src/math/realclosure/realclosure.cpp math/realclosure/realclosure.h.node math/realclosure/rcf_params.hpp.node util/array.h.node util/mpbq.h.node math/realclosure/mpz_matrix.h.node math/interval/interval_def.h.node util/obj_ref.h.node util/ref_vector.h.node util/ref_buffer.h.node util/common_msgs.h.node
	@echo src/math/realclosure/realclosure.cpp
	@$(CXX) $(CXXFLAGS) $(includes_7) $(CXX_OUT_FLAG)math/realclosure/realclosure$(OBJ_EXT) ../src/math/realclosure/realclosure.cpp
math/realclosure/realclosure$(LIB_EXT): math/realclosure/mpz_matrix$(OBJ_EXT) math/realclosure/realclosure$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/realclosure/realclosure$(LIB_EXT) math/realclosure/mpz_matrix$(OBJ_EXT) math/realclosure/realclosure$(OBJ_EXT)
realclosure: math/realclosure/realclosure$(LIB_EXT)

includes_8 = -I../src/math/interval -I../src/util -I../src
util/parray.h.node: ../src/util/parray.h util/vector.h.node util/trace.h.node
	@echo done > util/parray.h.node
math/subpaving/subpaving_types.h.node: ../src/math/subpaving/subpaving_types.h
	@echo done > math/subpaving/subpaving_types.h.node
math/subpaving/subpaving_t.h.node: ../src/math/subpaving/subpaving_t.h util/tptr.h.node util/small_object_allocator.h.node util/chashtable.h.node util/parray.h.node math/interval/interval.h.node util/scoped_numeral_vector.h.node math/subpaving/subpaving_types.h.node util/params.h.node util/statistics.h.node util/lbool.h.node util/id_gen.h.node util/rlimit.h.node
	@echo done > math/subpaving/subpaving_t.h.node
util/f2n.h.node: ../src/util/f2n.h util/mpf.h.node
	@echo done > util/f2n.h.node
math/subpaving/subpaving_mpf.h.node: ../src/math/subpaving/subpaving_mpf.h math/subpaving/subpaving_t.h.node util/mpf.h.node util/f2n.h.node
	@echo done > math/subpaving/subpaving_mpf.h.node
math/subpaving/subpaving_t_def.h.node: ../src/math/subpaving/subpaving_t_def.h math/subpaving/subpaving_t.h.node math/interval/interval_def.h.node util/buffer.h.node util/z3_exception.h.node util/common_msgs.h.node
	@echo done > math/subpaving/subpaving_t_def.h.node
math/subpaving/subpaving_mpf$(OBJ_EXT): ../src/math/subpaving/subpaving_mpf.cpp math/subpaving/subpaving_mpf.h.node math/subpaving/subpaving_t_def.h.node
	@echo src/math/subpaving/subpaving_mpf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_8) $(CXX_OUT_FLAG)math/subpaving/subpaving_mpf$(OBJ_EXT) ../src/math/subpaving/subpaving_mpf.cpp
math/subpaving/subpaving_mpq.h.node: ../src/math/subpaving/subpaving_mpq.h math/subpaving/subpaving_t.h.node util/mpq.h.node
	@echo done > math/subpaving/subpaving_mpq.h.node
math/subpaving/subpaving_mpq$(OBJ_EXT): ../src/math/subpaving/subpaving_mpq.cpp math/subpaving/subpaving_mpq.h.node math/subpaving/subpaving_t_def.h.node
	@echo src/math/subpaving/subpaving_mpq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_8) $(CXX_OUT_FLAG)math/subpaving/subpaving_mpq$(OBJ_EXT) ../src/math/subpaving/subpaving_mpq.cpp
math/subpaving/subpaving_hwf.h.node: ../src/math/subpaving/subpaving_hwf.h math/subpaving/subpaving_t.h.node util/f2n.h.node util/hwf.h.node
	@echo done > math/subpaving/subpaving_hwf.h.node
math/subpaving/subpaving_hwf$(OBJ_EXT): ../src/math/subpaving/subpaving_hwf.cpp math/subpaving/subpaving_hwf.h.node math/subpaving/subpaving_t_def.h.node
	@echo src/math/subpaving/subpaving_hwf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_8) $(CXX_OUT_FLAG)math/subpaving/subpaving_hwf$(OBJ_EXT) ../src/math/subpaving/subpaving_hwf.cpp
math/subpaving/subpaving_mpfx.h.node: ../src/math/subpaving/subpaving_mpfx.h math/subpaving/subpaving_t.h.node util/mpfx.h.node
	@echo done > math/subpaving/subpaving_mpfx.h.node
math/subpaving/subpaving_mpfx$(OBJ_EXT): ../src/math/subpaving/subpaving_mpfx.cpp math/subpaving/subpaving_mpfx.h.node math/subpaving/subpaving_t_def.h.node
	@echo src/math/subpaving/subpaving_mpfx.cpp
	@$(CXX) $(CXXFLAGS) $(includes_8) $(CXX_OUT_FLAG)math/subpaving/subpaving_mpfx$(OBJ_EXT) ../src/math/subpaving/subpaving_mpfx.cpp
math/subpaving/subpaving_mpff.h.node: ../src/math/subpaving/subpaving_mpff.h math/subpaving/subpaving_t.h.node util/mpff.h.node
	@echo done > math/subpaving/subpaving_mpff.h.node
math/subpaving/subpaving_mpff$(OBJ_EXT): ../src/math/subpaving/subpaving_mpff.cpp math/subpaving/subpaving_mpff.h.node math/subpaving/subpaving_t_def.h.node
	@echo src/math/subpaving/subpaving_mpff.cpp
	@$(CXX) $(CXXFLAGS) $(includes_8) $(CXX_OUT_FLAG)math/subpaving/subpaving_mpff$(OBJ_EXT) ../src/math/subpaving/subpaving_mpff.cpp
math/subpaving/subpaving.h.node: ../src/math/subpaving/subpaving.h util/mpq.h.node math/subpaving/subpaving_types.h.node util/params.h.node util/statistics.h.node
	@echo done > math/subpaving/subpaving.h.node
math/subpaving/subpaving$(OBJ_EXT): ../src/math/subpaving/subpaving.cpp math/subpaving/subpaving.h.node math/subpaving/subpaving_types.h.node math/subpaving/subpaving_mpq.h.node math/subpaving/subpaving_mpf.h.node math/subpaving/subpaving_hwf.h.node math/subpaving/subpaving_mpff.h.node math/subpaving/subpaving_mpfx.h.node
	@echo src/math/subpaving/subpaving.cpp
	@$(CXX) $(CXXFLAGS) $(includes_8) $(CXX_OUT_FLAG)math/subpaving/subpaving$(OBJ_EXT) ../src/math/subpaving/subpaving.cpp
math/subpaving/subpaving$(LIB_EXT): math/subpaving/subpaving_mpf$(OBJ_EXT) math/subpaving/subpaving_mpq$(OBJ_EXT) math/subpaving/subpaving_hwf$(OBJ_EXT) math/subpaving/subpaving_mpfx$(OBJ_EXT) math/subpaving/subpaving_mpff$(OBJ_EXT) math/subpaving/subpaving$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/subpaving/subpaving$(LIB_EXT) math/subpaving/subpaving_mpf$(OBJ_EXT) math/subpaving/subpaving_mpq$(OBJ_EXT) math/subpaving/subpaving_hwf$(OBJ_EXT) math/subpaving/subpaving_mpfx$(OBJ_EXT) math/subpaving/subpaving_mpff$(OBJ_EXT) math/subpaving/subpaving$(OBJ_EXT)
subpaving: math/subpaving/subpaving$(LIB_EXT)

includes_9 = -I../src/util -I../src/math/polynomial -I../src
util/optional.h.node: ../src/util/optional.h
	@echo done > util/optional.h.node
util/symbol_table.h.node: ../src/util/symbol_table.h util/vector.h.node util/hashtable.h.node util/hash.h.node util/symbol.h.node
	@echo done > util/symbol_table.h.node
util/ref_pair_vector.h.node: ../src/util/ref_pair_vector.h util/vector.h.node util/obj_ref.h.node util/ref_vector.h.node
	@echo done > util/ref_pair_vector.h.node
util/obj_mark.h.node: ../src/util/obj_mark.h util/bit_vector.h.node
	@echo done > util/obj_mark.h.node
ast/ast.h.node: ../src/ast/ast.h util/vector.h.node util/hashtable.h.node util/buffer.h.node util/symbol.h.node util/rational.h.node util/hash.h.node util/optional.h.node util/trace.h.node util/bit_vector.h.node util/symbol_table.h.node util/tptr.h.node util/memory_manager.h.node util/small_object_allocator.h.node util/obj_ref.h.node util/ref_vector.h.node util/ref_pair_vector.h.node util/ref_buffer.h.node util/obj_mark.h.node util/obj_hashtable.h.node util/id_gen.h.node util/map.h.node util/parray.h.node util/dictionary.h.node util/chashtable.h.node util/z3_exception.h.node util/dependency.h.node util/rlimit.h.node
	@echo done > ast/ast.h.node
ast/for_each_expr.h.node: ../src/ast/for_each_expr.h ast/ast.h.node util/trace.h.node
	@echo done > ast/for_each_expr.h.node
ast/for_each_expr$(OBJ_EXT): ../src/ast/for_each_expr.cpp ast/for_each_expr.h.node
	@echo src/ast/for_each_expr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/for_each_expr$(OBJ_EXT) ../src/ast/for_each_expr.cpp
ast/expr_abstract.h.node: ../src/ast/expr_abstract.h ast/ast.h.node
	@echo done > ast/expr_abstract.h.node
ast/format.h.node: ../src/ast/format.h ast/ast.h.node
	@echo done > ast/format.h.node
ast/arith_decl_plugin.h.node: ../src/ast/arith_decl_plugin.h ast/ast.h.node
	@echo done > ast/arith_decl_plugin.h.node
ast/bv_decl_plugin.h.node: ../src/ast/bv_decl_plugin.h ast/ast.h.node
	@echo done > ast/bv_decl_plugin.h.node
ast/array_decl_plugin.h.node: ../src/ast/array_decl_plugin.h ast/ast.h.node
	@echo done > ast/array_decl_plugin.h.node
ast/fpa_decl_plugin.h.node: ../src/ast/fpa_decl_plugin.h ast/ast.h.node util/id_gen.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node util/mpf.h.node
	@echo done > ast/fpa_decl_plugin.h.node
ast/dl_decl_plugin.h.node: ../src/ast/dl_decl_plugin.h ast/ast.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node
	@echo done > ast/dl_decl_plugin.h.node
ast/seq_decl_plugin.h.node: ../src/ast/seq_decl_plugin.h ast/ast.h.node ast/bv_decl_plugin.h.node util/lbool.h.node
	@echo done > ast/seq_decl_plugin.h.node
ast/datatype_decl_plugin.h.node: ../src/ast/datatype_decl_plugin.h ast/ast.h.node util/buffer.h.node util/symbol_table.h.node util/obj_hashtable.h.node
	@echo done > ast/datatype_decl_plugin.h.node
ast/ast_smt_pp.h.node: ../src/ast/ast_smt_pp.h ast/ast.h.node util/map.h.node
	@echo done > ast/ast_smt_pp.h.node
ast/ast_smt2_pp.h.node: ../src/ast/ast_smt2_pp.h ast/format.h.node util/params.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/fpa_decl_plugin.h.node ast/dl_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/ast_smt_pp.h.node util/smt2_util.h.node
	@echo done > ast/ast_smt2_pp.h.node
ast/ast_pp.h.node: ../src/ast/ast_pp.h ast/ast_smt2_pp.h.node
	@echo done > ast/ast_pp.h.node
ast/expr_abstract$(OBJ_EXT): ../src/ast/expr_abstract.cpp ast/expr_abstract.h.node util/map.h.node ast/ast_pp.h.node
	@echo src/ast/expr_abstract.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr_abstract$(OBJ_EXT) ../src/ast/expr_abstract.cpp
ast/recurse_expr.h.node: ../src/ast/recurse_expr.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/recurse_expr.h.node
ast/recurse_expr_def.h.node: ../src/ast/recurse_expr_def.h ast/recurse_expr.h.node
	@echo done > ast/recurse_expr_def.h.node
ast/format$(OBJ_EXT): ../src/ast/format.cpp ast/format.h.node ast/recurse_expr_def.h.node
	@echo src/ast/format.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/format$(OBJ_EXT) ../src/ast/format.cpp
ast/expr_map.h.node: ../src/ast/expr_map.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/expr_map.h.node
ast/expr_functors.h.node: ../src/ast/expr_functors.h ast/ast.h.node ast/expr_map.h.node
	@echo done > ast/expr_functors.h.node
ast/expr_functors$(OBJ_EXT): ../src/ast/expr_functors.cpp ast/expr_functors.h.node
	@echo src/ast/expr_functors.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr_functors$(OBJ_EXT) ../src/ast/expr_functors.cpp
ast/bv_decl_plugin$(OBJ_EXT): ../src/ast/bv_decl_plugin.cpp ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node util/warning.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/bv_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/bv_decl_plugin$(OBJ_EXT) ../src/ast/bv_decl_plugin.cpp
ast/expr2polynomial.h.node: ../src/ast/expr2polynomial.h ast/ast.h.node math/polynomial/polynomial.h.node
	@echo done > ast/expr2polynomial.h.node
ast/expr2var.h.node: ../src/ast/expr2var.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/expr2var.h.node
ast/expr2polynomial$(OBJ_EXT): ../src/ast/expr2polynomial.cpp ast/expr2polynomial.h.node ast/expr2var.h.node ast/arith_decl_plugin.h.node ast/ast_smt2_pp.h.node util/z3_exception.h.node util/common_msgs.h.node
	@echo src/ast/expr2polynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr2polynomial$(OBJ_EXT) ../src/ast/expr2polynomial.cpp
ast/pb_decl_plugin.h.node: ../src/ast/pb_decl_plugin.h ast/ast.h.node
	@echo done > ast/pb_decl_plugin.h.node
ast/ast_util.h.node: ../src/ast/ast_util.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/ast_util.h.node
ast/pb_decl_plugin$(OBJ_EXT): ../src/ast/pb_decl_plugin.cpp ast/pb_decl_plugin.h.node ast/ast_util.h.node ast/ast_pp.h.node
	@echo src/ast/pb_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/pb_decl_plugin$(OBJ_EXT) ../src/ast/pb_decl_plugin.cpp
ast/arith_decl_plugin$(OBJ_EXT): ../src/ast/arith_decl_plugin.cpp ast/arith_decl_plugin.h.node util/warning.h.node math/polynomial/algebraic_numbers.h.node util/id_gen.h.node ast/ast_smt2_pp.h.node util/gparams.h.node
	@echo src/ast/arith_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/arith_decl_plugin$(OBJ_EXT) ../src/ast/arith_decl_plugin.cpp
ast/well_sorted.h.node: ../src/ast/well_sorted.h
	@echo done > ast/well_sorted.h.node
ast/ast_ll_pp.h.node: ../src/ast/ast_ll_pp.h ast/ast.h.node
	@echo done > ast/ast_ll_pp.h.node
ast/well_sorted$(OBJ_EXT): ../src/ast/well_sorted.cpp ast/for_each_expr.h.node ast/well_sorted.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node util/warning.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/well_sorted.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/well_sorted$(OBJ_EXT) ../src/ast/well_sorted.cpp
ast/special_relations_decl_plugin.h.node: ../src/ast/special_relations_decl_plugin.h ast/ast.h.node
	@echo done > ast/special_relations_decl_plugin.h.node
ast/static_features.h.node: ../src/ast/static_features.h ast/ast.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/fpa_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/special_relations_decl_plugin.h.node util/map.h.node
	@echo done > ast/static_features.h.node
ast/static_features$(OBJ_EXT): ../src/ast/static_features.cpp ast/static_features.h.node ast/ast_pp.h.node
	@echo src/ast/static_features.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/static_features$(OBJ_EXT) ../src/ast/static_features.cpp
ast/recfun_decl_plugin.h.node: ../src/ast/recfun_decl_plugin.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/recfun_decl_plugin.h.node
ast/reg_decl_plugins$(OBJ_EXT): ../src/ast/reg_decl_plugins.cpp ast/ast.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/recfun_decl_plugin.h.node ast/dl_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/fpa_decl_plugin.h.node ast/special_relations_decl_plugin.h.node
	@echo src/ast/reg_decl_plugins.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/reg_decl_plugins$(OBJ_EXT) ../src/ast/reg_decl_plugins.cpp
ast/ast_translation.h.node: ../src/ast/ast_translation.h ast/ast.h.node
	@echo done > ast/ast_translation.h.node
ast/datatype_decl_plugin$(OBJ_EXT): ../src/ast/datatype_decl_plugin.cpp util/warning.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/ast_translation.h.node
	@echo src/ast/datatype_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/datatype_decl_plugin$(OBJ_EXT) ../src/ast/datatype_decl_plugin.cpp
ast/value_generator.h.node: ../src/ast/value_generator.h util/scoped_ptr_vector.h.node ast/ast.h.node
	@echo done > ast/value_generator.h.node
ast/value_generator$(OBJ_EXT): ../src/ast/value_generator.cpp ast/value_generator.h.node ast/datatype_decl_plugin.h.node ast/array_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/seq_decl_plugin.h.node
	@echo src/ast/value_generator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/value_generator$(OBJ_EXT) ../src/ast/value_generator.cpp
ast/for_each_ast.h.node: ../src/ast/for_each_ast.h ast/ast.h.node util/trace.h.node util/map.h.node
	@echo done > ast/for_each_ast.h.node
ast/ast_ll_pp$(OBJ_EXT): ../src/ast/ast_ll_pp.cpp ast/for_each_ast.h.node ast/arith_decl_plugin.h.node
	@echo src/ast/ast_ll_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_ll_pp$(OBJ_EXT) ../src/ast/ast_ll_pp.cpp
ast/ast_printer.h.node: ../src/ast/ast_printer.h ast/ast.h.node ast/ast_smt2_pp.h.node
	@echo done > ast/ast_printer.h.node
ast/pp.h.node: ../src/ast/pp.h ast/format.h.node util/params.h.node
	@echo done > ast/pp.h.node
ast/ast_printer$(OBJ_EXT): ../src/ast/ast_printer.cpp ast/ast_printer.h.node ast/pp.h.node
	@echo src/ast/ast_printer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_printer$(OBJ_EXT) ../src/ast/ast_printer.cpp
ast/pp_params.hpp.node: ../src/ast/pp_params.hpp util/params.h.node util/gparams.h.node
	@echo done > ast/pp_params.hpp.node
ast/pp$(OBJ_EXT): ../src/ast/pp.cpp ast/pp.h.node ast/pp_params.hpp.node
	@echo src/ast/pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/pp$(OBJ_EXT) ../src/ast/pp.cpp
util/top_sort.h.node: ../src/util/top_sort.h util/obj_hashtable.h.node util/vector.h.node util/memory_manager.h.node
	@echo done > util/top_sort.h.node
ast/decl_collector.h.node: ../src/ast/decl_collector.h util/top_sort.h.node ast/ast.h.node ast/datatype_decl_plugin.h.node
	@echo done > ast/decl_collector.h.node
ast/ast_smt_pp$(OBJ_EXT): ../src/ast/ast_smt_pp.cpp util/vector.h.node util/smt2_util.h.node ast/ast_smt_pp.h.node ast/ast_smt2_pp.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/fpa_decl_plugin.h.node ast/for_each_ast.h.node ast/decl_collector.h.node math/polynomial/algebraic_numbers.h.node
	@echo src/ast/ast_smt_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_smt_pp$(OBJ_EXT) ../src/ast/ast_smt_pp.cpp
ast/seq_decl_plugin$(OBJ_EXT): ../src/ast/seq_decl_plugin.cpp ast/seq_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/ast_pp.h.node ast/bv_decl_plugin.h.node
	@echo src/ast/seq_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/seq_decl_plugin$(OBJ_EXT) ../src/ast/seq_decl_plugin.cpp
util/ref_util.h.node: ../src/util/ref_util.h
	@echo done > util/ref_util.h.node
ast/expr2var$(OBJ_EXT): ../src/ast/expr2var.cpp ast/expr2var.h.node ast/ast_smt2_pp.h.node util/ref_util.h.node
	@echo src/ast/expr2var.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr2var$(OBJ_EXT) ../src/ast/expr2var.cpp
ast/macro_substitution.h.node: ../src/ast/macro_substitution.h ast/ast.h.node
	@echo done > ast/macro_substitution.h.node
ast/macro_substitution$(OBJ_EXT): ../src/ast/macro_substitution.cpp ast/macro_substitution.h.node util/ref_util.h.node
	@echo src/ast/macro_substitution.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/macro_substitution$(OBJ_EXT) ../src/ast/macro_substitution.cpp
ast/shared_occs.h.node: ../src/ast/shared_occs.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/shared_occs.h.node
ast/shared_occs$(OBJ_EXT): ../src/ast/shared_occs.cpp ast/shared_occs.h.node ast/ast_smt2_pp.h.node util/ref_util.h.node
	@echo src/ast/shared_occs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/shared_occs$(OBJ_EXT) ../src/ast/shared_occs.cpp
ast/decl_collector$(OBJ_EXT): ../src/ast/decl_collector.cpp ast/decl_collector.h.node ast/ast_pp.h.node
	@echo src/ast/decl_collector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/decl_collector$(OBJ_EXT) ../src/ast/decl_collector.cpp
ast/ast$(OBJ_EXT): ../src/ast/ast.cpp ast/ast.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node util/buffer.h.node util/warning.h.node util/string_buffer.h.node ast/ast_util.h.node ast/ast_smt2_pp.h.node ast/array_decl_plugin.h.node ast/ast_translation.h.node util/z3_version.h.node
	@echo src/ast/ast.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast$(OBJ_EXT) ../src/ast/ast.cpp
ast/expr_delta_pair.h.node: ../src/ast/expr_delta_pair.h
	@echo done > ast/expr_delta_pair.h.node
ast/has_free_vars$(OBJ_EXT): ../src/ast/has_free_vars.cpp ast/ast.h.node ast/expr_delta_pair.h.node util/hashtable.h.node
	@echo src/ast/has_free_vars.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/has_free_vars$(OBJ_EXT) ../src/ast/has_free_vars.cpp
ast/occurs.h.node: ../src/ast/occurs.h
	@echo done > ast/occurs.h.node
ast/occurs$(OBJ_EXT): ../src/ast/occurs.cpp ast/occurs.h.node ast/for_each_expr.h.node
	@echo src/ast/occurs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/occurs$(OBJ_EXT) ../src/ast/occurs.cpp
ast/expr_stat.h.node: ../src/ast/expr_stat.h
	@echo done > ast/expr_stat.h.node
ast/expr_stat$(OBJ_EXT): ../src/ast/expr_stat.cpp ast/for_each_expr.h.node ast/expr_stat.h.node
	@echo src/ast/expr_stat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr_stat$(OBJ_EXT) ../src/ast/expr_stat.cpp
ast/for_each_ast$(OBJ_EXT): ../src/ast/for_each_ast.cpp ast/for_each_ast.h.node
	@echo src/ast/for_each_ast.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/for_each_ast$(OBJ_EXT) ../src/ast/for_each_ast.cpp
ast/expr_substitution.h.node: ../src/ast/expr_substitution.h ast/ast.h.node
	@echo done > ast/expr_substitution.h.node
ast/expr_substitution$(OBJ_EXT): ../src/ast/expr_substitution.cpp util/ref_util.h.node ast/expr_substitution.h.node ast/ast_pp.h.node
	@echo src/ast/expr_substitution.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr_substitution$(OBJ_EXT) ../src/ast/expr_substitution.cpp
ast/ast_pp_dot.h.node: ../src/ast/ast_pp_dot.h ast/ast_pp.h.node
	@echo done > ast/ast_pp_dot.h.node
ast/ast_pp_dot$(OBJ_EXT): ../src/ast/ast_pp_dot.cpp util/util.h.node util/map.h.node ast/ast_pp_dot.h.node
	@echo src/ast/ast_pp_dot.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_pp_dot$(OBJ_EXT) ../src/ast/ast_pp_dot.cpp
ast/ast_util$(OBJ_EXT): ../src/ast/ast_util.cpp ast/ast_util.h.node ast/ast_pp.h.node ast/for_each_expr.h.node ast/arith_decl_plugin.h.node
	@echo src/ast/ast_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_util$(OBJ_EXT) ../src/ast/ast_util.cpp
ast/special_relations_decl_plugin$(OBJ_EXT): ../src/ast/special_relations_decl_plugin.cpp ast/ast.h.node ast/special_relations_decl_plugin.h.node
	@echo src/ast/special_relations_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/special_relations_decl_plugin$(OBJ_EXT) ../src/ast/special_relations_decl_plugin.cpp
ast/recfun_decl_plugin$(OBJ_EXT): ../src/ast/recfun_decl_plugin.cpp ast/expr_functors.h.node ast/recfun_decl_plugin.h.node ast/ast_pp.h.node ast/for_each_expr.h.node util/scoped_ptr_vector.h.node
	@echo src/ast/recfun_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/recfun_decl_plugin$(OBJ_EXT) ../src/ast/recfun_decl_plugin.cpp
ast/display_dimacs.h.node: ../src/ast/display_dimacs.h ast/ast.h.node
	@echo done > ast/display_dimacs.h.node
ast/display_dimacs$(OBJ_EXT): ../src/ast/display_dimacs.cpp ast/ast.h.node ast/display_dimacs.h.node
	@echo src/ast/display_dimacs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/display_dimacs$(OBJ_EXT) ../src/ast/display_dimacs.cpp
ast/array_decl_plugin$(OBJ_EXT): ../src/ast/array_decl_plugin.cpp ast/array_decl_plugin.h.node util/warning.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/arith_decl_plugin.h.node
	@echo src/ast/array_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/array_decl_plugin$(OBJ_EXT) ../src/ast/array_decl_plugin.cpp
ast/func_decl_dependencies.h.node: ../src/ast/func_decl_dependencies.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/func_decl_dependencies.h.node
ast/func_decl_dependencies$(OBJ_EXT): ../src/ast/func_decl_dependencies.cpp ast/func_decl_dependencies.h.node ast/for_each_expr.h.node ast/ast_util.h.node
	@echo src/ast/func_decl_dependencies.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/func_decl_dependencies$(OBJ_EXT) ../src/ast/func_decl_dependencies.cpp
ast/fpa_decl_plugin$(OBJ_EXT): ../src/ast/fpa_decl_plugin.cpp ast/fpa_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/fpa_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/fpa_decl_plugin$(OBJ_EXT) ../src/ast/fpa_decl_plugin.cpp
ast/ast_translation$(OBJ_EXT): ../src/ast/ast_translation.cpp ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/array_decl_plugin.h.node ast/format.h.node ast/ast_translation.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node
	@echo src/ast/ast_translation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_translation$(OBJ_EXT) ../src/ast/ast_translation.cpp
util/dec_ref_util.h.node: ../src/util/dec_ref_util.h
	@echo done > util/dec_ref_util.h.node
ast/expr_map$(OBJ_EXT): ../src/ast/expr_map.cpp ast/expr_map.h.node util/dec_ref_util.h.node
	@echo src/ast/expr_map.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/expr_map$(OBJ_EXT) ../src/ast/expr_map.cpp
ast/ast_lt$(OBJ_EXT): ../src/ast/ast_lt.cpp ast/ast.h.node
	@echo src/ast/ast_lt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_lt$(OBJ_EXT) ../src/ast/ast_lt.cpp
ast/used_vars.h.node: ../src/ast/used_vars.h ast/ast.h.node ast/expr_delta_pair.h.node
	@echo done > ast/used_vars.h.node
ast/used_vars$(OBJ_EXT): ../src/ast/used_vars.cpp ast/used_vars.h.node
	@echo src/ast/used_vars.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/used_vars$(OBJ_EXT) ../src/ast/used_vars.cpp
ast/ast_smt2_pp$(OBJ_EXT): ../src/ast/ast_smt2_pp.cpp ast/ast_smt2_pp.h.node ast/shared_occs.h.node ast/pp.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node math/polynomial/algebraic_numbers.h.node ast/pp_params.hpp.node
	@echo src/ast/ast_smt2_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_smt2_pp$(OBJ_EXT) ../src/ast/ast_smt2_pp.cpp
ast/act_cache.h.node: ../src/ast/act_cache.h ast/ast.h.node util/obj_hashtable.h.node util/chashtable.h.node
	@echo done > ast/act_cache.h.node
ast/act_cache$(OBJ_EXT): ../src/ast/act_cache.cpp ast/act_cache.h.node
	@echo src/ast/act_cache.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/act_cache$(OBJ_EXT) ../src/ast/act_cache.cpp
ast/reg_decl_plugins.h.node: ../src/ast/reg_decl_plugins.h
	@echo done > ast/reg_decl_plugins.h.node
ast/dl_decl_plugin$(OBJ_EXT): ../src/ast/dl_decl_plugin.cpp ast/ast_pp.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/dl_decl_plugin.h.node util/warning.h.node ast/reg_decl_plugins.h.node
	@echo src/ast/dl_decl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/dl_decl_plugin$(OBJ_EXT) ../src/ast/dl_decl_plugin.cpp
ast/ast_pp_util.h.node: ../src/ast/ast_pp_util.h ast/decl_collector.h.node ast/ast_smt2_pp.h.node util/obj_hashtable.h.node
	@echo done > ast/ast_pp_util.h.node
ast/ast_pp_util$(OBJ_EXT): ../src/ast/ast_pp_util.cpp ast/ast_pp_util.h.node ast/ast_smt2_pp.h.node ast/ast_smt_pp.h.node ast/recfun_decl_plugin.h.node
	@echo src/ast/ast_pp_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/ast_pp_util$(OBJ_EXT) ../src/ast/ast_pp_util.cpp
ast/num_occurs.h.node: ../src/ast/num_occurs.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > ast/num_occurs.h.node
ast/num_occurs$(OBJ_EXT): ../src/ast/num_occurs.cpp ast/num_occurs.h.node
	@echo src/ast/num_occurs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_9) $(CXX_OUT_FLAG)ast/num_occurs$(OBJ_EXT) ../src/ast/num_occurs.cpp
ast/ast$(LIB_EXT): ast/for_each_expr$(OBJ_EXT) ast/expr_abstract$(OBJ_EXT) ast/format$(OBJ_EXT) ast/expr_functors$(OBJ_EXT) ast/bv_decl_plugin$(OBJ_EXT) ast/expr2polynomial$(OBJ_EXT) ast/pb_decl_plugin$(OBJ_EXT) ast/arith_decl_plugin$(OBJ_EXT) ast/well_sorted$(OBJ_EXT) ast/static_features$(OBJ_EXT) ast/reg_decl_plugins$(OBJ_EXT) ast/datatype_decl_plugin$(OBJ_EXT) ast/value_generator$(OBJ_EXT) ast/ast_ll_pp$(OBJ_EXT) ast/ast_printer$(OBJ_EXT) ast/pp$(OBJ_EXT) ast/ast_smt_pp$(OBJ_EXT) ast/seq_decl_plugin$(OBJ_EXT) ast/expr2var$(OBJ_EXT) ast/macro_substitution$(OBJ_EXT) ast/shared_occs$(OBJ_EXT) ast/decl_collector$(OBJ_EXT) ast/ast$(OBJ_EXT) ast/has_free_vars$(OBJ_EXT) ast/occurs$(OBJ_EXT) ast/expr_stat$(OBJ_EXT) ast/for_each_ast$(OBJ_EXT) ast/expr_substitution$(OBJ_EXT) ast/ast_pp_dot$(OBJ_EXT) ast/ast_util$(OBJ_EXT) ast/special_relations_decl_plugin$(OBJ_EXT) ast/recfun_decl_plugin$(OBJ_EXT) ast/display_dimacs$(OBJ_EXT) ast/array_decl_plugin$(OBJ_EXT) ast/func_decl_dependencies$(OBJ_EXT) ast/fpa_decl_plugin$(OBJ_EXT) ast/ast_translation$(OBJ_EXT) ast/expr_map$(OBJ_EXT) ast/ast_lt$(OBJ_EXT) ast/used_vars$(OBJ_EXT) ast/ast_smt2_pp$(OBJ_EXT) ast/act_cache$(OBJ_EXT) ast/dl_decl_plugin$(OBJ_EXT) ast/ast_pp_util$(OBJ_EXT) ast/num_occurs$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/ast$(LIB_EXT) ast/for_each_expr$(OBJ_EXT) ast/expr_abstract$(OBJ_EXT) ast/format$(OBJ_EXT) ast/expr_functors$(OBJ_EXT) ast/bv_decl_plugin$(OBJ_EXT) ast/expr2polynomial$(OBJ_EXT) ast/pb_decl_plugin$(OBJ_EXT) ast/arith_decl_plugin$(OBJ_EXT) ast/well_sorted$(OBJ_EXT) ast/static_features$(OBJ_EXT) ast/reg_decl_plugins$(OBJ_EXT) ast/datatype_decl_plugin$(OBJ_EXT) ast/value_generator$(OBJ_EXT) ast/ast_ll_pp$(OBJ_EXT) ast/ast_printer$(OBJ_EXT) ast/pp$(OBJ_EXT) ast/ast_smt_pp$(OBJ_EXT) ast/seq_decl_plugin$(OBJ_EXT) ast/expr2var$(OBJ_EXT) ast/macro_substitution$(OBJ_EXT) ast/shared_occs$(OBJ_EXT) ast/decl_collector$(OBJ_EXT) ast/ast$(OBJ_EXT) ast/has_free_vars$(OBJ_EXT) ast/occurs$(OBJ_EXT) ast/expr_stat$(OBJ_EXT) ast/for_each_ast$(OBJ_EXT) ast/expr_substitution$(OBJ_EXT) ast/ast_pp_dot$(OBJ_EXT) ast/ast_util$(OBJ_EXT) ast/special_relations_decl_plugin$(OBJ_EXT) ast/recfun_decl_plugin$(OBJ_EXT) ast/display_dimacs$(OBJ_EXT) ast/array_decl_plugin$(OBJ_EXT) ast/func_decl_dependencies$(OBJ_EXT) ast/fpa_decl_plugin$(OBJ_EXT) ast/ast_translation$(OBJ_EXT) ast/expr_map$(OBJ_EXT) ast/ast_lt$(OBJ_EXT) ast/used_vars$(OBJ_EXT) ast/ast_smt2_pp$(OBJ_EXT) ast/act_cache$(OBJ_EXT) ast/dl_decl_plugin$(OBJ_EXT) ast/ast_pp_util$(OBJ_EXT) ast/num_occurs$(OBJ_EXT)
ast: ast/ast$(LIB_EXT)

includes_10 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src
params/pattern_inference_params.h.node: ../src/params/pattern_inference_params.h util/params.h.node
	@echo done > params/pattern_inference_params.h.node
params/pattern_inference_params_helper.hpp.node: ../src/params/pattern_inference_params_helper.hpp util/params.h.node util/gparams.h.node
	@echo done > params/pattern_inference_params_helper.hpp.node
params/pattern_inference_params$(OBJ_EXT): ../src/params/pattern_inference_params.cpp params/pattern_inference_params.h.node params/pattern_inference_params_helper.hpp.node
	@echo src/params/pattern_inference_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_10) $(CXX_OUT_FLAG)params/pattern_inference_params$(OBJ_EXT) ../src/params/pattern_inference_params.cpp
params/context_params.h.node: ../src/params/context_params.h util/params.h.node
	@echo done > params/context_params.h.node
params/context_params$(OBJ_EXT): ../src/params/context_params.cpp util/gparams.h.node util/params.h.node ast/ast.h.node params/context_params.h.node
	@echo src/params/context_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_10) $(CXX_OUT_FLAG)params/context_params$(OBJ_EXT) ../src/params/context_params.cpp
params/params$(LIB_EXT): params/pattern_inference_params$(OBJ_EXT) params/context_params$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)params/params$(LIB_EXT) params/pattern_inference_params$(OBJ_EXT) params/context_params$(OBJ_EXT)
params: params/params$(LIB_EXT)

includes_11 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src
util/id_var_list.h.node: ../src/util/id_var_list.h util/region.h.node util/debug.h.node
	@echo done > util/id_var_list.h.node
ast/euf/euf_justification.h.node: ../src/ast/euf/euf_justification.h
	@echo done > ast/euf/euf_justification.h.node
ast/euf/euf_enode.h.node: ../src/ast/euf/euf_enode.h util/vector.h.node util/id_var_list.h.node util/lbool.h.node ast/ast.h.node ast/euf/euf_justification.h.node
	@echo done > ast/euf/euf_enode.h.node
ast/euf/euf_etable.h.node: ../src/ast/euf/euf_etable.h ast/euf/euf_enode.h.node util/hashtable.h.node util/chashtable.h.node
	@echo done > ast/euf/euf_etable.h.node
ast/euf/euf_egraph.h.node: ../src/ast/euf/euf_egraph.h util/statistics.h.node util/trail.h.node util/lbool.h.node ast/euf/euf_enode.h.node ast/euf/euf_etable.h.node ast/ast_ll_pp.h.node
	@echo done > ast/euf/euf_egraph.h.node
ast/euf/euf_egraph$(OBJ_EXT): ../src/ast/euf/euf_egraph.cpp ast/euf/euf_egraph.h.node ast/ast_pp.h.node ast/ast_translation.h.node
	@echo src/ast/euf/euf_egraph.cpp
	@$(CXX) $(CXXFLAGS) $(includes_11) $(CXX_OUT_FLAG)ast/euf/euf_egraph$(OBJ_EXT) ../src/ast/euf/euf_egraph.cpp
ast/euf/euf_etable$(OBJ_EXT): ../src/ast/euf/euf_etable.cpp ast/euf/euf_etable.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/ast/euf/euf_etable.cpp
	@$(CXX) $(CXXFLAGS) $(includes_11) $(CXX_OUT_FLAG)ast/euf/euf_etable$(OBJ_EXT) ../src/ast/euf/euf_etable.cpp
ast/euf/euf_enode$(OBJ_EXT): ../src/ast/euf/euf_enode.cpp ast/euf/euf_enode.h.node ast/euf/euf_egraph.h.node
	@echo src/ast/euf/euf_enode.cpp
	@$(CXX) $(CXXFLAGS) $(includes_11) $(CXX_OUT_FLAG)ast/euf/euf_enode$(OBJ_EXT) ../src/ast/euf/euf_enode.cpp
ast/euf/euf$(LIB_EXT): ast/euf/euf_egraph$(OBJ_EXT) ast/euf/euf_etable$(OBJ_EXT) ast/euf/euf_enode$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/euf/euf$(LIB_EXT) ast/euf/euf_egraph$(OBJ_EXT) ast/euf/euf_etable$(OBJ_EXT) ast/euf/euf_enode$(OBJ_EXT)
euf: ast/euf/euf$(LIB_EXT)

includes_12 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src
math/grobner/pdd_solver.h.node: ../src/math/grobner/pdd_solver.h util/dependency.h.node util/obj_hashtable.h.node util/region.h.node util/rlimit.h.node util/statistics.h.node math/dd/dd_pdd.h.node
	@echo done > math/grobner/pdd_solver.h.node
math/grobner/pdd_simplifier.h.node: ../src/math/grobner/pdd_simplifier.h util/uint_set.h.node math/grobner/pdd_solver.h.node
	@echo done > math/grobner/pdd_simplifier.h.node
math/grobner/pdd_solver$(OBJ_EXT): ../src/math/grobner/pdd_solver.cpp math/grobner/pdd_solver.h.node math/grobner/pdd_simplifier.h.node util/uint_set.h.node
	@echo src/math/grobner/pdd_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_12) $(CXX_OUT_FLAG)math/grobner/pdd_solver$(OBJ_EXT) ../src/math/grobner/pdd_solver.cpp
math/grobner/pdd_simplifier$(OBJ_EXT): ../src/math/grobner/pdd_simplifier.cpp math/grobner/pdd_simplifier.h.node math/simplex/bit_matrix.h.node
	@echo src/math/grobner/pdd_simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_12) $(CXX_OUT_FLAG)math/grobner/pdd_simplifier$(OBJ_EXT) ../src/math/grobner/pdd_simplifier.cpp
math/grobner/grobner.h.node: ../src/math/grobner/grobner.h ast/ast.h.node ast/arith_decl_plugin.h.node util/heap.h.node util/obj_hashtable.h.node util/region.h.node util/dependency.h.node
	@echo done > math/grobner/grobner.h.node
math/grobner/grobner$(OBJ_EXT): ../src/math/grobner/grobner.cpp math/grobner/grobner.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node util/ref_util.h.node
	@echo src/math/grobner/grobner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_12) $(CXX_OUT_FLAG)math/grobner/grobner$(OBJ_EXT) ../src/math/grobner/grobner.cpp
math/grobner/grobner$(LIB_EXT): math/grobner/pdd_solver$(OBJ_EXT) math/grobner/pdd_simplifier$(OBJ_EXT) math/grobner/grobner$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/grobner/grobner$(LIB_EXT) math/grobner/pdd_solver$(OBJ_EXT) math/grobner/pdd_simplifier$(OBJ_EXT) math/grobner/grobner$(OBJ_EXT)
grobner: math/grobner/grobner$(LIB_EXT)

includes_13 = -I../src/util -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/ast -I../src/math/polynomial -I../src/math/simplex -I../src
sat/sat_types.h.node: ../src/sat/sat_types.h util/debug.h.node util/approx_set.h.node util/lbool.h.node util/z3_exception.h.node util/common_msgs.h.node util/vector.h.node util/uint_set.h.node util/stopwatch.h.node util/symbol.h.node
	@echo done > sat/sat_types.h.node
sat/sat_allocator.h.node: ../src/sat/sat_allocator.h util/vector.h.node util/machine.h.node
	@echo done > sat/sat_allocator.h.node
sat/sat_clause.h.node: ../src/sat/sat_clause.h util/small_object_allocator.h.node util/id_gen.h.node util/map.h.node sat/sat_types.h.node sat/sat_allocator.h.node
	@echo done > sat/sat_clause.h.node
sat/sat_binspr.h.node: ../src/sat/sat_binspr.h util/params.h.node util/statistics.h.node sat/sat_clause.h.node sat/sat_types.h.node
	@echo done > sat/sat_binspr.h.node
sat/sat_watched.h.node: ../src/sat/sat_watched.h sat/sat_types.h.node util/vector.h.node
	@echo done > sat/sat_watched.h.node
sat/sat_justification.h.node: ../src/sat/sat_justification.h
	@echo done > sat/sat_justification.h.node
sat/sat_var_queue.h.node: ../src/sat/sat_var_queue.h util/heap.h.node sat/sat_types.h.node
	@echo done > sat/sat_var_queue.h.node
sat/sat_extension.h.node: ../src/sat/sat_extension.h sat/sat_types.h.node util/params.h.node util/statistics.h.node
	@echo done > sat/sat_extension.h.node
sat/sat_config.h.node: ../src/sat/sat_config.h util/params.h.node
	@echo done > sat/sat_config.h.node
sat/sat_cleaner.h.node: ../src/sat/sat_cleaner.h sat/sat_types.h.node util/statistics.h.node
	@echo done > sat/sat_cleaner.h.node
sat/sat_clause_set.h.node: ../src/sat/sat_clause_set.h sat/sat_clause.h.node
	@echo done > sat/sat_clause_set.h.node
sat/sat_clause_use_list.h.node: ../src/sat/sat_clause_use_list.h sat/sat_types.h.node util/trace.h.node
	@echo done > sat/sat_clause_use_list.h.node
sat/sat_model_converter.h.node: ../src/sat/sat_model_converter.h sat/sat_types.h.node util/ref_vector.h.node
	@echo done > sat/sat_model_converter.h.node
sat/sat_simplifier.h.node: ../src/sat/sat_simplifier.h sat/sat_types.h.node sat/sat_clause.h.node sat/sat_clause_set.h.node sat/sat_clause_use_list.h.node sat/sat_extension.h.node sat/sat_watched.h.node sat/sat_model_converter.h.node util/heap.h.node util/statistics.h.node util/params.h.node
	@echo done > sat/sat_simplifier.h.node
sat/sat_big.h.node: ../src/sat/sat_big.h sat/sat_types.h.node util/statistics.h.node util/params.h.node
	@echo done > sat/sat_big.h.node
sat/sat_scc.h.node: ../src/sat/sat_scc.h util/statistics.h.node util/params.h.node sat/sat_types.h.node sat/sat_big.h.node
	@echo done > sat/sat_scc.h.node
sat/sat_asymm_branch.h.node: ../src/sat/sat_asymm_branch.h sat/sat_types.h.node sat/sat_big.h.node util/statistics.h.node util/params.h.node
	@echo done > sat/sat_asymm_branch.h.node
sat/sat_aig_finder.h.node: ../src/sat/sat_aig_finder.h util/params.h.node util/statistics.h.node sat/sat_clause.h.node sat/sat_types.h.node sat/sat_big.h.node
	@echo done > sat/sat_aig_finder.h.node
sat/sat_cutset.h.node: ../src/sat/sat_cutset.h util/region.h.node util/debug.h.node util/util.h.node util/lbool.h.node util/vector.h.node
	@echo done > sat/sat_cutset.h.node
sat/sat_aig_cuts.h.node: ../src/sat/sat_aig_cuts.h sat/sat_cutset.h.node sat/sat_types.h.node
	@echo done > sat/sat_aig_cuts.h.node
sat/sat_cut_simplifier.h.node: ../src/sat/sat_cut_simplifier.h util/union_find.h.node sat/sat_aig_finder.h.node sat/sat_aig_cuts.h.node
	@echo done > sat/sat_cut_simplifier.h.node
sat/sat_probing.h.node: ../src/sat/sat_probing.h sat/sat_types.h.node sat/sat_big.h.node util/params.h.node util/statistics.h.node
	@echo done > sat/sat_probing.h.node
sat/sat_mus.h.node: ../src/sat/sat_mus.h
	@echo done > sat/sat_mus.h.node
sat/sat_drat.h.node: ../src/sat/sat_drat.h sat/sat_types.h.node
	@echo done > sat/sat_drat.h.node
sat/sat_parallel.h.node: ../src/sat/sat_parallel.h sat/sat_types.h.node util/hashtable.h.node util/map.h.node util/rlimit.h.node util/scoped_ptr_vector.h.node util/mutex.h.node
	@echo done > sat/sat_parallel.h.node
util/ema.h.node: ../src/util/ema.h
	@echo done > util/ema.h.node
sat/sat_local_search.h.node: ../src/sat/sat_local_search.h util/vector.h.node sat/sat_types.h.node sat/sat_config.h.node util/rlimit.h.node util/ema.h.node util/statistics.h.node
	@echo done > sat/sat_local_search.h.node
sat/sat_solver_core.h.node: ../src/sat/sat_solver_core.h sat/sat_types.h.node
	@echo done > sat/sat_solver_core.h.node
util/scoped_limit_trail.h.node: ../src/util/scoped_limit_trail.h util/vector.h.node
	@echo done > util/scoped_limit_trail.h.node
sat/sat_solver.h.node: ../src/sat/sat_solver.h sat/sat_types.h.node sat/sat_clause.h.node sat/sat_watched.h.node sat/sat_justification.h.node sat/sat_var_queue.h.node sat/sat_extension.h.node sat/sat_config.h.node sat/sat_cleaner.h.node sat/sat_simplifier.h.node sat/sat_scc.h.node sat/sat_asymm_branch.h.node sat/sat_cut_simplifier.h.node sat/sat_probing.h.node sat/sat_mus.h.node sat/sat_binspr.h.node sat/sat_drat.h.node sat/sat_parallel.h.node sat/sat_local_search.h.node sat/sat_solver_core.h.node util/params.h.node util/statistics.h.node util/stopwatch.h.node util/ema.h.node util/trace.h.node util/rlimit.h.node util/scoped_ptr_vector.h.node util/scoped_limit_trail.h.node
	@echo done > sat/sat_solver.h.node
sat/sat_binspr$(OBJ_EXT): ../src/sat/sat_binspr.cpp sat/sat_binspr.h.node sat/sat_solver.h.node sat/sat_big.h.node
	@echo src/sat/sat_binspr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_binspr$(OBJ_EXT) ../src/sat/sat_binspr.cpp
sat/sat_gc$(OBJ_EXT): ../src/sat/sat_gc.cpp sat/sat_solver.h.node
	@echo src/sat/sat_gc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_gc$(OBJ_EXT) ../src/sat/sat_gc.cpp
sat/sat_cleaner$(OBJ_EXT): ../src/sat/sat_cleaner.cpp sat/sat_cleaner.h.node sat/sat_solver.h.node util/trace.h.node util/stopwatch.h.node
	@echo src/sat/sat_cleaner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_cleaner$(OBJ_EXT) ../src/sat/sat_cleaner.cpp
sat/sat_prob.h.node: ../src/sat/sat_prob.h util/uint_set.h.node util/rlimit.h.node sat/sat_clause.h.node sat/sat_types.h.node
	@echo done > sat/sat_prob.h.node
util/luby.h.node: ../src/util/luby.h
	@echo done > util/luby.h.node
sat/sat_prob$(OBJ_EXT): ../src/sat/sat_prob.cpp sat/sat_prob.h.node sat/sat_solver.h.node util/luby.h.node
	@echo src/sat/sat_prob.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_prob$(OBJ_EXT) ../src/sat/sat_prob.cpp
sat/sat_elim_eqs.h.node: ../src/sat/sat_elim_eqs.h sat/sat_types.h.node util/union_find.h.node
	@echo done > sat/sat_elim_eqs.h.node
sat/sat_scc_params.hpp.node: ../src/sat/sat_scc_params.hpp util/params.h.node util/gparams.h.node
	@echo done > sat/sat_scc_params.hpp.node
sat/sat_scc$(OBJ_EXT): ../src/sat/sat_scc.cpp sat/sat_scc.h.node sat/sat_solver.h.node sat/sat_elim_eqs.h.node util/stopwatch.h.node util/trace.h.node sat/sat_scc_params.hpp.node
	@echo src/sat/sat_scc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_scc$(OBJ_EXT) ../src/sat/sat_scc.cpp
sat/sat_big$(OBJ_EXT): ../src/sat/sat_big.cpp sat/sat_big.h.node sat/sat_solver.h.node
	@echo src/sat/sat_big.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_big$(OBJ_EXT) ../src/sat/sat_big.cpp
sat/sat_clause_set$(OBJ_EXT): ../src/sat/sat_clause_set.cpp sat/sat_clause_set.h.node
	@echo src/sat/sat_clause_set.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_clause_set$(OBJ_EXT) ../src/sat/sat_clause_set.cpp
sat/sat_xor_finder.h.node: ../src/sat/sat_xor_finder.h util/params.h.node util/statistics.h.node sat/sat_clause.h.node sat/sat_types.h.node sat/sat_solver.h.node
	@echo done > sat/sat_xor_finder.h.node
sat/sat_xor_finder$(OBJ_EXT): ../src/sat/sat_xor_finder.cpp sat/sat_xor_finder.h.node sat/sat_solver.h.node
	@echo src/sat/sat_xor_finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_xor_finder$(OBJ_EXT) ../src/sat/sat_xor_finder.cpp
sat/sat_params.hpp.node: ../src/sat/sat_params.hpp util/params.h.node util/gparams.h.node
	@echo done > sat/sat_params.hpp.node
sat/sat_simplifier_params.hpp.node: ../src/sat/sat_simplifier_params.hpp util/params.h.node util/gparams.h.node
	@echo done > sat/sat_simplifier_params.hpp.node
sat/sat_config$(OBJ_EXT): ../src/sat/sat_config.cpp sat/sat_config.h.node sat/sat_types.h.node sat/sat_params.hpp.node sat/sat_simplifier_params.hpp.node
	@echo src/sat/sat_config.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_config$(OBJ_EXT) ../src/sat/sat_config.cpp
sat/sat_aig_finder$(OBJ_EXT): ../src/sat/sat_aig_finder.cpp sat/sat_aig_finder.h.node sat/sat_solver.h.node
	@echo src/sat/sat_aig_finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_aig_finder$(OBJ_EXT) ../src/sat/sat_aig_finder.cpp
sat/sat_anf_simplifier.h.node: ../src/sat/sat_anf_simplifier.h util/params.h.node util/statistics.h.node sat/sat_clause.h.node sat/sat_types.h.node sat/sat_solver.h.node
	@echo done > sat/sat_anf_simplifier.h.node
sat/sat_anf_simplifier$(OBJ_EXT): ../src/sat/sat_anf_simplifier.cpp util/union_find.h.node sat/sat_anf_simplifier.h.node sat/sat_solver.h.node sat/sat_elim_eqs.h.node sat/sat_xor_finder.h.node sat/sat_aig_finder.h.node math/grobner/pdd_solver.h.node
	@echo src/sat/sat_anf_simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_anf_simplifier$(OBJ_EXT) ../src/sat/sat_anf_simplifier.cpp
sat/sat_integrity_checker.h.node: ../src/sat/sat_integrity_checker.h sat/sat_types.h.node sat/sat_watched.h.node
	@echo done > sat/sat_integrity_checker.h.node
sat/sat_integrity_checker$(OBJ_EXT): ../src/sat/sat_integrity_checker.cpp sat/sat_integrity_checker.h.node sat/sat_solver.h.node util/trace.h.node
	@echo src/sat/sat_integrity_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_integrity_checker$(OBJ_EXT) ../src/sat/sat_integrity_checker.cpp
sat/sat_cutset_compute_shift.h.node: ../src/sat/sat_cutset_compute_shift.h
	@echo done > sat/sat_cutset_compute_shift.h.node
sat/sat_cutset$(OBJ_EXT): ../src/sat/sat_cutset.cpp util/hashtable.h.node sat/sat_cutset.h.node sat/sat_cutset_compute_shift.h.node
	@echo src/sat/sat_cutset.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_cutset$(OBJ_EXT) ../src/sat/sat_cutset.cpp
sat/sat_elim_vars.h.node: ../src/sat/sat_elim_vars.h sat/sat_types.h.node math/dd/dd_bdd.h.node
	@echo done > sat/sat_elim_vars.h.node
sat/sat_elim_vars$(OBJ_EXT): ../src/sat/sat_elim_vars.cpp sat/sat_simplifier.h.node sat/sat_elim_vars.h.node sat/sat_solver.h.node
	@echo src/sat/sat_elim_vars.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_elim_vars$(OBJ_EXT) ../src/sat/sat_elim_vars.cpp
sat/sat_ddfw.h.node: ../src/sat/sat_ddfw.h util/uint_set.h.node util/rlimit.h.node util/params.h.node util/ema.h.node sat/sat_clause.h.node sat/sat_types.h.node
	@echo done > sat/sat_ddfw.h.node
sat/sat_ddfw$(OBJ_EXT): ../src/sat/sat_ddfw.cpp util/luby.h.node sat/sat_ddfw.h.node sat/sat_solver.h.node sat/sat_params.hpp.node
	@echo src/sat/sat_ddfw.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_ddfw$(OBJ_EXT) ../src/sat/sat_ddfw.cpp
sat/sat_bcd.h.node: ../src/sat/sat_bcd.h sat/sat_types.h.node sat/sat_simplifier.h.node util/union_find.h.node
	@echo done > sat/sat_bcd.h.node
sat/sat_bcd$(OBJ_EXT): ../src/sat/sat_bcd.cpp util/trace.h.node util/bit_vector.h.node util/map.h.node sat/sat_bcd.h.node sat/sat_solver.h.node
	@echo src/sat/sat_bcd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_bcd$(OBJ_EXT) ../src/sat/sat_bcd.cpp
sat/sat_mus$(OBJ_EXT): ../src/sat/sat_mus.cpp sat/sat_solver.h.node sat/sat_mus.h.node
	@echo src/sat/sat_mus.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_mus$(OBJ_EXT) ../src/sat/sat_mus.cpp
sat/dimacs.h.node: ../src/sat/dimacs.h sat/sat_types.h.node
	@echo done > sat/dimacs.h.node
sat/dimacs$(OBJ_EXT): ../src/sat/dimacs.cpp sat/dimacs.h.node sat/sat_solver.h.node
	@echo src/sat/dimacs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/dimacs$(OBJ_EXT) ../src/sat/dimacs.cpp
sat/sat_simplifier$(OBJ_EXT): ../src/sat/sat_simplifier.cpp sat/sat_simplifier.h.node sat/sat_simplifier_params.hpp.node sat/sat_solver.h.node sat/sat_elim_vars.h.node sat/sat_integrity_checker.h.node util/stopwatch.h.node util/trace.h.node
	@echo src/sat/sat_simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_simplifier$(OBJ_EXT) ../src/sat/sat_simplifier.cpp
sat/sat_parallel$(OBJ_EXT): ../src/sat/sat_parallel.cpp sat/sat_parallel.h.node sat/sat_clause.h.node sat/sat_solver.h.node
	@echo src/sat/sat_parallel.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_parallel$(OBJ_EXT) ../src/sat/sat_parallel.cpp
util/max_cliques.h.node: ../src/util/max_cliques.h util/vector.h.node util/uint_set.h.node
	@echo done > util/max_cliques.h.node
sat/sat_lookahead.h.node: ../src/sat/sat_lookahead.h util/small_object_allocator.h.node sat/sat_elim_eqs.h.node
	@echo done > sat/sat_lookahead.h.node
sat/sat_solver$(OBJ_EXT): ../src/sat/sat_solver.cpp util/luby.h.node util/trace.h.node util/max_cliques.h.node util/gparams.h.node sat/sat_solver.h.node sat/sat_integrity_checker.h.node sat/sat_lookahead.h.node sat/sat_ddfw.h.node sat/sat_prob.h.node sat/sat_anf_simplifier.h.node sat/sat_cut_simplifier.h.node
	@echo src/sat/sat_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_solver$(OBJ_EXT) ../src/sat/sat_solver.cpp
sat/sat_model_converter$(OBJ_EXT): ../src/sat/sat_model_converter.cpp sat/sat_model_converter.h.node sat/sat_clause.h.node sat/sat_solver.h.node util/trace.h.node
	@echo src/sat/sat_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_model_converter$(OBJ_EXT) ../src/sat/sat_model_converter.cpp
util/timer.h.node: ../src/util/timer.h util/stopwatch.h.node
	@echo done > util/timer.h.node
sat/sat_local_search$(OBJ_EXT): ../src/sat/sat_local_search.cpp sat/sat_local_search.h.node sat/sat_solver.h.node sat/sat_params.hpp.node util/timer.h.node
	@echo src/sat/sat_local_search.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_local_search$(OBJ_EXT) ../src/sat/sat_local_search.cpp
sat/sat_elim_eqs$(OBJ_EXT): ../src/sat/sat_elim_eqs.cpp sat/sat_elim_eqs.h.node sat/sat_solver.h.node util/trace.h.node
	@echo src/sat/sat_elim_eqs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_elim_eqs$(OBJ_EXT) ../src/sat/sat_elim_eqs.cpp
sat/sat_npn3_finder.h.node: ../src/sat/sat_npn3_finder.h util/params.h.node util/scoped_ptr_vector.h.node util/statistics.h.node sat/sat_clause.h.node sat/sat_types.h.node sat/sat_big.h.node
	@echo done > sat/sat_npn3_finder.h.node
sat/sat_npn3_finder$(OBJ_EXT): ../src/sat/sat_npn3_finder.cpp sat/sat_npn3_finder.h.node sat/sat_solver.h.node
	@echo src/sat/sat_npn3_finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_npn3_finder$(OBJ_EXT) ../src/sat/sat_npn3_finder.cpp
sat/sat_probing$(OBJ_EXT): ../src/sat/sat_probing.cpp sat/sat_probing.h.node sat/sat_solver.h.node sat/sat_elim_eqs.h.node sat/sat_simplifier_params.hpp.node
	@echo src/sat/sat_probing.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_probing$(OBJ_EXT) ../src/sat/sat_probing.cpp
sat/sat_watched$(OBJ_EXT): ../src/sat/sat_watched.cpp sat/sat_watched.h.node sat/sat_clause.h.node sat/sat_extension.h.node
	@echo src/sat/sat_watched.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_watched$(OBJ_EXT) ../src/sat/sat_watched.cpp
sat/sat_clause_use_list$(OBJ_EXT): ../src/sat/sat_clause_use_list.cpp sat/sat_clause.h.node sat/sat_clause_use_list.h.node
	@echo src/sat/sat_clause_use_list.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_clause_use_list$(OBJ_EXT) ../src/sat/sat_clause_use_list.cpp
sat/sat_lut_finder.h.node: ../src/sat/sat_lut_finder.h util/params.h.node util/statistics.h.node sat/sat_clause.h.node sat/sat_types.h.node sat/sat_solver.h.node
	@echo done > sat/sat_lut_finder.h.node
sat/sat_cut_simplifier$(OBJ_EXT): ../src/sat/sat_cut_simplifier.cpp sat/sat_cut_simplifier.h.node sat/sat_xor_finder.h.node sat/sat_lut_finder.h.node sat/sat_npn3_finder.h.node sat/sat_elim_eqs.h.node
	@echo src/sat/sat_cut_simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_cut_simplifier$(OBJ_EXT) ../src/sat/sat_cut_simplifier.cpp
sat/sat_asymm_branch_params.hpp.node: ../src/sat/sat_asymm_branch_params.hpp util/params.h.node util/gparams.h.node
	@echo done > sat/sat_asymm_branch_params.hpp.node
sat/sat_asymm_branch$(OBJ_EXT): ../src/sat/sat_asymm_branch.cpp sat/sat_asymm_branch.h.node sat/sat_asymm_branch_params.hpp.node sat/sat_solver.h.node sat/sat_big.h.node util/stopwatch.h.node util/trace.h.node
	@echo src/sat/sat_asymm_branch.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_asymm_branch$(OBJ_EXT) ../src/sat/sat_asymm_branch.cpp
sat/sat_drat$(OBJ_EXT): ../src/sat/sat_drat.cpp sat/sat_solver.h.node sat/sat_drat.h.node
	@echo src/sat/sat_drat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_drat$(OBJ_EXT) ../src/sat/sat_drat.cpp
sat/sat_lookahead$(OBJ_EXT): ../src/sat/sat_lookahead.cpp sat/sat_solver.h.node sat/sat_lookahead.h.node sat/sat_scc.h.node util/union_find.h.node
	@echo src/sat/sat_lookahead.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_lookahead$(OBJ_EXT) ../src/sat/sat_lookahead.cpp
sat/sat_lut_finder$(OBJ_EXT): ../src/sat/sat_lut_finder.cpp sat/sat_lut_finder.h.node sat/sat_solver.h.node
	@echo src/sat/sat_lut_finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_lut_finder$(OBJ_EXT) ../src/sat/sat_lut_finder.cpp
sat/sat_aig_cuts$(OBJ_EXT): ../src/sat/sat_aig_cuts.cpp util/trace.h.node sat/sat_aig_cuts.h.node sat/sat_solver.h.node sat/sat_lut_finder.h.node
	@echo src/sat/sat_aig_cuts.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_aig_cuts$(OBJ_EXT) ../src/sat/sat_aig_cuts.cpp
sat/sat_clause$(OBJ_EXT): ../src/sat/sat_clause.cpp sat/sat_clause.h.node util/z3_exception.h.node util/trace.h.node
	@echo src/sat/sat_clause.cpp
	@$(CXX) $(CXXFLAGS) $(includes_13) $(CXX_OUT_FLAG)sat/sat_clause$(OBJ_EXT) ../src/sat/sat_clause.cpp
sat/sat$(LIB_EXT): sat/sat_binspr$(OBJ_EXT) sat/sat_gc$(OBJ_EXT) sat/sat_cleaner$(OBJ_EXT) sat/sat_prob$(OBJ_EXT) sat/sat_scc$(OBJ_EXT) sat/sat_big$(OBJ_EXT) sat/sat_clause_set$(OBJ_EXT) sat/sat_xor_finder$(OBJ_EXT) sat/sat_config$(OBJ_EXT) sat/sat_aig_finder$(OBJ_EXT) sat/sat_anf_simplifier$(OBJ_EXT) sat/sat_integrity_checker$(OBJ_EXT) sat/sat_cutset$(OBJ_EXT) sat/sat_elim_vars$(OBJ_EXT) sat/sat_ddfw$(OBJ_EXT) sat/sat_bcd$(OBJ_EXT) sat/sat_mus$(OBJ_EXT) sat/dimacs$(OBJ_EXT) sat/sat_simplifier$(OBJ_EXT) sat/sat_parallel$(OBJ_EXT) sat/sat_solver$(OBJ_EXT) sat/sat_model_converter$(OBJ_EXT) sat/sat_local_search$(OBJ_EXT) sat/sat_elim_eqs$(OBJ_EXT) sat/sat_npn3_finder$(OBJ_EXT) sat/sat_probing$(OBJ_EXT) sat/sat_watched$(OBJ_EXT) sat/sat_clause_use_list$(OBJ_EXT) sat/sat_cut_simplifier$(OBJ_EXT) sat/sat_asymm_branch$(OBJ_EXT) sat/sat_drat$(OBJ_EXT) sat/sat_lookahead$(OBJ_EXT) sat/sat_lut_finder$(OBJ_EXT) sat/sat_aig_cuts$(OBJ_EXT) sat/sat_clause$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)sat/sat$(LIB_EXT) sat/sat_binspr$(OBJ_EXT) sat/sat_gc$(OBJ_EXT) sat/sat_cleaner$(OBJ_EXT) sat/sat_prob$(OBJ_EXT) sat/sat_scc$(OBJ_EXT) sat/sat_big$(OBJ_EXT) sat/sat_clause_set$(OBJ_EXT) sat/sat_xor_finder$(OBJ_EXT) sat/sat_config$(OBJ_EXT) sat/sat_aig_finder$(OBJ_EXT) sat/sat_anf_simplifier$(OBJ_EXT) sat/sat_integrity_checker$(OBJ_EXT) sat/sat_cutset$(OBJ_EXT) sat/sat_elim_vars$(OBJ_EXT) sat/sat_ddfw$(OBJ_EXT) sat/sat_bcd$(OBJ_EXT) sat/sat_mus$(OBJ_EXT) sat/dimacs$(OBJ_EXT) sat/sat_simplifier$(OBJ_EXT) sat/sat_parallel$(OBJ_EXT) sat/sat_solver$(OBJ_EXT) sat/sat_model_converter$(OBJ_EXT) sat/sat_local_search$(OBJ_EXT) sat/sat_elim_eqs$(OBJ_EXT) sat/sat_npn3_finder$(OBJ_EXT) sat/sat_probing$(OBJ_EXT) sat/sat_watched$(OBJ_EXT) sat/sat_clause_use_list$(OBJ_EXT) sat/sat_cut_simplifier$(OBJ_EXT) sat/sat_asymm_branch$(OBJ_EXT) sat/sat_drat$(OBJ_EXT) sat/sat_lookahead$(OBJ_EXT) sat/sat_lut_finder$(OBJ_EXT) sat/sat_aig_cuts$(OBJ_EXT) sat/sat_clause$(OBJ_EXT)
sat: sat/sat$(LIB_EXT)

includes_14 = -I../src/math/polynomial -I../src/util -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/ast -I../src/math/simplex -I../src
nlsat/nlsat_types.h.node: ../src/nlsat/nlsat_types.h math/polynomial/polynomial.h.node util/buffer.h.node sat/sat_types.h.node util/z3_exception.h.node
	@echo done > nlsat/nlsat_types.h.node
nlsat/nlsat_clause.h.node: ../src/nlsat/nlsat_clause.h nlsat/nlsat_types.h.node util/vector.h.node
	@echo done > nlsat/nlsat_clause.h.node
nlsat/nlsat_clause$(OBJ_EXT): ../src/nlsat/nlsat_clause.cpp nlsat/nlsat_clause.h.node
	@echo src/nlsat/nlsat_clause.cpp
	@$(CXX) $(CXXFLAGS) $(includes_14) $(CXX_OUT_FLAG)nlsat/nlsat_clause$(OBJ_EXT) ../src/nlsat/nlsat_clause.cpp
nlsat/nlsat_types$(OBJ_EXT): ../src/nlsat/nlsat_types.cpp nlsat/nlsat_types.h.node util/debug.h.node util/hash.h.node math/polynomial/polynomial.h.node
	@echo src/nlsat/nlsat_types.cpp
	@$(CXX) $(CXXFLAGS) $(includes_14) $(CXX_OUT_FLAG)nlsat/nlsat_types$(OBJ_EXT) ../src/nlsat/nlsat_types.cpp
nlsat/nlsat_solver.h.node: ../src/nlsat/nlsat_solver.h nlsat/nlsat_types.h.node util/params.h.node util/statistics.h.node util/rlimit.h.node
	@echo done > nlsat/nlsat_solver.h.node
nlsat/nlsat_assignment.h.node: ../src/nlsat/nlsat_assignment.h nlsat/nlsat_types.h.node math/polynomial/algebraic_numbers.h.node
	@echo done > nlsat/nlsat_assignment.h.node
nlsat/nlsat_justification.h.node: ../src/nlsat/nlsat_justification.h nlsat/nlsat_types.h.node util/tptr.h.node
	@echo done > nlsat/nlsat_justification.h.node
nlsat/nlsat_interval_set.h.node: ../src/nlsat/nlsat_interval_set.h nlsat/nlsat_types.h.node
	@echo done > nlsat/nlsat_interval_set.h.node
nlsat/nlsat_evaluator.h.node: ../src/nlsat/nlsat_evaluator.h nlsat/nlsat_types.h.node nlsat/nlsat_assignment.h.node nlsat/nlsat_interval_set.h.node
	@echo done > nlsat/nlsat_evaluator.h.node
nlsat/nlsat_scoped_literal_vector.h.node: ../src/nlsat/nlsat_scoped_literal_vector.h nlsat/nlsat_solver.h.node
	@echo done > nlsat/nlsat_scoped_literal_vector.h.node
nlsat/nlsat_explain.h.node: ../src/nlsat/nlsat_explain.h nlsat/nlsat_solver.h.node nlsat/nlsat_scoped_literal_vector.h.node math/polynomial/polynomial_cache.h.node math/polynomial/algebraic_numbers.h.node
	@echo done > nlsat/nlsat_explain.h.node
nlsat/nlsat_params.hpp.node: ../src/nlsat/nlsat_params.hpp util/params.h.node util/gparams.h.node
	@echo done > nlsat/nlsat_params.hpp.node
nlsat/nlsat_solver$(OBJ_EXT): ../src/nlsat/nlsat_solver.cpp util/z3_exception.h.node util/chashtable.h.node util/id_gen.h.node util/map.h.node util/dependency.h.node util/permutation.h.node math/polynomial/algebraic_numbers.h.node math/polynomial/polynomial_cache.h.node nlsat/nlsat_solver.h.node nlsat/nlsat_clause.h.node nlsat/nlsat_assignment.h.node nlsat/nlsat_justification.h.node nlsat/nlsat_evaluator.h.node nlsat/nlsat_explain.h.node nlsat/nlsat_params.hpp.node
	@echo src/nlsat/nlsat_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_14) $(CXX_OUT_FLAG)nlsat/nlsat_solver$(OBJ_EXT) ../src/nlsat/nlsat_solver.cpp
nlsat/nlsat_interval_set$(OBJ_EXT): ../src/nlsat/nlsat_interval_set.cpp nlsat/nlsat_interval_set.h.node math/polynomial/algebraic_numbers.h.node util/buffer.h.node
	@echo src/nlsat/nlsat_interval_set.cpp
	@$(CXX) $(CXXFLAGS) $(includes_14) $(CXX_OUT_FLAG)nlsat/nlsat_interval_set$(OBJ_EXT) ../src/nlsat/nlsat_interval_set.cpp
nlsat/nlsat_explain$(OBJ_EXT): ../src/nlsat/nlsat_explain.cpp nlsat/nlsat_explain.h.node nlsat/nlsat_assignment.h.node nlsat/nlsat_evaluator.h.node math/polynomial/algebraic_numbers.h.node util/ref_buffer.h.node
	@echo src/nlsat/nlsat_explain.cpp
	@$(CXX) $(CXXFLAGS) $(includes_14) $(CXX_OUT_FLAG)nlsat/nlsat_explain$(OBJ_EXT) ../src/nlsat/nlsat_explain.cpp
nlsat/nlsat_evaluator$(OBJ_EXT): ../src/nlsat/nlsat_evaluator.cpp nlsat/nlsat_evaluator.h.node nlsat/nlsat_solver.h.node
	@echo src/nlsat/nlsat_evaluator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_14) $(CXX_OUT_FLAG)nlsat/nlsat_evaluator$(OBJ_EXT) ../src/nlsat/nlsat_evaluator.cpp
nlsat/nlsat$(LIB_EXT): nlsat/nlsat_clause$(OBJ_EXT) nlsat/nlsat_types$(OBJ_EXT) nlsat/nlsat_solver$(OBJ_EXT) nlsat/nlsat_interval_set$(OBJ_EXT) nlsat/nlsat_explain$(OBJ_EXT) nlsat/nlsat_evaluator$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)nlsat/nlsat$(LIB_EXT) nlsat/nlsat_clause$(OBJ_EXT) nlsat/nlsat_types$(OBJ_EXT) nlsat/nlsat_solver$(OBJ_EXT) nlsat/nlsat_interval_set$(OBJ_EXT) nlsat/nlsat_explain$(OBJ_EXT) nlsat/nlsat_evaluator$(OBJ_EXT)
nlsat: nlsat/nlsat$(LIB_EXT)

includes_15 = -I../src/util -I../src/nlsat -I../src/math/polynomial -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/ast -I../src/math/simplex -I../src
util/sstream.h.node: ../src/util/sstream.h
	@echo done > util/sstream.h.node
math/lp/numeric_pair.h.node: ../src/math/lp/numeric_pair.h util/rational.h.node util/sstream.h.node util/z3_exception.h.node
	@echo done > math/lp/numeric_pair.h.node
math/lp/lp_types.h.node: ../src/math/lp/lp_types.h
	@echo done > math/lp/lp_types.h.node
math/lp/lp_utils.h.node: ../src/math/lp/lp_utils.h math/lp/numeric_pair.h.node math/lp/lp_types.h.node util/debug.h.node
	@echo done > math/lp/lp_utils.h.node
math/lp/lp_settings.h.node: ../src/math/lp/lp_settings.h util/vector.h.node math/lp/lp_utils.h.node util/stopwatch.h.node util/statistics.h.node math/lp/lp_types.h.node
	@echo done > math/lp/lp_settings.h.node
math/lp/sparse_vector.h.node: ../src/math/lp/sparse_vector.h util/vector.h.node util/debug.h.node math/lp/lp_utils.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/sparse_vector.h.node
math/lp/indexed_vector.h.node: ../src/math/lp/indexed_vector.h util/vector.h.node util/debug.h.node math/lp/lp_utils.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/indexed_vector.h.node
math/lp/matrix.h.node: ../src/math/lp/matrix.h math/lp/numeric_pair.h.node util/vector.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/matrix.h.node
math/lp/tail_matrix.h.node: ../src/math/lp/tail_matrix.h util/vector.h.node math/lp/indexed_vector.h.node math/lp/matrix.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/tail_matrix.h.node
math/lp/permutation_matrix.h.node: ../src/math/lp/permutation_matrix.h util/vector.h.node util/debug.h.node math/lp/sparse_vector.h.node math/lp/indexed_vector.h.node math/lp/lp_settings.h.node math/lp/matrix.h.node math/lp/tail_matrix.h.node
	@echo done > math/lp/permutation_matrix.h.node
math/lp/static_matrix.h.node: ../src/math/lp/static_matrix.h util/vector.h.node math/lp/sparse_vector.h.node math/lp/indexed_vector.h.node math/lp/permutation_matrix.h.node
	@echo done > math/lp/static_matrix.h.node
math/lp/u_set.h.node: ../src/math/lp/u_set.h util/vector.h.node
	@echo done > math/lp/u_set.h.node
math/lp/lar_term.h.node: ../src/math/lp/lar_term.h math/lp/indexed_vector.h.node util/map.h.node
	@echo done > math/lp/lar_term.h.node
util/stacked_value.h.node: ../src/util/stacked_value.h
	@echo done > util/stacked_value.h.node
math/lp/column_info.h.node: ../src/math/lp/column_info.h util/vector.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/column_info.h.node
math/lp/ul_pair.h.node: ../src/math/lp/ul_pair.h util/vector.h.node math/lp/column_info.h.node math/lp/lp_types.h.node
	@echo done > math/lp/ul_pair.h.node
math/lp/column_namer.h.node: ../src/math/lp/column_namer.h math/lp/static_matrix.h.node
	@echo done > math/lp/column_namer.h.node
math/lp/lar_constraints.h.node: ../src/math/lp/lar_constraints.h util/vector.h.node util/region.h.node util/stacked_value.h.node math/lp/lp_utils.h.node math/lp/ul_pair.h.node math/lp/lar_term.h.node math/lp/column_namer.h.node
	@echo done > math/lp/lar_constraints.h.node
util/ext_gcd.h.node: ../src/util/ext_gcd.h
	@echo done > util/ext_gcd.h.node
math/lp/hnf.h.node: ../src/math/lp/hnf.h math/lp/numeric_pair.h.node util/ext_gcd.h.node
	@echo done > math/lp/hnf.h.node
math/lp/general_matrix.h.node: ../src/math/lp/general_matrix.h
	@echo done > math/lp/general_matrix.h.node
math/lp/var_register.h.node: ../src/math/lp/var_register.h math/lp/lp_types.h.node
	@echo done > math/lp/var_register.h.node
math/lp/lia_move.h.node: ../src/math/lp/lia_move.h
	@echo done > math/lp/lia_move.h.node
math/lp/explanation.h.node: ../src/math/lp/explanation.h math/lp/lp_utils.h.node util/map.h.node util/hashtable.h.node
	@echo done > math/lp/explanation.h.node
math/lp/hnf_cutter.h.node: ../src/math/lp/hnf_cutter.h math/lp/lar_term.h.node math/lp/hnf.h.node math/lp/general_matrix.h.node math/lp/var_register.h.node math/lp/lia_move.h.node math/lp/explanation.h.node
	@echo done > math/lp/hnf_cutter.h.node
math/lp/int_gcd_test.h.node: ../src/math/lp/int_gcd_test.h math/lp/lia_move.h.node
	@echo done > math/lp/int_gcd_test.h.node
math/lp/int_solver.h.node: ../src/math/lp/int_solver.h math/lp/lp_settings.h.node math/lp/static_matrix.h.node math/lp/u_set.h.node math/lp/lar_term.h.node math/lp/lar_constraints.h.node math/lp/hnf_cutter.h.node math/lp/int_gcd_test.h.node math/lp/lia_move.h.node math/lp/explanation.h.node
	@echo done > math/lp/int_solver.h.node
math/lp/core_solver_pretty_printer.h.node: ../src/math/lp/core_solver_pretty_printer.h util/vector.h.node math/lp/lp_settings.h.node math/lp/indexed_vector.h.node
	@echo done > math/lp/core_solver_pretty_printer.h.node
math/lp/indexed_value.h.node: ../src/math/lp/indexed_value.h
	@echo done > math/lp/indexed_value.h.node
math/lp/eta_matrix.h.node: ../src/math/lp/eta_matrix.h util/vector.h.node math/lp/tail_matrix.h.node math/lp/permutation_matrix.h.node
	@echo done > math/lp/eta_matrix.h.node
math/lp/binary_heap_priority_queue.h.node: ../src/math/lp/binary_heap_priority_queue.h util/vector.h.node util/debug.h.node math/lp/lp_utils.h.node
	@echo done > math/lp/binary_heap_priority_queue.h.node
math/lp/binary_heap_upair_queue.h.node: ../src/math/lp/binary_heap_upair_queue.h util/vector.h.node math/lp/binary_heap_priority_queue.h.node
	@echo done > math/lp/binary_heap_upair_queue.h.node
math/lp/square_sparse_matrix.h.node: ../src/math/lp/square_sparse_matrix.h util/vector.h.node math/lp/permutation_matrix.h.node math/lp/static_matrix.h.node math/lp/indexed_value.h.node math/lp/indexed_vector.h.node math/lp/lp_settings.h.node math/lp/eta_matrix.h.node math/lp/binary_heap_upair_queue.h.node math/lp/numeric_pair.h.node math/lp/u_set.h.node
	@echo done > math/lp/square_sparse_matrix.h.node
math/lp/row_eta_matrix.h.node: ../src/math/lp/row_eta_matrix.h util/vector.h.node util/debug.h.node math/lp/sparse_vector.h.node math/lp/indexed_vector.h.node math/lp/permutation_matrix.h.node
	@echo done > math/lp/row_eta_matrix.h.node
math/lp/square_dense_submatrix.h.node: ../src/math/lp/square_dense_submatrix.h util/vector.h.node math/lp/permutation_matrix.h.node math/lp/static_matrix.h.node math/lp/indexed_value.h.node math/lp/indexed_vector.h.node math/lp/lp_settings.h.node math/lp/eta_matrix.h.node math/lp/binary_heap_upair_queue.h.node math/lp/square_sparse_matrix.h.node
	@echo done > math/lp/square_dense_submatrix.h.node
math/lp/dense_matrix.h.node: ../src/math/lp/dense_matrix.h util/vector.h.node math/lp/matrix.h.node
	@echo done > math/lp/dense_matrix.h.node
math/lp/lu.h.node: ../src/math/lp/lu.h util/vector.h.node util/debug.h.node math/lp/square_sparse_matrix.h.node math/lp/static_matrix.h.node math/lp/numeric_pair.h.node math/lp/row_eta_matrix.h.node math/lp/square_dense_submatrix.h.node math/lp/dense_matrix.h.node
	@echo done > math/lp/lu.h.node
math/lp/lp_core_solver_base.h.node: ../src/math/lp/lp_core_solver_base.h util/vector.h.node math/lp/lp_utils.h.node math/lp/core_solver_pretty_printer.h.node math/lp/numeric_pair.h.node math/lp/static_matrix.h.node math/lp/lu.h.node math/lp/permutation_matrix.h.node math/lp/column_namer.h.node
	@echo done > math/lp/lp_core_solver_base.h.node
math/lp/breakpoint.h.node: ../src/math/lp/breakpoint.h
	@echo done > math/lp/breakpoint.h.node
math/lp/scaler.h.node: ../src/math/lp/scaler.h util/vector.h.node math/lp/lp_utils.h.node math/lp/static_matrix.h.node
	@echo done > math/lp/scaler.h.node
math/lp/implied_bound.h.node: ../src/math/lp/implied_bound.h math/lp/lp_settings.h.node math/lp/lar_constraints.h.node
	@echo done > math/lp/implied_bound.h.node
math/lp/test_bound_analyzer.h.node: ../src/math/lp/test_bound_analyzer.h util/vector.h.node math/lp/implied_bound.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/test_bound_analyzer.h.node
math/lp/bound_analyzer_on_row.h.node: ../src/math/lp/bound_analyzer_on_row.h util/vector.h.node math/lp/implied_bound.h.node math/lp/test_bound_analyzer.h.node
	@echo done > math/lp/bound_analyzer_on_row.h.node
math/lp/lp_solver.h.node: ../src/math/lp/lp_solver.h util/vector.h.node math/lp/lp_settings.h.node math/lp/column_info.h.node math/lp/static_matrix.h.node math/lp/lp_core_solver_base.h.node math/lp/scaler.h.node math/lp/bound_analyzer_on_row.h.node
	@echo done > math/lp/lp_solver.h.node
math/lp/lp_primal_core_solver.h.node: ../src/math/lp/lp_primal_core_solver.h util/vector.h.node math/lp/lu.h.node math/lp/lp_solver.h.node math/lp/static_matrix.h.node math/lp/core_solver_pretty_printer.h.node math/lp/lp_core_solver_base.h.node math/lp/breakpoint.h.node math/lp/binary_heap_priority_queue.h.node math/lp/u_set.h.node
	@echo done > math/lp/lp_primal_core_solver.h.node
math/lp/stacked_vector.h.node: ../src/math/lp/stacked_vector.h util/vector.h.node
	@echo done > math/lp/stacked_vector.h.node
math/lp/lar_solution_signature.h.node: ../src/math/lp/lar_solution_signature.h util/vector.h.node util/debug.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/lar_solution_signature.h.node
math/lp/lar_core_solver.h.node: ../src/math/lp/lar_core_solver.h util/vector.h.node math/lp/lp_core_solver_base.h.node math/lp/indexed_vector.h.node math/lp/binary_heap_priority_queue.h.node math/lp/breakpoint.h.node math/lp/lp_primal_core_solver.h.node math/lp/stacked_vector.h.node math/lp/lar_solution_signature.h.node util/stacked_value.h.node
	@echo done > math/lp/lar_core_solver.h.node
math/lp/random_updater.h.node: ../src/math/lp/random_updater.h util/vector.h.node math/lp/lp_settings.h.node math/lp/u_set.h.node
	@echo done > math/lp/random_updater.h.node
math/lp/conversion_helper.h.node: ../src/math/lp/conversion_helper.h
	@echo done > math/lp/conversion_helper.h.node
math/lp/nra_solver.h.node: ../src/math/lp/nra_solver.h util/vector.h.node math/lp/lp_settings.h.node util/rlimit.h.node util/params.h.node nlsat/nlsat_solver.h.node
	@echo done > math/lp/nra_solver.h.node
math/lp/lp_bound_propagator.h.node: ../src/math/lp/lp_bound_propagator.h math/lp/lp_settings.h.node
	@echo done > math/lp/lp_bound_propagator.h.node
math/lp/lar_solver.h.node: ../src/math/lp/lar_solver.h util/vector.h.node util/debug.h.node util/buffer.h.node math/lp/lar_constraints.h.node math/lp/lar_core_solver.h.node math/lp/numeric_pair.h.node math/lp/scaler.h.node math/lp/lp_primal_core_solver.h.node math/lp/random_updater.h.node util/stacked_value.h.node math/lp/stacked_vector.h.node math/lp/implied_bound.h.node math/lp/bound_analyzer_on_row.h.node math/lp/conversion_helper.h.node math/lp/int_solver.h.node math/lp/nra_solver.h.node math/lp/lp_types.h.node math/lp/lp_bound_propagator.h.node
	@echo done > math/lp/lar_solver.h.node
math/lp/int_cube.h.node: ../src/math/lp/int_cube.h math/lp/lia_move.h.node
	@echo done > math/lp/int_cube.h.node
math/lp/int_cube$(OBJ_EXT): ../src/math/lp/int_cube.cpp math/lp/int_solver.h.node math/lp/lar_solver.h.node math/lp/int_cube.h.node
	@echo src/math/lp/int_cube.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/int_cube$(OBJ_EXT) ../src/math/lp/int_cube.cpp
math/lp/lar_solver$(OBJ_EXT): ../src/math/lp/lar_solver.cpp math/lp/lar_solver.h.node
	@echo src/math/lp/lar_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lar_solver$(OBJ_EXT) ../src/math/lp/lar_solver.cpp
math/lp/lp_dual_core_solver.h.node: ../src/math/lp/lp_dual_core_solver.h math/lp/static_matrix.h.node math/lp/lp_core_solver_base.h.node util/vector.h.node
	@echo done > math/lp/lp_dual_core_solver.h.node
math/lp/lp_dual_core_solver_def.h.node: ../src/math/lp/lp_dual_core_solver_def.h util/vector.h.node math/lp/lp_dual_core_solver.h.node
	@echo done > math/lp/lp_dual_core_solver_def.h.node
math/lp/lp_dual_core_solver$(OBJ_EXT): ../src/math/lp/lp_dual_core_solver.cpp util/vector.h.node math/lp/lp_dual_core_solver_def.h.node
	@echo src/math/lp/lp_dual_core_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_dual_core_solver$(OBJ_EXT) ../src/math/lp/lp_dual_core_solver.cpp
math/lp/scaler_def.h.node: ../src/math/lp/scaler_def.h math/lp/scaler.h.node math/lp/numeric_pair.h.node
	@echo done > math/lp/scaler_def.h.node
math/lp/scaler$(OBJ_EXT): ../src/math/lp/scaler.cpp math/lp/scaler_def.h.node
	@echo src/math/lp/scaler.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/scaler$(OBJ_EXT) ../src/math/lp/scaler.cpp
math/lp/indexed_vector_def.h.node: ../src/math/lp/indexed_vector_def.h util/vector.h.node math/lp/indexed_vector.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/indexed_vector_def.h.node
math/lp/indexed_vector$(OBJ_EXT): ../src/math/lp/indexed_vector.cpp util/vector.h.node math/lp/indexed_vector_def.h.node
	@echo src/math/lp/indexed_vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/indexed_vector$(OBJ_EXT) ../src/math/lp/indexed_vector.cpp
math/lp/matrix_def.h.node: ../src/math/lp/matrix_def.h math/lp/matrix.h.node
	@echo done > math/lp/matrix_def.h.node
math/lp/matrix$(OBJ_EXT): ../src/math/lp/matrix.cpp math/lp/lp_settings.h.node math/lp/matrix_def.h.node math/lp/static_matrix.h.node
	@echo src/math/lp/matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/matrix$(OBJ_EXT) ../src/math/lp/matrix.cpp
math/lp/nla_defs.h.node: ../src/math/lp/nla_defs.h math/lp/lp_types.h.node math/lp/column_info.h.node math/lp/explanation.h.node
	@echo done > math/lp/nla_defs.h.node
math/lp/incremental_vector.h.node: ../src/math/lp/incremental_vector.h util/vector.h.node
	@echo done > math/lp/incremental_vector.h.node
math/lp/var_eqs.h.node: ../src/math/lp/var_eqs.h util/union_find.h.node math/lp/nla_defs.h.node util/rational.h.node math/lp/explanation.h.node math/lp/incremental_vector.h.node
	@echo done > math/lp/var_eqs.h.node
math/lp/monic.h.node: ../src/math/lp/monic.h math/lp/lp_settings.h.node util/vector.h.node math/lp/lar_solver.h.node math/lp/nla_defs.h.node
	@echo done > math/lp/monic.h.node
math/lp/emonics.h.node: ../src/math/lp/emonics.h math/lp/lp_utils.h.node math/lp/var_eqs.h.node math/lp/monic.h.node util/region.h.node util/map.h.node
	@echo done > math/lp/emonics.h.node
math/lp/factorization.h.node: ../src/math/lp/factorization.h util/rational.h.node math/lp/monic.h.node math/lp/nla_defs.h.node
	@echo done > math/lp/factorization.h.node
math/lp/nex.h.node: ../src/math/lp/nex.h math/lp/nla_defs.h.node
	@echo done > math/lp/nex.h.node
math/lp/nex_creator.h.node: ../src/math/lp/nex_creator.h util/map.h.node math/lp/nex.h.node
	@echo done > math/lp/nex_creator.h.node
math/lp/nla_common.h.node: ../src/math/lp/nla_common.h util/rational.h.node math/lp/nla_defs.h.node math/lp/lar_term.h.node math/lp/monic.h.node math/lp/emonics.h.node math/lp/factorization.h.node util/dependency.h.node util/region.h.node math/lp/nex_creator.h.node
	@echo done > math/lp/nla_common.h.node
math/lp/nla_tangent_lemmas.h.node: ../src/math/lp/nla_tangent_lemmas.h util/rational.h.node math/lp/factorization.h.node math/lp/nla_common.h.node
	@echo done > math/lp/nla_tangent_lemmas.h.node
math/lp/nla_basics_lemmas.h.node: ../src/math/lp/nla_basics_lemmas.h math/lp/monic.h.node math/lp/factorization.h.node math/lp/nla_common.h.node
	@echo done > math/lp/nla_basics_lemmas.h.node
math/lp/nla_order_lemmas.h.node: ../src/math/lp/nla_order_lemmas.h math/lp/factorization.h.node math/lp/nla_common.h.node
	@echo done > math/lp/nla_order_lemmas.h.node
math/lp/nla_monotone_lemmas.h.node: ../src/math/lp/nla_monotone_lemmas.h
	@echo done > math/lp/nla_monotone_lemmas.h.node
math/lp/nla_settings.h.node: ../src/math/lp/nla_settings.h
	@echo done > math/lp/nla_settings.h.node
math/lp/nla_intervals.h.node: ../src/math/lp/nla_intervals.h util/dependency.h.node util/region.h.node math/lp/nla_common.h.node math/lp/lar_solver.h.node math/interval/interval.h.node math/interval/dep_intervals.h.node util/dependency.h.node
	@echo done > math/lp/nla_intervals.h.node
math/lp/cross_nested.h.node: ../src/math/lp/cross_nested.h math/lp/nex.h.node math/lp/nex_creator.h.node
	@echo done > math/lp/cross_nested.h.node
math/lp/horner.h.node: ../src/math/lp/horner.h math/lp/nla_common.h.node math/lp/nla_intervals.h.node math/lp/nex.h.node math/lp/cross_nested.h.node math/lp/u_set.h.node
	@echo done > math/lp/horner.h.node
math/lp/monomial_bounds.h.node: ../src/math/lp/monomial_bounds.h math/lp/nla_common.h.node math/lp/nla_intervals.h.node math/lp/u_set.h.node
	@echo done > math/lp/monomial_bounds.h.node
math/lp/nla_core.h.node: ../src/math/lp/nla_core.h math/lp/factorization.h.node math/lp/lp_types.h.node math/lp/var_eqs.h.node math/lp/nla_tangent_lemmas.h.node math/lp/nla_basics_lemmas.h.node math/lp/nla_order_lemmas.h.node math/lp/nla_monotone_lemmas.h.node math/lp/emonics.h.node math/lp/nla_settings.h.node math/lp/nex.h.node math/lp/horner.h.node math/lp/monomial_bounds.h.node math/lp/nla_intervals.h.node math/grobner/pdd_solver.h.node nlsat/nlsat_solver.h.node
	@echo done > math/lp/nla_core.h.node
math/lp/emonics$(OBJ_EXT): ../src/math/lp/emonics.cpp math/lp/emonics.h.node math/lp/nla_defs.h.node math/lp/nla_core.h.node
	@echo src/math/lp/emonics.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/emonics$(OBJ_EXT) ../src/math/lp/emonics.cpp
math/lp/int_branch.h.node: ../src/math/lp/int_branch.h math/lp/lar_solver.h.node math/lp/int_solver.h.node math/lp/lia_move.h.node
	@echo done > math/lp/int_branch.h.node
math/lp/int_branch$(OBJ_EXT): ../src/math/lp/int_branch.cpp math/lp/int_solver.h.node math/lp/lar_solver.h.node math/lp/int_branch.h.node
	@echo src/math/lp/int_branch.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/int_branch$(OBJ_EXT) ../src/math/lp/int_branch.cpp
math/lp/factorization_factory_imp.h.node: ../src/math/lp/factorization_factory_imp.h math/lp/factorization.h.node
	@echo done > math/lp/factorization_factory_imp.h.node
math/lp/nla_order_lemmas$(OBJ_EXT): ../src/math/lp/nla_order_lemmas.cpp math/lp/nla_order_lemmas.h.node math/lp/nla_core.h.node math/lp/nla_common.h.node math/lp/factorization_factory_imp.h.node
	@echo src/math/lp/nla_order_lemmas.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_order_lemmas$(OBJ_EXT) ../src/math/lp/nla_order_lemmas.cpp
math/lp/lp_primal_core_solver_def.h.node: ../src/math/lp/lp_primal_core_solver_def.h util/vector.h.node math/lp/lp_primal_core_solver.h.node
	@echo done > math/lp/lp_primal_core_solver_def.h.node
math/lp/lp_primal_core_solver_tableau_def.h.node: ../src/math/lp/lp_primal_core_solver_tableau_def.h math/lp/lp_primal_core_solver.h.node
	@echo done > math/lp/lp_primal_core_solver_tableau_def.h.node
math/lp/lp_primal_core_solver$(OBJ_EXT): ../src/math/lp/lp_primal_core_solver.cpp util/vector.h.node math/lp/lar_solver.h.node math/lp/lp_primal_core_solver_def.h.node math/lp/lp_primal_core_solver_tableau_def.h.node
	@echo src/math/lp/lp_primal_core_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_primal_core_solver$(OBJ_EXT) ../src/math/lp/lp_primal_core_solver.cpp
math/lp/square_dense_submatrix_def.h.node: ../src/math/lp/square_dense_submatrix_def.h util/vector.h.node math/lp/square_dense_submatrix.h.node
	@echo done > math/lp/square_dense_submatrix_def.h.node
math/lp/square_dense_submatrix$(OBJ_EXT): ../src/math/lp/square_dense_submatrix.cpp util/vector.h.node math/lp/square_dense_submatrix_def.h.node
	@echo src/math/lp/square_dense_submatrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/square_dense_submatrix$(OBJ_EXT) ../src/math/lp/square_dense_submatrix.cpp
math/lp/lp_core_solver_base_def.h.node: ../src/math/lp/lp_core_solver_base_def.h util/vector.h.node math/lp/lp_utils.h.node math/lp/lp_core_solver_base.h.node
	@echo done > math/lp/lp_core_solver_base_def.h.node
math/lp/lp_core_solver_base$(OBJ_EXT): ../src/math/lp/lp_core_solver_base.cpp util/vector.h.node math/lp/lp_core_solver_base_def.h.node
	@echo src/math/lp/lp_core_solver_base.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_core_solver_base$(OBJ_EXT) ../src/math/lp/lp_core_solver_base.cpp
math/lp/nla_solver.h.node: ../src/math/lp/nla_solver.h util/vector.h.node math/lp/lp_settings.h.node util/rlimit.h.node util/params.h.node math/lp/lar_solver.h.node math/lp/monic.h.node math/lp/nla_settings.h.node math/lp/nla_core.h.node
	@echo done > math/lp/nla_solver.h.node
math/lp/nla_solver$(OBJ_EXT): ../src/math/lp/nla_solver.cpp math/lp/nla_solver.h.node math/lp/monic.h.node math/lp/lp_utils.h.node math/lp/var_eqs.h.node math/lp/factorization.h.node math/lp/nla_solver.h.node math/lp/nla_core.h.node
	@echo src/math/lp/nla_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_solver$(OBJ_EXT) ../src/math/lp/nla_solver.cpp
math/lp/lp_settings_def.h.node: ../src/math/lp/lp_settings_def.h util/vector.h.node math/lp/lp_settings.h.node
	@echo done > math/lp/lp_settings_def.h.node
math/lp/lp_settings$(OBJ_EXT): ../src/math/lp/lp_settings.cpp util/vector.h.node math/lp/lp_settings_def.h.node
	@echo src/math/lp/lp_settings.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_settings$(OBJ_EXT) ../src/math/lp/lp_settings.cpp
math/lp/random_updater_def.h.node: ../src/math/lp/random_updater_def.h math/lp/random_updater.h.node math/lp/static_matrix.h.node math/lp/lar_solver.h.node util/vector.h.node
	@echo done > math/lp/random_updater_def.h.node
math/lp/random_updater$(OBJ_EXT): ../src/math/lp/random_updater.cpp math/lp/random_updater_def.h.node
	@echo src/math/lp/random_updater.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/random_updater$(OBJ_EXT) ../src/math/lp/random_updater.cpp
math/lp/core_solver_pretty_printer_def.h.node: ../src/math/lp/core_solver_pretty_printer_def.h math/lp/lp_utils.h.node math/lp/lp_core_solver_base.h.node math/lp/core_solver_pretty_printer.h.node math/lp/numeric_pair.h.node
	@echo done > math/lp/core_solver_pretty_printer_def.h.node
math/lp/core_solver_pretty_printer$(OBJ_EXT): ../src/math/lp/core_solver_pretty_printer.cpp math/lp/numeric_pair.h.node math/lp/core_solver_pretty_printer_def.h.node
	@echo src/math/lp/core_solver_pretty_printer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/core_solver_pretty_printer$(OBJ_EXT) ../src/math/lp/core_solver_pretty_printer.cpp
math/lp/gomory.h.node: ../src/math/lp/gomory.h math/lp/lar_term.h.node math/lp/lia_move.h.node math/lp/explanation.h.node math/lp/static_matrix.h.node
	@echo done > math/lp/gomory.h.node
math/lp/int_solver$(OBJ_EXT): ../src/math/lp/int_solver.cpp math/lp/int_solver.h.node math/lp/lar_solver.h.node math/lp/lp_utils.h.node math/lp/monic.h.node math/lp/gomory.h.node math/lp/int_branch.h.node math/lp/int_cube.h.node
	@echo src/math/lp/int_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/int_solver$(OBJ_EXT) ../src/math/lp/int_solver.cpp
math/lp/static_matrix_def.h.node: ../src/math/lp/static_matrix_def.h util/vector.h.node math/lp/static_matrix.h.node
	@echo done > math/lp/static_matrix_def.h.node
math/lp/lp_dual_simplex.h.node: ../src/math/lp/lp_dual_simplex.h util/vector.h.node math/lp/lp_utils.h.node math/lp/lp_solver.h.node math/lp/lp_dual_core_solver.h.node
	@echo done > math/lp/lp_dual_simplex.h.node
math/lp/static_matrix$(OBJ_EXT): ../src/math/lp/static_matrix.cpp util/vector.h.node math/lp/static_matrix_def.h.node math/lp/lp_core_solver_base.h.node math/lp/lp_dual_core_solver.h.node math/lp/lp_dual_simplex.h.node math/lp/lp_primal_core_solver.h.node math/lp/scaler.h.node math/lp/lar_solver.h.node
	@echo src/math/lp/static_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/static_matrix$(OBJ_EXT) ../src/math/lp/static_matrix.cpp
math/lp/lp_primal_simplex.h.node: ../src/math/lp/lp_primal_simplex.h util/vector.h.node math/lp/lp_utils.h.node math/lp/column_info.h.node math/lp/lp_primal_core_solver.h.node math/lp/lp_solver.h.node
	@echo done > math/lp/lp_primal_simplex.h.node
math/lp/lp_primal_simplex_def.h.node: ../src/math/lp/lp_primal_simplex_def.h util/vector.h.node math/lp/lp_primal_simplex.h.node
	@echo done > math/lp/lp_primal_simplex_def.h.node
math/lp/lp_primal_simplex$(OBJ_EXT): ../src/math/lp/lp_primal_simplex.cpp util/vector.h.node math/lp/lp_primal_simplex_def.h.node
	@echo src/math/lp/lp_primal_simplex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_primal_simplex$(OBJ_EXT) ../src/math/lp/lp_primal_simplex.cpp
math/lp/hnf_cutter$(OBJ_EXT): ../src/math/lp/hnf_cutter.cpp math/lp/int_solver.h.node math/lp/lar_solver.h.node math/lp/hnf_cutter.h.node
	@echo src/math/lp/hnf_cutter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/hnf_cutter$(OBJ_EXT) ../src/math/lp/hnf_cutter.cpp
math/lp/lp_solver_def.h.node: ../src/math/lp/lp_solver_def.h util/vector.h.node math/lp/lp_solver.h.node
	@echo done > math/lp/lp_solver_def.h.node
math/lp/lp_solver$(OBJ_EXT): ../src/math/lp/lp_solver.cpp math/lp/lp_solver_def.h.node
	@echo src/math/lp/lp_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_solver$(OBJ_EXT) ../src/math/lp/lp_solver.cpp
math/lp/nla_monotone_lemmas$(OBJ_EXT): ../src/math/lp/nla_monotone_lemmas.cpp math/lp/nla_basics_lemmas.h.node math/lp/nla_core.h.node
	@echo src/math/lp/nla_monotone_lemmas.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_monotone_lemmas$(OBJ_EXT) ../src/math/lp/nla_monotone_lemmas.cpp
math/lp/row_eta_matrix_def.h.node: ../src/math/lp/row_eta_matrix_def.h util/vector.h.node math/lp/row_eta_matrix.h.node
	@echo done > math/lp/row_eta_matrix_def.h.node
math/lp/row_eta_matrix$(OBJ_EXT): ../src/math/lp/row_eta_matrix.cpp util/vector.h.node math/lp/row_eta_matrix_def.h.node math/lp/lu.h.node
	@echo src/math/lp/row_eta_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/row_eta_matrix$(OBJ_EXT) ../src/math/lp/row_eta_matrix.cpp
math/lp/nex_creator$(OBJ_EXT): ../src/math/lp/nex_creator.cpp util/lbool.h.node math/lp/nex_creator.h.node
	@echo src/math/lp/nex_creator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nex_creator$(OBJ_EXT) ../src/math/lp/nex_creator.cpp
math/lp/mon_eq$(OBJ_EXT): ../src/math/lp/mon_eq.cpp math/lp/lar_solver.h.node math/lp/monic.h.node
	@echo src/math/lp/mon_eq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/mon_eq$(OBJ_EXT) ../src/math/lp/mon_eq.cpp
math/lp/gomory$(OBJ_EXT): ../src/math/lp/gomory.cpp math/lp/gomory.h.node math/lp/int_solver.h.node math/lp/lar_solver.h.node math/lp/lp_utils.h.node
	@echo src/math/lp/gomory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/gomory$(OBJ_EXT) ../src/math/lp/gomory.cpp
math/lp/factorization_factory_imp$(OBJ_EXT): ../src/math/lp/factorization_factory_imp.cpp math/lp/factorization_factory_imp.h.node math/lp/nla_core.h.node
	@echo src/math/lp/factorization_factory_imp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/factorization_factory_imp$(OBJ_EXT) ../src/math/lp/factorization_factory_imp.cpp
math/lp/polynomial.h.node: ../src/math/lp/polynomial.h
	@echo done > math/lp/polynomial.h.node
math/lp/nra_solver$(OBJ_EXT): ../src/math/lp/nra_solver.cpp math/lp/lar_solver.h.node math/lp/nra_solver.h.node nlsat/nlsat_solver.h.node math/lp/polynomial.h.node math/polynomial/algebraic_numbers.h.node util/map.h.node math/lp/u_set.h.node math/lp/nla_core.h.node
	@echo src/math/lp/nra_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nra_solver$(OBJ_EXT) ../src/math/lp/nra_solver.cpp
math/lp/nla_tangent_lemmas$(OBJ_EXT): ../src/math/lp/nla_tangent_lemmas.cpp math/lp/nla_tangent_lemmas.h.node math/lp/nla_core.h.node
	@echo src/math/lp/nla_tangent_lemmas.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_tangent_lemmas$(OBJ_EXT) ../src/math/lp/nla_tangent_lemmas.cpp
math/lp/permutation_matrix_def.h.node: ../src/math/lp/permutation_matrix_def.h util/vector.h.node math/lp/permutation_matrix.h.node
	@echo done > math/lp/permutation_matrix_def.h.node
math/lp/permutation_matrix$(OBJ_EXT): ../src/math/lp/permutation_matrix.cpp util/vector.h.node math/lp/permutation_matrix_def.h.node math/lp/numeric_pair.h.node
	@echo src/math/lp/permutation_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/permutation_matrix$(OBJ_EXT) ../src/math/lp/permutation_matrix.cpp
math/lp/lp_utils$(OBJ_EXT): ../src/math/lp/lp_utils.cpp math/lp/lp_utils.h.node
	@echo src/math/lp/lp_utils.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_utils$(OBJ_EXT) ../src/math/lp/lp_utils.cpp
math/lp/monomial_bounds$(OBJ_EXT): ../src/math/lp/monomial_bounds.cpp math/lp/monomial_bounds.h.node math/lp/nla_core.h.node math/lp/nla_intervals.h.node
	@echo src/math/lp/monomial_bounds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/monomial_bounds$(OBJ_EXT) ../src/math/lp/monomial_bounds.cpp
math/lp/int_gcd_test$(OBJ_EXT): ../src/math/lp/int_gcd_test.cpp math/lp/int_solver.h.node math/lp/lar_solver.h.node math/lp/int_gcd_test.h.node
	@echo src/math/lp/int_gcd_test.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/int_gcd_test$(OBJ_EXT) ../src/math/lp/int_gcd_test.cpp
math/lp/square_sparse_matrix_def.h.node: ../src/math/lp/square_sparse_matrix_def.h util/vector.h.node math/lp/square_sparse_matrix.h.node
	@echo done > math/lp/square_sparse_matrix_def.h.node
math/lp/square_sparse_matrix$(OBJ_EXT): ../src/math/lp/square_sparse_matrix.cpp util/vector.h.node math/lp/lp_settings.h.node math/lp/lu.h.node math/lp/square_sparse_matrix_def.h.node math/lp/dense_matrix.h.node
	@echo src/math/lp/square_sparse_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/square_sparse_matrix$(OBJ_EXT) ../src/math/lp/square_sparse_matrix.cpp
math/lp/lp_dual_simplex_def.h.node: ../src/math/lp/lp_dual_simplex_def.h math/lp/lp_dual_simplex.h.node
	@echo done > math/lp/lp_dual_simplex_def.h.node
math/lp/lp_dual_simplex$(OBJ_EXT): ../src/math/lp/lp_dual_simplex.cpp math/lp/lp_dual_simplex_def.h.node
	@echo src/math/lp/lp_dual_simplex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lp_dual_simplex$(OBJ_EXT) ../src/math/lp/lp_dual_simplex.cpp
math/lp/nla_basics_lemmas$(OBJ_EXT): ../src/math/lp/nla_basics_lemmas.cpp math/lp/nla_basics_lemmas.h.node math/lp/nla_core.h.node math/lp/factorization_factory_imp.h.node
	@echo src/math/lp/nla_basics_lemmas.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_basics_lemmas$(OBJ_EXT) ../src/math/lp/nla_basics_lemmas.cpp
math/lp/factorization$(OBJ_EXT): ../src/math/lp/factorization.cpp util/vector.h.node math/lp/factorization.h.node
	@echo src/math/lp/factorization.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/factorization$(OBJ_EXT) ../src/math/lp/factorization.cpp
math/lp/eta_matrix_def.h.node: ../src/math/lp/eta_matrix_def.h util/vector.h.node math/lp/eta_matrix.h.node
	@echo done > math/lp/eta_matrix_def.h.node
math/lp/eta_matrix$(OBJ_EXT): ../src/math/lp/eta_matrix.cpp util/vector.h.node math/lp/numeric_pair.h.node math/lp/eta_matrix_def.h.node
	@echo src/math/lp/eta_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/eta_matrix$(OBJ_EXT) ../src/math/lp/eta_matrix.cpp
math/lp/lar_core_solver_def.h.node: ../src/math/lp/lar_core_solver_def.h util/vector.h.node math/lp/lar_core_solver.h.node math/lp/lar_solution_signature.h.node
	@echo done > math/lp/lar_core_solver_def.h.node
math/lp/lar_core_solver$(OBJ_EXT): ../src/math/lp/lar_core_solver.cpp util/vector.h.node math/lp/lar_core_solver_def.h.node
	@echo src/math/lp/lar_core_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lar_core_solver$(OBJ_EXT) ../src/math/lp/lar_core_solver.cpp
math/lp/lu_def.h.node: ../src/math/lp/lu_def.h util/vector.h.node util/debug.h.node math/lp/lu.h.node
	@echo done > math/lp/lu_def.h.node
math/lp/lu$(OBJ_EXT): ../src/math/lp/lu.cpp util/vector.h.node util/debug.h.node math/lp/lu_def.h.node
	@echo src/math/lp/lu.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/lu$(OBJ_EXT) ../src/math/lp/lu.cpp
math/lp/nla_common$(OBJ_EXT): ../src/math/lp/nla_common.cpp math/lp/nla_common.h.node math/lp/nla_core.h.node
	@echo src/math/lp/nla_common.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_common$(OBJ_EXT) ../src/math/lp/nla_common.cpp
math/lp/nla_intervals$(OBJ_EXT): ../src/math/lp/nla_intervals.cpp math/lp/nla_core.h.node math/interval/interval_def.h.node math/lp/nla_intervals.h.node util/mpq.h.node
	@echo src/math/lp/nla_intervals.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_intervals$(OBJ_EXT) ../src/math/lp/nla_intervals.cpp
math/lp/binary_heap_priority_queue_def.h.node: ../src/math/lp/binary_heap_priority_queue_def.h util/vector.h.node math/lp/binary_heap_priority_queue.h.node
	@echo done > math/lp/binary_heap_priority_queue_def.h.node
math/lp/binary_heap_priority_queue$(OBJ_EXT): ../src/math/lp/binary_heap_priority_queue.cpp math/lp/numeric_pair.h.node math/lp/binary_heap_priority_queue_def.h.node
	@echo src/math/lp/binary_heap_priority_queue.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/binary_heap_priority_queue$(OBJ_EXT) ../src/math/lp/binary_heap_priority_queue.cpp
math/lp/dense_matrix_def.h.node: ../src/math/lp/dense_matrix_def.h math/lp/lp_settings.h.node util/vector.h.node math/lp/numeric_pair.h.node math/lp/dense_matrix.h.node
	@echo done > math/lp/dense_matrix_def.h.node
math/lp/dense_matrix$(OBJ_EXT): ../src/math/lp/dense_matrix.cpp math/lp/lp_settings.h.node math/lp/dense_matrix_def.h.node util/vector.h.node
	@echo src/math/lp/dense_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/dense_matrix$(OBJ_EXT) ../src/math/lp/dense_matrix.cpp
math/dd/pdd_interval.h.node: ../src/math/dd/pdd_interval.h math/dd/dd_pdd.h.node math/interval/dep_intervals.h.node
	@echo done > math/dd/pdd_interval.h.node
math/dd/pdd_eval.h.node: ../src/math/dd/pdd_eval.h math/dd/dd_pdd.h.node
	@echo done > math/dd/pdd_eval.h.node
math/lp/nla_core$(OBJ_EXT): ../src/math/lp/nla_core.cpp math/lp/nla_core.h.node math/lp/factorization_factory_imp.h.node math/lp/nex.h.node math/grobner/pdd_solver.h.node math/dd/pdd_interval.h.node math/dd/pdd_eval.h.node
	@echo src/math/lp/nla_core.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/nla_core$(OBJ_EXT) ../src/math/lp/nla_core.cpp
math/lp/binary_heap_upair_queue_def.h.node: ../src/math/lp/binary_heap_upair_queue_def.h math/lp/lp_utils.h.node math/lp/binary_heap_upair_queue.h.node
	@echo done > math/lp/binary_heap_upair_queue_def.h.node
math/lp/binary_heap_upair_queue$(OBJ_EXT): ../src/math/lp/binary_heap_upair_queue.cpp math/lp/binary_heap_upair_queue_def.h.node
	@echo src/math/lp/binary_heap_upair_queue.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/binary_heap_upair_queue$(OBJ_EXT) ../src/math/lp/binary_heap_upair_queue.cpp
math/lp/horner$(OBJ_EXT): ../src/math/lp/horner.cpp math/lp/horner.h.node math/lp/nla_core.h.node math/lp/lp_utils.h.node math/lp/cross_nested.h.node
	@echo src/math/lp/horner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_15) $(CXX_OUT_FLAG)math/lp/horner$(OBJ_EXT) ../src/math/lp/horner.cpp
math/lp/lp$(LIB_EXT): math/lp/int_cube$(OBJ_EXT) math/lp/lar_solver$(OBJ_EXT) math/lp/lp_dual_core_solver$(OBJ_EXT) math/lp/scaler$(OBJ_EXT) math/lp/indexed_vector$(OBJ_EXT) math/lp/matrix$(OBJ_EXT) math/lp/emonics$(OBJ_EXT) math/lp/int_branch$(OBJ_EXT) math/lp/nla_order_lemmas$(OBJ_EXT) math/lp/lp_primal_core_solver$(OBJ_EXT) math/lp/square_dense_submatrix$(OBJ_EXT) math/lp/lp_core_solver_base$(OBJ_EXT) math/lp/nla_solver$(OBJ_EXT) math/lp/lp_settings$(OBJ_EXT) math/lp/random_updater$(OBJ_EXT) math/lp/core_solver_pretty_printer$(OBJ_EXT) math/lp/int_solver$(OBJ_EXT) math/lp/static_matrix$(OBJ_EXT) math/lp/lp_primal_simplex$(OBJ_EXT) math/lp/hnf_cutter$(OBJ_EXT) math/lp/lp_solver$(OBJ_EXT) math/lp/nla_monotone_lemmas$(OBJ_EXT) math/lp/row_eta_matrix$(OBJ_EXT) math/lp/nex_creator$(OBJ_EXT) math/lp/mon_eq$(OBJ_EXT) math/lp/gomory$(OBJ_EXT) math/lp/factorization_factory_imp$(OBJ_EXT) math/lp/nra_solver$(OBJ_EXT) math/lp/nla_tangent_lemmas$(OBJ_EXT) math/lp/permutation_matrix$(OBJ_EXT) math/lp/lp_utils$(OBJ_EXT) math/lp/monomial_bounds$(OBJ_EXT) math/lp/int_gcd_test$(OBJ_EXT) math/lp/square_sparse_matrix$(OBJ_EXT) math/lp/lp_dual_simplex$(OBJ_EXT) math/lp/nla_basics_lemmas$(OBJ_EXT) math/lp/factorization$(OBJ_EXT) math/lp/eta_matrix$(OBJ_EXT) math/lp/lar_core_solver$(OBJ_EXT) math/lp/lu$(OBJ_EXT) math/lp/nla_common$(OBJ_EXT) math/lp/nla_intervals$(OBJ_EXT) math/lp/binary_heap_priority_queue$(OBJ_EXT) math/lp/dense_matrix$(OBJ_EXT) math/lp/nla_core$(OBJ_EXT) math/lp/binary_heap_upair_queue$(OBJ_EXT) math/lp/horner$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/lp/lp$(LIB_EXT) math/lp/int_cube$(OBJ_EXT) math/lp/lar_solver$(OBJ_EXT) math/lp/lp_dual_core_solver$(OBJ_EXT) math/lp/scaler$(OBJ_EXT) math/lp/indexed_vector$(OBJ_EXT) math/lp/matrix$(OBJ_EXT) math/lp/emonics$(OBJ_EXT) math/lp/int_branch$(OBJ_EXT) math/lp/nla_order_lemmas$(OBJ_EXT) math/lp/lp_primal_core_solver$(OBJ_EXT) math/lp/square_dense_submatrix$(OBJ_EXT) math/lp/lp_core_solver_base$(OBJ_EXT) math/lp/nla_solver$(OBJ_EXT) math/lp/lp_settings$(OBJ_EXT) math/lp/random_updater$(OBJ_EXT) math/lp/core_solver_pretty_printer$(OBJ_EXT) math/lp/int_solver$(OBJ_EXT) math/lp/static_matrix$(OBJ_EXT) math/lp/lp_primal_simplex$(OBJ_EXT) math/lp/hnf_cutter$(OBJ_EXT) math/lp/lp_solver$(OBJ_EXT) math/lp/nla_monotone_lemmas$(OBJ_EXT) math/lp/row_eta_matrix$(OBJ_EXT) math/lp/nex_creator$(OBJ_EXT) math/lp/mon_eq$(OBJ_EXT) math/lp/gomory$(OBJ_EXT) math/lp/factorization_factory_imp$(OBJ_EXT) math/lp/nra_solver$(OBJ_EXT) math/lp/nla_tangent_lemmas$(OBJ_EXT) math/lp/permutation_matrix$(OBJ_EXT) math/lp/lp_utils$(OBJ_EXT) math/lp/monomial_bounds$(OBJ_EXT) math/lp/int_gcd_test$(OBJ_EXT) math/lp/square_sparse_matrix$(OBJ_EXT) math/lp/lp_dual_simplex$(OBJ_EXT) math/lp/nla_basics_lemmas$(OBJ_EXT) math/lp/factorization$(OBJ_EXT) math/lp/eta_matrix$(OBJ_EXT) math/lp/lar_core_solver$(OBJ_EXT) math/lp/lu$(OBJ_EXT) math/lp/nla_common$(OBJ_EXT) math/lp/nla_intervals$(OBJ_EXT) math/lp/binary_heap_priority_queue$(OBJ_EXT) math/lp/dense_matrix$(OBJ_EXT) math/lp/nla_core$(OBJ_EXT) math/lp/binary_heap_upair_queue$(OBJ_EXT) math/lp/horner$(OBJ_EXT)
lp: math/lp/lp$(LIB_EXT)

includes_16 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src
ast/rewriter/rewriter_types.h.node: ../src/ast/rewriter/rewriter_types.h util/z3_exception.h.node util/common_msgs.h.node
	@echo done > ast/rewriter/rewriter_types.h.node
ast/rewriter/rewriter.h.node: ../src/ast/rewriter/rewriter.h ast/ast.h.node ast/rewriter/rewriter_types.h.node ast/act_cache.h.node
	@echo done > ast/rewriter/rewriter.h.node
ast/rewriter/maximize_ac_sharing.h.node: ../src/ast/rewriter/maximize_ac_sharing.h util/hashtable.h.node util/region.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter.h.node
	@echo done > ast/rewriter/maximize_ac_sharing.h.node
ast/rewriter/maximize_ac_sharing$(OBJ_EXT): ../src/ast/rewriter/maximize_ac_sharing.cpp ast/rewriter/maximize_ac_sharing.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/maximize_ac_sharing.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/maximize_ac_sharing$(OBJ_EXT) ../src/ast/rewriter/maximize_ac_sharing.cpp
ast/rewriter/bv_elim.h.node: ../src/ast/rewriter/bv_elim.h ast/ast.h.node ast/rewriter/rewriter.h.node
	@echo done > ast/rewriter/bv_elim.h.node
ast/rewriter/var_subst.h.node: ../src/ast/rewriter/var_subst.h ast/rewriter/rewriter.h.node ast/used_vars.h.node util/params.h.node
	@echo done > ast/rewriter/var_subst.h.node
ast/rewriter/rewriter_def.h.node: ../src/ast/rewriter/rewriter_def.h ast/rewriter/rewriter.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node
	@echo done > ast/rewriter/rewriter_def.h.node
ast/rewriter/bv_elim$(OBJ_EXT): ../src/ast/rewriter/bv_elim.cpp ast/rewriter/bv_elim.h.node ast/bv_decl_plugin.h.node ast/rewriter/var_subst.h.node ast/rewriter/rewriter_def.h.node
	@echo src/ast/rewriter/bv_elim.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/bv_elim$(OBJ_EXT) ../src/ast/rewriter/bv_elim.cpp
params/bv_rewriter_params.hpp.node: ../src/params/bv_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/bv_rewriter_params.hpp.node
ast/rewriter/poly_rewriter.h.node: ../src/ast/rewriter/poly_rewriter.h ast/ast.h.node util/obj_hashtable.h.node ast/rewriter/rewriter_types.h.node util/params.h.node
	@echo done > ast/rewriter/poly_rewriter.h.node
ast/rewriter/mk_extract_proc.h.node: ../src/ast/rewriter/mk_extract_proc.h ast/ast.h.node ast/bv_decl_plugin.h.node
	@echo done > ast/rewriter/mk_extract_proc.h.node
ast/rewriter/bv_rewriter.h.node: ../src/ast/rewriter/bv_rewriter.h ast/rewriter/poly_rewriter.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/rewriter/mk_extract_proc.h.node
	@echo done > ast/rewriter/bv_rewriter.h.node
util/container_util.h.node: ../src/util/container_util.h
	@echo done > util/container_util.h.node
params/poly_rewriter_params.hpp.node: ../src/params/poly_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/poly_rewriter_params.hpp.node
params/arith_rewriter_params.hpp.node: ../src/params/arith_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/arith_rewriter_params.hpp.node
ast/ast_lt.h.node: ../src/ast/ast_lt.h
	@echo done > ast/ast_lt.h.node
ast/rewriter/poly_rewriter_def.h.node: ../src/ast/rewriter/poly_rewriter_def.h util/container_util.h.node ast/rewriter/poly_rewriter.h.node params/poly_rewriter_params.hpp.node params/arith_rewriter_params.hpp.node ast/ast_lt.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node
	@echo done > ast/rewriter/poly_rewriter_def.h.node
ast/rewriter/bv_rewriter$(OBJ_EXT): ../src/ast/rewriter/bv_rewriter.cpp params/bv_rewriter_params.hpp.node ast/rewriter/bv_rewriter.h.node ast/rewriter/poly_rewriter_def.h.node ast/ast_smt2_pp.h.node ast/ast_lt.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/bv_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/bv_rewriter$(OBJ_EXT) ../src/ast/rewriter/bv_rewriter.cpp
ast/rewriter/push_app_ite.h.node: ../src/ast/rewriter/push_app_ite.h ast/ast.h.node ast/rewriter/rewriter.h.node
	@echo done > ast/rewriter/push_app_ite.h.node
ast/rewriter/push_app_ite$(OBJ_EXT): ../src/ast/rewriter/push_app_ite.cpp ast/rewriter/push_app_ite.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/push_app_ite.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/push_app_ite$(OBJ_EXT) ../src/ast/rewriter/push_app_ite.cpp
ast/rewriter/expr_safe_replace.h.node: ../src/ast/rewriter/expr_safe_replace.h ast/ast.h.node
	@echo done > ast/rewriter/expr_safe_replace.h.node
ast/rewriter/factor_equivs.h.node: ../src/ast/rewriter/factor_equivs.h util/union_find.h.node ast/ast_util.h.node
	@echo done > ast/rewriter/factor_equivs.h.node
ast/rewriter/factor_equivs$(OBJ_EXT): ../src/ast/rewriter/factor_equivs.cpp ast/arith_decl_plugin.h.node ast/for_each_expr.h.node ast/ast_pp.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/factor_equivs.h.node
	@echo src/ast/rewriter/factor_equivs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/factor_equivs$(OBJ_EXT) ../src/ast/rewriter/factor_equivs.cpp
ast/rewriter/bv_bounds.h.node: ../src/ast/rewriter/bv_bounds.h ast/ast.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter_types.h.node
	@echo done > ast/rewriter/bv_bounds.h.node
ast/rewriter/bv_bounds$(OBJ_EXT): ../src/ast/rewriter/bv_bounds.cpp ast/rewriter/bv_bounds.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/rewriter/bv_bounds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/bv_bounds$(OBJ_EXT) ../src/ast/rewriter/bv_bounds.cpp
ast/rewriter/enum2bv_rewriter.h.node: ../src/ast/rewriter/enum2bv_rewriter.h ast/datatype_decl_plugin.h.node ast/rewriter/rewriter_types.h.node ast/expr_functors.h.node
	@echo done > ast/rewriter/enum2bv_rewriter.h.node
ast/rewriter/enum2bv_rewriter$(OBJ_EXT): ../src/ast/rewriter/enum2bv_rewriter.cpp ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/enum2bv_rewriter.h.node ast/ast_util.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/enum2bv_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/enum2bv_rewriter$(OBJ_EXT) ../src/ast/rewriter/enum2bv_rewriter.cpp
ast/rewriter/bit2int.h.node: ../src/ast/rewriter/bit2int.h ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/expr_map.h.node ast/rewriter/bv_rewriter.h.node
	@echo done > ast/rewriter/bit2int.h.node
ast/rewriter/bit2int$(OBJ_EXT): ../src/ast/rewriter/bit2int.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/for_each_ast.h.node ast/rewriter/bit2int.h.node
	@echo src/ast/rewriter/bit2int.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/bit2int$(OBJ_EXT) ../src/ast/rewriter/bit2int.cpp
ast/rewriter/array_rewriter.h.node: ../src/ast/rewriter/array_rewriter.h ast/array_decl_plugin.h.node ast/rewriter/rewriter_types.h.node util/lbool.h.node util/params.h.node
	@echo done > ast/rewriter/array_rewriter.h.node
params/array_rewriter_params.hpp.node: ../src/params/array_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/array_rewriter_params.hpp.node
ast/rewriter/array_rewriter$(OBJ_EXT): ../src/ast/rewriter/array_rewriter.cpp ast/rewriter/array_rewriter.h.node ast/ast_lt.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/rewriter/var_subst.h.node params/array_rewriter_params.hpp.node
	@echo src/ast/rewriter/array_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/array_rewriter$(OBJ_EXT) ../src/ast/rewriter/array_rewriter.cpp
ast/rewriter/var_subst$(OBJ_EXT): ../src/ast/rewriter/var_subst.cpp ast/rewriter/var_subst.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node ast/well_sorted.h.node ast/for_each_expr.h.node
	@echo src/ast/rewriter/var_subst.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/var_subst$(OBJ_EXT) ../src/ast/rewriter/var_subst.cpp
ast/rewriter/label_rewriter.h.node: ../src/ast/rewriter/label_rewriter.h ast/ast.h.node ast/rewriter/rewriter_types.h.node
	@echo done > ast/rewriter/label_rewriter.h.node
ast/rewriter/label_rewriter$(OBJ_EXT): ../src/ast/rewriter/label_rewriter.cpp ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/label_rewriter.h.node
	@echo src/ast/rewriter/label_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/label_rewriter$(OBJ_EXT) ../src/ast/rewriter/label_rewriter.cpp
ast/rewriter/arith_rewriter.h.node: ../src/ast/rewriter/arith_rewriter.h ast/rewriter/poly_rewriter.h.node ast/arith_decl_plugin.h.node ast/seq_decl_plugin.h.node
	@echo done > ast/rewriter/arith_rewriter.h.node
ast/rewriter/arith_rewriter$(OBJ_EXT): ../src/ast/rewriter/arith_rewriter.cpp params/arith_rewriter_params.hpp.node ast/rewriter/arith_rewriter.h.node ast/rewriter/poly_rewriter_def.h.node math/polynomial/algebraic_numbers.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/arith_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/arith_rewriter$(OBJ_EXT) ../src/ast/rewriter/arith_rewriter.cpp
params/rewriter_params.hpp.node: ../src/params/rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/rewriter_params.hpp.node
ast/rewriter/th_rewriter.h.node: ../src/ast/rewriter/th_rewriter.h ast/ast.h.node ast/rewriter/rewriter_types.h.node util/params.h.node
	@echo done > ast/rewriter/th_rewriter.h.node
ast/rewriter/bool_rewriter.h.node: ../src/ast/rewriter/bool_rewriter.h ast/ast.h.node ast/rewriter/rewriter.h.node util/params.h.node
	@echo done > ast/rewriter/bool_rewriter.h.node
ast/rewriter/datatype_rewriter.h.node: ../src/ast/rewriter/datatype_rewriter.h ast/datatype_decl_plugin.h.node ast/rewriter/rewriter_types.h.node
	@echo done > ast/rewriter/datatype_rewriter.h.node
ast/rewriter/fpa_rewriter.h.node: ../src/ast/rewriter/fpa_rewriter.h ast/ast.h.node ast/rewriter/rewriter.h.node ast/fpa_decl_plugin.h.node ast/expr_map.h.node util/params.h.node util/mpf.h.node
	@echo done > ast/rewriter/fpa_rewriter.h.node
ast/rewriter/dl_rewriter.h.node: ../src/ast/rewriter/dl_rewriter.h ast/dl_decl_plugin.h.node ast/rewriter/rewriter_types.h.node
	@echo done > ast/rewriter/dl_rewriter.h.node
ast/rewriter/pb_rewriter.h.node: ../src/ast/rewriter/pb_rewriter.h ast/pb_decl_plugin.h.node ast/rewriter/rewriter_types.h.node util/params.h.node util/lbool.h.node
	@echo done > ast/rewriter/pb_rewriter.h.node
ast/rewriter/recfun_rewriter.h.node: ../src/ast/rewriter/recfun_rewriter.h ast/recfun_decl_plugin.h.node ast/rewriter/rewriter.h.node
	@echo done > ast/rewriter/recfun_rewriter.h.node
math/automata/boolean_algebra.h.node: ../src/math/automata/boolean_algebra.h util/util.h.node
	@echo done > math/automata/boolean_algebra.h.node
math/automata/symbolic_automata.h.node: ../src/math/automata/symbolic_automata.h math/automata/automaton.h.node math/automata/boolean_algebra.h.node
	@echo done > math/automata/symbolic_automata.h.node
ast/rewriter/seq_rewriter.h.node: ../src/ast/rewriter/seq_rewriter.h ast/seq_decl_plugin.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/rewriter/rewriter_types.h.node ast/rewriter/bool_rewriter.h.node util/params.h.node util/lbool.h.node util/sign.h.node math/automata/automaton.h.node math/automata/symbolic_automata.h.node
	@echo done > ast/rewriter/seq_rewriter.h.node
ast/rewriter/th_rewriter$(OBJ_EXT): ../src/ast/rewriter/th_rewriter.cpp params/rewriter_params.hpp.node ast/rewriter/th_rewriter.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/bv_rewriter.h.node ast/rewriter/datatype_rewriter.h.node ast/rewriter/array_rewriter.h.node ast/rewriter/fpa_rewriter.h.node ast/rewriter/dl_rewriter.h.node ast/rewriter/pb_rewriter.h.node ast/rewriter/recfun_rewriter.h.node ast/rewriter/seq_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node ast/rewriter/expr_safe_replace.h.node ast/expr_substitution.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/well_sorted.h.node
	@echo src/ast/rewriter/th_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/th_rewriter$(OBJ_EXT) ../src/ast/rewriter/th_rewriter.cpp
ast/rewriter/inj_axiom.h.node: ../src/ast/rewriter/inj_axiom.h ast/ast.h.node
	@echo done > ast/rewriter/inj_axiom.h.node
ast/has_free_vars.h.node: ../src/ast/has_free_vars.h
	@echo done > ast/has_free_vars.h.node
ast/rewriter/inj_axiom$(OBJ_EXT): ../src/ast/rewriter/inj_axiom.cpp ast/rewriter/inj_axiom.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/has_free_vars.h.node ast/well_sorted.h.node
	@echo src/ast/rewriter/inj_axiom.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/inj_axiom$(OBJ_EXT) ../src/ast/rewriter/inj_axiom.cpp
params/fpa_rewriter_params.hpp.node: ../src/params/fpa_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/fpa_rewriter_params.hpp.node
ast/rewriter/fpa_rewriter$(OBJ_EXT): ../src/ast/rewriter/fpa_rewriter.cpp ast/rewriter/fpa_rewriter.h.node params/fpa_rewriter_params.hpp.node ast/ast_smt2_pp.h.node
	@echo src/ast/rewriter/fpa_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/fpa_rewriter$(OBJ_EXT) ../src/ast/rewriter/fpa_rewriter.cpp
ast/rewriter/elim_bounds.h.node: ../src/ast/rewriter/elim_bounds.h ast/ast.h.node ast/arith_decl_plugin.h.node ast/rewriter/rewriter.h.node
	@echo done > ast/rewriter/elim_bounds.h.node
ast/rewriter/elim_bounds$(OBJ_EXT): ../src/ast/rewriter/elim_bounds.cpp ast/used_vars.h.node util/obj_hashtable.h.node ast/rewriter/var_subst.h.node ast/rewriter/elim_bounds.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/elim_bounds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/elim_bounds$(OBJ_EXT) ../src/ast/rewriter/elim_bounds.cpp
ast/rewriter/hoist_rewriter.h.node: ../src/ast/rewriter/hoist_rewriter.h ast/ast.h.node ast/rewriter/rewriter.h.node ast/rewriter/expr_safe_replace.h.node util/params.h.node util/union_find.h.node util/obj_hashtable.h.node
	@echo done > ast/rewriter/hoist_rewriter.h.node
ast/rewriter/hoist_rewriter$(OBJ_EXT): ../src/ast/rewriter/hoist_rewriter.cpp ast/rewriter/hoist_rewriter.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/ast/rewriter/hoist_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/hoist_rewriter$(OBJ_EXT) ../src/ast/rewriter/hoist_rewriter.cpp
ast/rewriter/expr_replacer.h.node: ../src/ast/rewriter/expr_replacer.h ast/ast.h.node ast/expr_substitution.h.node util/params.h.node
	@echo done > ast/rewriter/expr_replacer.h.node
ast/rewriter/expr_replacer$(OBJ_EXT): ../src/ast/rewriter/expr_replacer.cpp ast/rewriter/expr_replacer.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/th_rewriter.h.node
	@echo src/ast/rewriter/expr_replacer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/expr_replacer$(OBJ_EXT) ../src/ast/rewriter/expr_replacer.cpp
ast/rewriter/mk_simplified_app.h.node: ../src/ast/rewriter/mk_simplified_app.h ast/ast.h.node util/params.h.node ast/rewriter/rewriter_types.h.node
	@echo done > ast/rewriter/mk_simplified_app.h.node
ast/rewriter/mk_simplified_app$(OBJ_EXT): ../src/ast/rewriter/mk_simplified_app.cpp ast/rewriter/mk_simplified_app.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/bv_rewriter.h.node ast/rewriter/datatype_rewriter.h.node ast/rewriter/array_rewriter.h.node ast/rewriter/fpa_rewriter.h.node
	@echo src/ast/rewriter/mk_simplified_app.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/mk_simplified_app$(OBJ_EXT) ../src/ast/rewriter/mk_simplified_app.cpp
ast/rewriter/mk_extract_proc$(OBJ_EXT): ../src/ast/rewriter/mk_extract_proc.cpp ast/rewriter/mk_extract_proc.h.node
	@echo src/ast/rewriter/mk_extract_proc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/mk_extract_proc$(OBJ_EXT) ../src/ast/rewriter/mk_extract_proc.cpp
ast/rewriter/expr_safe_replace$(OBJ_EXT): ../src/ast/rewriter/expr_safe_replace.cpp ast/rewriter/expr_safe_replace.h.node ast/rewriter/var_subst.h.node ast/ast_pp.h.node
	@echo src/ast/rewriter/expr_safe_replace.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/expr_safe_replace$(OBJ_EXT) ../src/ast/rewriter/expr_safe_replace.cpp
ast/rewriter/func_decl_replace.h.node: ../src/ast/rewriter/func_decl_replace.h ast/ast.h.node
	@echo done > ast/rewriter/func_decl_replace.h.node
ast/rewriter/func_decl_replace$(OBJ_EXT): ../src/ast/rewriter/func_decl_replace.cpp ast/rewriter/func_decl_replace.h.node
	@echo src/ast/rewriter/func_decl_replace.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/func_decl_replace$(OBJ_EXT) ../src/ast/rewriter/func_decl_replace.cpp
ast/rewriter/rewriter$(OBJ_EXT): ../src/ast/rewriter/rewriter.cpp ast/rewriter/rewriter_def.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/rewriter/rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/rewriter$(OBJ_EXT) ../src/ast/rewriter/rewriter.cpp
ast/rewriter/value_sweep.h.node: ../src/ast/rewriter/value_sweep.h ast/value_generator.h.node ast/recfun_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/rewriter/th_rewriter.h.node
	@echo done > ast/rewriter/value_sweep.h.node
ast/rewriter/value_sweep$(OBJ_EXT): ../src/ast/rewriter/value_sweep.cpp ast/for_each_expr.h.node ast/ast_pp.h.node ast/rewriter/value_sweep.h.node
	@echo src/ast/rewriter/value_sweep.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/value_sweep$(OBJ_EXT) ../src/ast/rewriter/value_sweep.cpp
ast/rewriter/der.h.node: ../src/ast/rewriter/der.h ast/ast.h.node ast/rewriter/var_subst.h.node
	@echo done > ast/rewriter/der.h.node
ast/rewriter/der$(OBJ_EXT): ../src/ast/rewriter/der.cpp ast/rewriter/der.h.node ast/occurs.h.node ast/for_each_expr.h.node ast/rewriter/rewriter_def.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/rewriter/der.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/der$(OBJ_EXT) ../src/ast/rewriter/der.cpp
ast/rewriter/pb2bv_rewriter.h.node: ../src/ast/rewriter/pb2bv_rewriter.h ast/pb_decl_plugin.h.node ast/rewriter/rewriter_types.h.node ast/expr_functors.h.node
	@echo done > ast/rewriter/pb2bv_rewriter.h.node
util/sorting_network.h.node: ../src/util/sorting_network.h util/vector.h.node
	@echo done > util/sorting_network.h.node
ast/rewriter/pb2bv_rewriter$(OBJ_EXT): ../src/ast/rewriter/pb2bv_rewriter.cpp util/statistics.h.node util/lbool.h.node util/uint_set.h.node util/gparams.h.node util/debug.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/pb2bv_rewriter.h.node ast/ast_util.h.node ast/ast_pp.h.node util/sorting_network.h.node
	@echo src/ast/rewriter/pb2bv_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/pb2bv_rewriter$(OBJ_EXT) ../src/ast/rewriter/pb2bv_rewriter.cpp
ast/rewriter/recfun_rewriter$(OBJ_EXT): ../src/ast/rewriter/recfun_rewriter.cpp ast/rewriter/recfun_rewriter.h.node ast/rewriter/var_subst.h.node
	@echo src/ast/rewriter/recfun_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/recfun_rewriter$(OBJ_EXT) ../src/ast/rewriter/recfun_rewriter.cpp
ast/rewriter/dl_rewriter$(OBJ_EXT): ../src/ast/rewriter/dl_rewriter.cpp ast/rewriter/dl_rewriter.h.node
	@echo src/ast/rewriter/dl_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/dl_rewriter$(OBJ_EXT) ../src/ast/rewriter/dl_rewriter.cpp
params/bool_rewriter_params.hpp.node: ../src/params/bool_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/bool_rewriter_params.hpp.node
ast/rewriter/bool_rewriter$(OBJ_EXT): ../src/ast/rewriter/bool_rewriter.cpp ast/rewriter/bool_rewriter.h.node params/bool_rewriter_params.hpp.node ast/rewriter/rewriter_def.h.node ast/ast_lt.h.node
	@echo src/ast/rewriter/bool_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/bool_rewriter$(OBJ_EXT) ../src/ast/rewriter/bool_rewriter.cpp
ast/rewriter/factor_rewriter.h.node: ../src/ast/rewriter/factor_rewriter.h ast/ast.h.node ast/rewriter/rewriter.h.node ast/arith_decl_plugin.h.node
	@echo done > ast/rewriter/factor_rewriter.h.node
ast/rewriter/factor_rewriter$(OBJ_EXT): ../src/ast/rewriter/factor_rewriter.cpp ast/rewriter/factor_rewriter.h.node ast/ast_pp.h.node ast/rewriter/rewriter_def.h.node
	@echo src/ast/rewriter/factor_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/factor_rewriter$(OBJ_EXT) ../src/ast/rewriter/factor_rewriter.cpp
ast/rewriter/datatype_rewriter$(OBJ_EXT): ../src/ast/rewriter/datatype_rewriter.cpp ast/rewriter/datatype_rewriter.h.node
	@echo src/ast/rewriter/datatype_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/datatype_rewriter$(OBJ_EXT) ../src/ast/rewriter/datatype_rewriter.cpp
params/seq_rewriter_params.hpp.node: ../src/params/seq_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/seq_rewriter_params.hpp.node
math/automata/symbolic_automata_def.h.node: ../src/math/automata/symbolic_automata_def.h math/automata/symbolic_automata.h.node util/hashtable.h.node util/vector.h.node
	@echo done > math/automata/symbolic_automata_def.h.node
ast/rewriter/seq_rewriter$(OBJ_EXT): ../src/ast/rewriter/seq_rewriter.cpp util/uint_set.h.node ast/rewriter/seq_rewriter.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_util.h.node ast/well_sorted.h.node ast/rewriter/var_subst.h.node ast/rewriter/expr_safe_replace.h.node params/seq_rewriter_params.hpp.node math/automata/automaton.h.node math/automata/symbolic_automata_def.h.node
	@echo src/ast/rewriter/seq_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/seq_rewriter$(OBJ_EXT) ../src/ast/rewriter/seq_rewriter.cpp
ast/rewriter/ast_counter.h.node: ../src/ast/rewriter/ast_counter.h ast/ast.h.node util/map.h.node util/uint_set.h.node ast/rewriter/var_subst.h.node
	@echo done > ast/rewriter/ast_counter.h.node
ast/rewriter/ast_counter$(OBJ_EXT): ../src/ast/rewriter/ast_counter.cpp ast/rewriter/ast_counter.h.node
	@echo src/ast/rewriter/ast_counter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/ast_counter$(OBJ_EXT) ../src/ast/rewriter/ast_counter.cpp
ast/rewriter/pb_rewriter_def.h.node: ../src/ast/rewriter/pb_rewriter_def.h ast/rewriter/pb_rewriter.h.node
	@echo done > ast/rewriter/pb_rewriter_def.h.node
ast/rewriter/pb_rewriter$(OBJ_EXT): ../src/ast/rewriter/pb_rewriter.cpp ast/rewriter/pb_rewriter.h.node ast/rewriter/pb_rewriter_def.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/ast_smt_pp.h.node
	@echo src/ast/rewriter/pb_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/pb_rewriter$(OBJ_EXT) ../src/ast/rewriter/pb_rewriter.cpp
ast/rewriter/quant_hoist.h.node: ../src/ast/rewriter/quant_hoist.h ast/ast.h.node
	@echo done > ast/rewriter/quant_hoist.h.node
ast/rewriter/quant_hoist$(OBJ_EXT): ../src/ast/rewriter/quant_hoist.cpp ast/rewriter/quant_hoist.h.node ast/expr_functors.h.node ast/ast_smt_pp.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/var_subst.h.node ast/ast_pp.h.node ast/rewriter/ast_counter.h.node ast/rewriter/expr_safe_replace.h.node
	@echo src/ast/rewriter/quant_hoist.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/quant_hoist$(OBJ_EXT) ../src/ast/rewriter/quant_hoist.cpp
ast/rewriter/distribute_forall.h.node: ../src/ast/rewriter/distribute_forall.h ast/ast.h.node ast/act_cache.h.node
	@echo done > ast/rewriter/distribute_forall.h.node
ast/rewriter/distribute_forall$(OBJ_EXT): ../src/ast/rewriter/distribute_forall.cpp ast/rewriter/var_subst.h.node ast/ast_ll_pp.h.node ast/ast_util.h.node ast/rewriter/distribute_forall.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/ast/rewriter/distribute_forall.cpp
	@$(CXX) $(CXXFLAGS) $(includes_16) $(CXX_OUT_FLAG)ast/rewriter/distribute_forall$(OBJ_EXT) ../src/ast/rewriter/distribute_forall.cpp
ast/rewriter/rewriter$(LIB_EXT): ast/rewriter/maximize_ac_sharing$(OBJ_EXT) ast/rewriter/bv_elim$(OBJ_EXT) ast/rewriter/bv_rewriter$(OBJ_EXT) ast/rewriter/push_app_ite$(OBJ_EXT) ast/rewriter/factor_equivs$(OBJ_EXT) ast/rewriter/bv_bounds$(OBJ_EXT) ast/rewriter/enum2bv_rewriter$(OBJ_EXT) ast/rewriter/bit2int$(OBJ_EXT) ast/rewriter/array_rewriter$(OBJ_EXT) ast/rewriter/var_subst$(OBJ_EXT) ast/rewriter/label_rewriter$(OBJ_EXT) ast/rewriter/arith_rewriter$(OBJ_EXT) ast/rewriter/th_rewriter$(OBJ_EXT) ast/rewriter/inj_axiom$(OBJ_EXT) ast/rewriter/fpa_rewriter$(OBJ_EXT) ast/rewriter/elim_bounds$(OBJ_EXT) ast/rewriter/hoist_rewriter$(OBJ_EXT) ast/rewriter/expr_replacer$(OBJ_EXT) ast/rewriter/mk_simplified_app$(OBJ_EXT) ast/rewriter/mk_extract_proc$(OBJ_EXT) ast/rewriter/expr_safe_replace$(OBJ_EXT) ast/rewriter/func_decl_replace$(OBJ_EXT) ast/rewriter/rewriter$(OBJ_EXT) ast/rewriter/value_sweep$(OBJ_EXT) ast/rewriter/der$(OBJ_EXT) ast/rewriter/pb2bv_rewriter$(OBJ_EXT) ast/rewriter/recfun_rewriter$(OBJ_EXT) ast/rewriter/dl_rewriter$(OBJ_EXT) ast/rewriter/bool_rewriter$(OBJ_EXT) ast/rewriter/factor_rewriter$(OBJ_EXT) ast/rewriter/datatype_rewriter$(OBJ_EXT) ast/rewriter/seq_rewriter$(OBJ_EXT) ast/rewriter/ast_counter$(OBJ_EXT) ast/rewriter/pb_rewriter$(OBJ_EXT) ast/rewriter/quant_hoist$(OBJ_EXT) ast/rewriter/distribute_forall$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/rewriter/rewriter$(LIB_EXT) ast/rewriter/maximize_ac_sharing$(OBJ_EXT) ast/rewriter/bv_elim$(OBJ_EXT) ast/rewriter/bv_rewriter$(OBJ_EXT) ast/rewriter/push_app_ite$(OBJ_EXT) ast/rewriter/factor_equivs$(OBJ_EXT) ast/rewriter/bv_bounds$(OBJ_EXT) ast/rewriter/enum2bv_rewriter$(OBJ_EXT) ast/rewriter/bit2int$(OBJ_EXT) ast/rewriter/array_rewriter$(OBJ_EXT) ast/rewriter/var_subst$(OBJ_EXT) ast/rewriter/label_rewriter$(OBJ_EXT) ast/rewriter/arith_rewriter$(OBJ_EXT) ast/rewriter/th_rewriter$(OBJ_EXT) ast/rewriter/inj_axiom$(OBJ_EXT) ast/rewriter/fpa_rewriter$(OBJ_EXT) ast/rewriter/elim_bounds$(OBJ_EXT) ast/rewriter/hoist_rewriter$(OBJ_EXT) ast/rewriter/expr_replacer$(OBJ_EXT) ast/rewriter/mk_simplified_app$(OBJ_EXT) ast/rewriter/mk_extract_proc$(OBJ_EXT) ast/rewriter/expr_safe_replace$(OBJ_EXT) ast/rewriter/func_decl_replace$(OBJ_EXT) ast/rewriter/rewriter$(OBJ_EXT) ast/rewriter/value_sweep$(OBJ_EXT) ast/rewriter/der$(OBJ_EXT) ast/rewriter/pb2bv_rewriter$(OBJ_EXT) ast/rewriter/recfun_rewriter$(OBJ_EXT) ast/rewriter/dl_rewriter$(OBJ_EXT) ast/rewriter/bool_rewriter$(OBJ_EXT) ast/rewriter/factor_rewriter$(OBJ_EXT) ast/rewriter/datatype_rewriter$(OBJ_EXT) ast/rewriter/seq_rewriter$(OBJ_EXT) ast/rewriter/ast_counter$(OBJ_EXT) ast/rewriter/pb_rewriter$(OBJ_EXT) ast/rewriter/quant_hoist$(OBJ_EXT) ast/rewriter/distribute_forall$(OBJ_EXT)
rewriter: ast/rewriter/rewriter$(LIB_EXT)

includes_17 = -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src
ast/macros/cond_macro.h.node: ../src/ast/macros/cond_macro.h ast/ast_ll_pp.h.node
	@echo done > ast/macros/cond_macro.h.node
ast/macros/macro_util.h.node: ../src/ast/macros/macro_util.h ast/ast.h.node util/obj_hashtable.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/bv_rewriter.h.node
	@echo done > ast/macros/macro_util.h.node
ast/macros/quantifier_macro_info.h.node: ../src/ast/macros/quantifier_macro_info.h util/scoped_ptr_vector.h.node ast/ast_ll_pp.h.node ast/macros/cond_macro.h.node ast/macros/macro_util.h.node ast/func_decl_dependencies.h.node
	@echo done > ast/macros/quantifier_macro_info.h.node
ast/macros/quantifier_macro_info$(OBJ_EXT): ../src/ast/macros/quantifier_macro_info.cpp ast/ast_pp.h.node ast/macros/quantifier_macro_info.h.node
	@echo src/ast/macros/quantifier_macro_info.cpp
	@$(CXX) $(CXXFLAGS) $(includes_17) $(CXX_OUT_FLAG)ast/macros/quantifier_macro_info$(OBJ_EXT) ../src/ast/macros/quantifier_macro_info.cpp
ast/justified_expr.h.node: ../src/ast/justified_expr.h ast/ast.h.node
	@echo done > ast/justified_expr.h.node
ast/macros/macro_manager.h.node: ../src/ast/macros/macro_manager.h util/obj_hashtable.h.node ast/ast_util.h.node ast/justified_expr.h.node ast/recurse_expr.h.node ast/func_decl_dependencies.h.node ast/macros/macro_util.h.node
	@echo done > ast/macros/macro_manager.h.node
ast/macros/macro_finder.h.node: ../src/ast/macros/macro_finder.h ast/macros/macro_manager.h.node
	@echo done > ast/macros/macro_finder.h.node
ast/macros/macro_finder$(OBJ_EXT): ../src/ast/macros/macro_finder.cpp ast/macros/macro_finder.h.node ast/occurs.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/ast/macros/macro_finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_17) $(CXX_OUT_FLAG)ast/macros/macro_finder$(OBJ_EXT) ../src/ast/macros/macro_finder.cpp
ast/macros/macro_util$(OBJ_EXT): ../src/ast/macros/macro_util.cpp ast/macros/macro_util.h.node ast/occurs.h.node ast/ast_util.h.node ast/rewriter/var_subst.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/for_each_expr.h.node ast/well_sorted.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/ast/macros/macro_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_17) $(CXX_OUT_FLAG)ast/macros/macro_util$(OBJ_EXT) ../src/ast/macros/macro_util.cpp
ast/macros/quasi_macros.h.node: ../src/ast/macros/quasi_macros.h ast/justified_expr.h.node ast/macros/macro_manager.h.node ast/rewriter/th_rewriter.h.node
	@echo done > ast/macros/quasi_macros.h.node
ast/macros/quasi_macros$(OBJ_EXT): ../src/ast/macros/quasi_macros.cpp ast/macros/quasi_macros.h.node ast/for_each_expr.h.node ast/ast_pp.h.node util/uint_set.h.node ast/rewriter/var_subst.h.node
	@echo src/ast/macros/quasi_macros.cpp
	@$(CXX) $(CXXFLAGS) $(includes_17) $(CXX_OUT_FLAG)ast/macros/quasi_macros$(OBJ_EXT) ../src/ast/macros/quasi_macros.cpp
ast/macros/macro_manager$(OBJ_EXT): ../src/ast/macros/macro_manager.cpp ast/macros/macro_manager.h.node ast/for_each_expr.h.node ast/rewriter/var_subst.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node ast/ast_translation.h.node ast/recurse_expr_def.h.node
	@echo src/ast/macros/macro_manager.cpp
	@$(CXX) $(CXXFLAGS) $(includes_17) $(CXX_OUT_FLAG)ast/macros/macro_manager$(OBJ_EXT) ../src/ast/macros/macro_manager.cpp
ast/macros/macros$(LIB_EXT): ast/macros/quantifier_macro_info$(OBJ_EXT) ast/macros/macro_finder$(OBJ_EXT) ast/macros/macro_util$(OBJ_EXT) ast/macros/quasi_macros$(OBJ_EXT) ast/macros/macro_manager$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/macros/macros$(LIB_EXT) ast/macros/quantifier_macro_info$(OBJ_EXT) ast/macros/macro_finder$(OBJ_EXT) ast/macros/macro_util$(OBJ_EXT) ast/macros/quasi_macros$(OBJ_EXT) ast/macros/macro_manager$(OBJ_EXT)
macros: ast/macros/macros$(LIB_EXT)

includes_18 = -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src
ast/normal_forms/defined_names.h.node: ../src/ast/normal_forms/defined_names.h ast/ast.h.node
	@echo done > ast/normal_forms/defined_names.h.node
ast/normal_forms/defined_names$(OBJ_EXT): ../src/ast/normal_forms/defined_names.cpp util/obj_hashtable.h.node ast/normal_forms/defined_names.h.node ast/used_vars.h.node ast/rewriter/var_subst.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/array_decl_plugin.h.node
	@echo src/ast/normal_forms/defined_names.cpp
	@$(CXX) $(CXXFLAGS) $(includes_18) $(CXX_OUT_FLAG)ast/normal_forms/defined_names$(OBJ_EXT) ../src/ast/normal_forms/defined_names.cpp
ast/normal_forms/name_exprs.h.node: ../src/ast/normal_forms/name_exprs.h ast/ast.h.node ast/normal_forms/defined_names.h.node
	@echo done > ast/normal_forms/name_exprs.h.node
ast/normal_forms/name_exprs$(OBJ_EXT): ../src/ast/normal_forms/name_exprs.cpp ast/normal_forms/name_exprs.h.node ast/rewriter/rewriter_def.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/normal_forms/name_exprs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_18) $(CXX_OUT_FLAG)ast/normal_forms/name_exprs$(OBJ_EXT) ../src/ast/normal_forms/name_exprs.cpp
ast/normal_forms/nnf.h.node: ../src/ast/normal_forms/nnf.h ast/ast.h.node util/params.h.node ast/normal_forms/defined_names.h.node
	@echo done > ast/normal_forms/nnf.h.node
ast/normal_forms/nnf_params.hpp.node: ../src/ast/normal_forms/nnf_params.hpp util/params.h.node util/gparams.h.node
	@echo done > ast/normal_forms/nnf_params.hpp.node
ast/normal_forms/nnf$(OBJ_EXT): ../src/ast/normal_forms/nnf.cpp util/warning.h.node ast/normal_forms/nnf.h.node ast/normal_forms/nnf_params.hpp.node ast/used_vars.h.node ast/well_sorted.h.node ast/act_cache.h.node ast/rewriter/var_subst.h.node ast/normal_forms/name_exprs.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/normal_forms/nnf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_18) $(CXX_OUT_FLAG)ast/normal_forms/nnf$(OBJ_EXT) ../src/ast/normal_forms/nnf.cpp
ast/normal_forms/pull_quant.h.node: ../src/ast/normal_forms/pull_quant.h ast/ast.h.node
	@echo done > ast/normal_forms/pull_quant.h.node
ast/normal_forms/pull_quant$(OBJ_EXT): ../src/ast/normal_forms/pull_quant.cpp ast/normal_forms/pull_quant.h.node ast/rewriter/var_subst.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node
	@echo src/ast/normal_forms/pull_quant.cpp
	@$(CXX) $(CXXFLAGS) $(includes_18) $(CXX_OUT_FLAG)ast/normal_forms/pull_quant$(OBJ_EXT) ../src/ast/normal_forms/pull_quant.cpp
ast/normal_forms/elim_term_ite.h.node: ../src/ast/normal_forms/elim_term_ite.h ast/normal_forms/defined_names.h.node ast/rewriter/rewriter.h.node ast/justified_expr.h.node
	@echo done > ast/normal_forms/elim_term_ite.h.node
ast/normal_forms/elim_term_ite$(OBJ_EXT): ../src/ast/normal_forms/elim_term_ite.cpp ast/normal_forms/elim_term_ite.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/normal_forms/elim_term_ite.cpp
	@$(CXX) $(CXXFLAGS) $(includes_18) $(CXX_OUT_FLAG)ast/normal_forms/elim_term_ite$(OBJ_EXT) ../src/ast/normal_forms/elim_term_ite.cpp
ast/normal_forms/normal_forms$(LIB_EXT): ast/normal_forms/defined_names$(OBJ_EXT) ast/normal_forms/name_exprs$(OBJ_EXT) ast/normal_forms/nnf$(OBJ_EXT) ast/normal_forms/pull_quant$(OBJ_EXT) ast/normal_forms/elim_term_ite$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/normal_forms/normal_forms$(LIB_EXT) ast/normal_forms/defined_names$(OBJ_EXT) ast/normal_forms/name_exprs$(OBJ_EXT) ast/normal_forms/nnf$(OBJ_EXT) ast/normal_forms/pull_quant$(OBJ_EXT) ast/normal_forms/elim_term_ite$(OBJ_EXT)
normal_forms: ast/normal_forms/normal_forms$(LIB_EXT)

includes_19 = -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src
model/func_interp.h.node: ../src/model/func_interp.h ast/ast.h.node ast/ast_translation.h.node
	@echo done > model/func_interp.h.node
model/model_core.h.node: ../src/model/model_core.h ast/ast.h.node util/obj_hashtable.h.node model/func_interp.h.node
	@echo done > model/model_core.h.node
model/model_smt2_pp.h.node: ../src/model/model_smt2_pp.h ast/ast_printer.h.node model/model_core.h.node
	@echo done > model/model_smt2_pp.h.node
model/model_smt2_pp$(OBJ_EXT): ../src/model/model_smt2_pp.cpp model/model_smt2_pp.h.node ast/ast_smt2_pp.h.node ast/func_decl_dependencies.h.node ast/recfun_decl_plugin.h.node ast/pp.h.node
	@echo src/model/model_smt2_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_smt2_pp$(OBJ_EXT) ../src/model/model_smt2_pp.cpp
model/model_core$(OBJ_EXT): ../src/model/model_core.cpp model/model_core.h.node ast/ast_pp.h.node
	@echo src/model/model_core.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_core$(OBJ_EXT) ../src/model/model_core.cpp
model/value_factory.h.node: ../src/model/value_factory.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > model/value_factory.h.node
model/value_factory$(OBJ_EXT): ../src/model/value_factory.cpp model/value_factory.h.node ast/ast_pp.h.node
	@echo src/model/value_factory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/value_factory$(OBJ_EXT) ../src/model/value_factory.cpp
util/ref.h.node: ../src/util/ref.h
	@echo done > util/ref.h.node
util/plugin_manager.h.node: ../src/util/plugin_manager.h util/util.h.node
	@echo done > util/plugin_manager.h.node
model/model_evaluator.h.node: ../src/model/model_evaluator.h ast/ast.h.node ast/rewriter/rewriter_types.h.node util/params.h.node
	@echo done > model/model_evaluator.h.node
model/model.h.node: ../src/model/model.h util/ref.h.node util/vector.h.node ast/ast_translation.h.node util/plugin_manager.h.node model/model_core.h.node model/model_evaluator.h.node model/value_factory.h.node
	@echo done > model/model.h.node
model/model_evaluator_params.hpp.node: ../src/model/model_evaluator_params.hpp util/params.h.node util/gparams.h.node
	@echo done > model/model_evaluator_params.hpp.node
model/model_v2_pp.h.node: ../src/model/model_v2_pp.h
	@echo done > model/model_v2_pp.h.node
model/model_evaluator$(OBJ_EXT): ../src/model/model_evaluator.cpp ast/ast_pp.h.node ast/ast_util.h.node ast/rewriter/rewriter_types.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/bv_rewriter.h.node ast/rewriter/pb_rewriter.h.node ast/rewriter/seq_rewriter.h.node ast/rewriter/datatype_rewriter.h.node ast/rewriter/array_rewriter.h.node ast/rewriter/fpa_rewriter.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node model/model_smt2_pp.h.node model/model.h.node model/model_evaluator_params.hpp.node model/model_evaluator.h.node model/model_v2_pp.h.node
	@echo src/model/model_evaluator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_evaluator$(OBJ_EXT) ../src/model/model_evaluator.cpp
model/model2expr.h.node: ../src/model/model2expr.h model/model.h.node
	@echo done > model/model2expr.h.node
model/model2expr$(OBJ_EXT): ../src/model/model2expr.cpp model/model2expr.h.node ast/for_each_ast.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/var_subst.h.node
	@echo src/model/model2expr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model2expr$(OBJ_EXT) ../src/model/model2expr.cpp
model/struct_factory.h.node: ../src/model/struct_factory.h model/value_factory.h.node util/obj_hashtable.h.node
	@echo done > model/struct_factory.h.node
model/struct_factory$(OBJ_EXT): ../src/model/struct_factory.cpp model/struct_factory.h.node model/model_core.h.node
	@echo src/model/struct_factory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/struct_factory$(OBJ_EXT) ../src/model/struct_factory.cpp
model/array_factory.h.node: ../src/model/array_factory.h model/struct_factory.h.node
	@echo done > model/array_factory.h.node
model/array_factory$(OBJ_EXT): ../src/model/array_factory.cpp ast/array_decl_plugin.h.node ast/ast_pp.h.node model/func_interp.h.node model/model_core.h.node model/array_factory.h.node
	@echo src/model/array_factory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/array_factory$(OBJ_EXT) ../src/model/array_factory.cpp
model/datatype_factory.h.node: ../src/model/datatype_factory.h model/struct_factory.h.node ast/datatype_decl_plugin.h.node
	@echo done > model/datatype_factory.h.node
model/datatype_factory$(OBJ_EXT): ../src/model/datatype_factory.cpp model/datatype_factory.h.node model/model_core.h.node ast/ast_pp.h.node ast/expr_functors.h.node
	@echo src/model/datatype_factory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/datatype_factory$(OBJ_EXT) ../src/model/datatype_factory.cpp
util/obj_pair_hashtable.h.node: ../src/util/obj_pair_hashtable.h util/hash.h.node util/hashtable.h.node
	@echo done > util/obj_pair_hashtable.h.node
util/backtrackable_set.h.node: ../src/util/backtrackable_set.h util/vector.h.node
	@echo done > util/backtrackable_set.h.node
model/model_macro_solver.h.node: ../src/model/model_macro_solver.h util/obj_pair_hashtable.h.node util/backtrackable_set.h.node ast/macros/quantifier_macro_info.h.node model/model_core.h.node
	@echo done > model/model_macro_solver.h.node
model/model_macro_solver$(OBJ_EXT): ../src/model/model_macro_solver.cpp ast/for_each_expr.h.node ast/ast_pp.h.node model/model_macro_solver.h.node model/model_core.h.node
	@echo src/model/model_macro_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_macro_solver$(OBJ_EXT) ../src/model/model_macro_solver.cpp
model/func_interp$(OBJ_EXT): ../src/model/func_interp.cpp util/obj_hashtable.h.node ast/rewriter/var_subst.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node ast/ast_util.h.node model/func_interp.h.node ast/array_decl_plugin.h.node
	@echo src/model/func_interp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/func_interp$(OBJ_EXT) ../src/model/func_interp.cpp
model/model_pp.h.node: ../src/model/model_pp.h
	@echo done > model/model_pp.h.node
ast/used_symbols.h.node: ../src/ast/used_symbols.h ast/ast.h.node util/hashtable.h.node util/obj_hashtable.h.node
	@echo done > ast/used_symbols.h.node
model/model_pp$(OBJ_EXT): ../src/model/model_pp.cpp model/model_pp.h.node model/model_core.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node ast/used_symbols.h.node ast/pp.h.node
	@echo src/model/model_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_pp$(OBJ_EXT) ../src/model/model_pp.cpp
model/model_v2_pp$(OBJ_EXT): ../src/model/model_v2_pp.cpp model/model_v2_pp.h.node model/model_core.h.node ast/ast_pp.h.node
	@echo src/model/model_v2_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_v2_pp$(OBJ_EXT) ../src/model/model_v2_pp.cpp
model/model_params.hpp.node: ../src/model/model_params.hpp util/params.h.node util/gparams.h.node
	@echo done > model/model_params.hpp.node
model/seq_factory.h.node: ../src/model/seq_factory.h ast/seq_decl_plugin.h.node model/model_core.h.node
	@echo done > model/seq_factory.h.node
model/numeral_factory.h.node: ../src/model/numeral_factory.h ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node model/value_factory.h.node
	@echo done > model/numeral_factory.h.node
model/fpa_factory.h.node: ../src/model/fpa_factory.h ast/seq_decl_plugin.h.node model/model_core.h.node model/value_factory.h.node
	@echo done > model/fpa_factory.h.node
model/model$(OBJ_EXT): ../src/model/model.cpp ast/ast.h.node util/top_sort.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/rewriter/var_subst.h.node ast/rewriter/th_rewriter.h.node ast/array_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/well_sorted.h.node ast/used_symbols.h.node ast/for_each_expr.h.node ast/for_each_ast.h.node model/model.h.node model/model_params.hpp.node model/model_evaluator.h.node model/array_factory.h.node model/value_factory.h.node model/seq_factory.h.node model/datatype_factory.h.node model/numeral_factory.h.node model/fpa_factory.h.node
	@echo src/model/model.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model$(OBJ_EXT) ../src/model/model.cpp
model/model_implicant.h.node: ../src/model/model_implicant.h ast/ast.h.node ast/ast_pp.h.node util/obj_hashtable.h.node util/ref_vector.h.node util/trace.h.node util/vector.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node
	@echo done > model/model_implicant.h.node
ast/scoped_proof.h.node: ../src/ast/scoped_proof.h ast/ast.h.node
	@echo done > ast/scoped_proof.h.node
model/model_implicant$(OBJ_EXT): ../src/model/model_implicant.cpp ast/array_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node ast/for_each_expr.h.node model/model.h.node util/ref_vector.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node util/util.h.node model/model_implicant.h.node ast/arith_decl_plugin.h.node ast/rewriter/expr_replacer.h.node model/model_smt2_pp.h.node ast/rewriter/poly_rewriter.h.node ast/rewriter/poly_rewriter_def.h.node ast/rewriter/arith_rewriter.h.node ast/scoped_proof.h.node
	@echo src/model/model_implicant.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/model_implicant$(OBJ_EXT) ../src/model/model_implicant.cpp
model/numeral_factory$(OBJ_EXT): ../src/model/numeral_factory.cpp ast/ast_pp.h.node model/numeral_factory.h.node
	@echo src/model/numeral_factory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_19) $(CXX_OUT_FLAG)model/numeral_factory$(OBJ_EXT) ../src/model/numeral_factory.cpp
model/model$(LIB_EXT): model/model_smt2_pp$(OBJ_EXT) model/model_core$(OBJ_EXT) model/value_factory$(OBJ_EXT) model/model_evaluator$(OBJ_EXT) model/model2expr$(OBJ_EXT) model/struct_factory$(OBJ_EXT) model/array_factory$(OBJ_EXT) model/datatype_factory$(OBJ_EXT) model/model_macro_solver$(OBJ_EXT) model/func_interp$(OBJ_EXT) model/model_pp$(OBJ_EXT) model/model_v2_pp$(OBJ_EXT) model/model$(OBJ_EXT) model/model_implicant$(OBJ_EXT) model/numeral_factory$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)model/model$(LIB_EXT) model/model_smt2_pp$(OBJ_EXT) model/model_core$(OBJ_EXT) model/value_factory$(OBJ_EXT) model/model_evaluator$(OBJ_EXT) model/model2expr$(OBJ_EXT) model/struct_factory$(OBJ_EXT) model/array_factory$(OBJ_EXT) model/datatype_factory$(OBJ_EXT) model/model_macro_solver$(OBJ_EXT) model/func_interp$(OBJ_EXT) model/model_pp$(OBJ_EXT) model/model_v2_pp$(OBJ_EXT) model/model$(OBJ_EXT) model/model_implicant$(OBJ_EXT) model/numeral_factory$(OBJ_EXT)
model: model/model$(LIB_EXT)

includes_20 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src
tactic/sine_filter.h.node: ../src/tactic/sine_filter.h util/params.h.node
	@echo done > tactic/sine_filter.h.node
tactic/converter.h.node: ../src/tactic/converter.h util/vector.h.node util/ref.h.node ast/ast_translation.h.node
	@echo done > tactic/converter.h.node
tactic/model_converter.h.node: ../src/tactic/model_converter.h util/ref.h.node ast/ast_pp_util.h.node model/model.h.node tactic/converter.h.node
	@echo done > tactic/model_converter.h.node
tactic/proof_converter.h.node: ../src/tactic/proof_converter.h ast/ast.h.node util/ref.h.node tactic/converter.h.node
	@echo done > tactic/proof_converter.h.node
tactic/dependency_converter.h.node: ../src/tactic/dependency_converter.h util/ref.h.node ast/ast_pp_util.h.node model/model.h.node tactic/converter.h.node
	@echo done > tactic/dependency_converter.h.node
tactic/goal.h.node: ../src/tactic/goal.h ast/ast.h.node ast/ast_translation.h.node ast/ast_printer.h.node ast/for_each_expr.h.node util/ref.h.node util/ref_vector.h.node util/ref_buffer.h.node tactic/model_converter.h.node tactic/proof_converter.h.node tactic/dependency_converter.h.node
	@echo done > tactic/goal.h.node
tactic/tactic_exception.h.node: ../src/tactic/tactic_exception.h util/z3_exception.h.node util/common_msgs.h.node
	@echo done > tactic/tactic_exception.h.node
tactic/tactic.h.node: ../src/tactic/tactic.h tactic/goal.h.node util/params.h.node util/statistics.h.node tactic/tactic_exception.h.node util/lbool.h.node
	@echo done > tactic/tactic.h.node
tactic/probe.h.node: ../src/tactic/probe.h tactic/goal.h.node
	@echo done > tactic/probe.h.node
tactic/tactical.h.node: ../src/tactic/tactical.h tactic/tactic.h.node tactic/probe.h.node
	@echo done > tactic/tactical.h.node
tactic/generic_model_converter.h.node: ../src/tactic/generic_model_converter.h tactic/model_converter.h.node
	@echo done > tactic/generic_model_converter.h.node
tactic/sine_filter$(OBJ_EXT): ../src/tactic/sine_filter.cpp tactic/sine_filter.h.node tactic/tactical.h.node tactic/generic_model_converter.h.node ast/datatype_decl_plugin.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node ast/ast_util.h.node util/obj_pair_hashtable.h.node ast/ast_pp.h.node
	@echo src/tactic/sine_filter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/sine_filter$(OBJ_EXT) ../src/tactic/sine_filter.cpp
tactic/goal_util.h.node: ../src/tactic/goal_util.h
	@echo done > tactic/goal_util.h.node
tactic/goal_util$(OBJ_EXT): ../src/tactic/goal_util.cpp tactic/goal_util.h.node tactic/goal.h.node
	@echo src/tactic/goal_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/goal_util$(OBJ_EXT) ../src/tactic/goal_util.cpp
tactic/tactic$(OBJ_EXT): ../src/tactic/tactic.cpp tactic/tactic.h.node tactic/probe.h.node util/stopwatch.h.node model/model_v2_pp.h.node
	@echo src/tactic/tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/tactic$(OBJ_EXT) ../src/tactic/tactic.cpp
tactic/replace_proof_converter.h.node: ../src/tactic/replace_proof_converter.h tactic/proof_converter.h.node
	@echo done > tactic/replace_proof_converter.h.node
tactic/equiv_proof_converter.h.node: ../src/tactic/equiv_proof_converter.h tactic/replace_proof_converter.h.node
	@echo done > tactic/equiv_proof_converter.h.node
tactic/equiv_proof_converter$(OBJ_EXT): ../src/tactic/equiv_proof_converter.cpp tactic/equiv_proof_converter.h.node ast/ast_pp.h.node ast/scoped_proof.h.node
	@echo src/tactic/equiv_proof_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/equiv_proof_converter$(OBJ_EXT) ../src/tactic/equiv_proof_converter.cpp
tactic/probe$(OBJ_EXT): ../src/tactic/probe.cpp tactic/probe.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node tactic/goal_util.h.node ast/rewriter/bv_rewriter.h.node
	@echo src/tactic/probe.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/probe$(OBJ_EXT) ../src/tactic/probe.cpp
tactic/goal_num_occurs.h.node: ../src/tactic/goal_num_occurs.h ast/num_occurs.h.node
	@echo done > tactic/goal_num_occurs.h.node
tactic/goal_num_occurs$(OBJ_EXT): ../src/tactic/goal_num_occurs.cpp tactic/goal_num_occurs.h.node tactic/goal.h.node
	@echo src/tactic/goal_num_occurs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/goal_num_occurs$(OBJ_EXT) ../src/tactic/goal_num_occurs.cpp
tactic/horn_subsume_model_converter.h.node: ../src/tactic/horn_subsume_model_converter.h tactic/model_converter.h.node ast/rewriter/th_rewriter.h.node
	@echo done > tactic/horn_subsume_model_converter.h.node
tactic/horn_subsume_model_converter$(OBJ_EXT): ../src/tactic/horn_subsume_model_converter.cpp tactic/horn_subsume_model_converter.h.node ast/rewriter/var_subst.h.node ast/ast_pp.h.node model/model_smt2_pp.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node ast/for_each_expr.h.node ast/well_sorted.h.node
	@echo src/tactic/horn_subsume_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/horn_subsume_model_converter$(OBJ_EXT) ../src/tactic/horn_subsume_model_converter.cpp
tactic/goal$(OBJ_EXT): ../src/tactic/goal.cpp ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/for_each_expr.h.node ast/well_sorted.h.node ast/display_dimacs.h.node tactic/goal.h.node
	@echo src/tactic/goal.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/goal$(OBJ_EXT) ../src/tactic/goal.cpp
tactic/goal_shared_occs.h.node: ../src/tactic/goal_shared_occs.h tactic/goal.h.node ast/shared_occs.h.node
	@echo done > tactic/goal_shared_occs.h.node
tactic/goal_shared_occs$(OBJ_EXT): ../src/tactic/goal_shared_occs.cpp tactic/goal_shared_occs.h.node
	@echo src/tactic/goal_shared_occs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/goal_shared_occs$(OBJ_EXT) ../src/tactic/goal_shared_occs.cpp
tactic/generic_model_converter$(OBJ_EXT): ../src/tactic/generic_model_converter.cpp ast/ast_pp.h.node ast/for_each_expr.h.node ast/ast_util.h.node ast/occurs.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/th_rewriter.h.node tactic/generic_model_converter.h.node model/model_v2_pp.h.node model/model_evaluator.h.node
	@echo src/tactic/generic_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/generic_model_converter$(OBJ_EXT) ../src/tactic/generic_model_converter.cpp
tactic/model_converter$(OBJ_EXT): ../src/tactic/model_converter.cpp tactic/model_converter.h.node model/model_v2_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/model_converter$(OBJ_EXT) ../src/tactic/model_converter.cpp
tactic/dependency_converter$(OBJ_EXT): ../src/tactic/dependency_converter.cpp tactic/dependency_converter.h.node tactic/goal.h.node
	@echo src/tactic/dependency_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/dependency_converter$(OBJ_EXT) ../src/tactic/dependency_converter.cpp
tactic/replace_proof_converter$(OBJ_EXT): ../src/tactic/replace_proof_converter.cpp tactic/replace_proof_converter.h.node ast/expr_functors.h.node ast/ast_pp.h.node ast/for_each_expr.h.node
	@echo src/tactic/replace_proof_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/replace_proof_converter$(OBJ_EXT) ../src/tactic/replace_proof_converter.cpp
util/cancel_eh.h.node: ../src/util/cancel_eh.h util/event_handler.h.node
	@echo done > util/cancel_eh.h.node
tactic/tactical$(OBJ_EXT): ../src/tactic/tactical.cpp util/scoped_timer.h.node util/cancel_eh.h.node util/scoped_ptr_vector.h.node tactic/tactical.h.node
	@echo src/tactic/tactical.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/tactical$(OBJ_EXT) ../src/tactic/tactical.cpp
tactic/proof_converter$(OBJ_EXT): ../src/tactic/proof_converter.cpp tactic/proof_converter.h.node tactic/goal.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/proof_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_20) $(CXX_OUT_FLAG)tactic/proof_converter$(OBJ_EXT) ../src/tactic/proof_converter.cpp
tactic/tactic$(LIB_EXT): tactic/sine_filter$(OBJ_EXT) tactic/goal_util$(OBJ_EXT) tactic/tactic$(OBJ_EXT) tactic/equiv_proof_converter$(OBJ_EXT) tactic/probe$(OBJ_EXT) tactic/goal_num_occurs$(OBJ_EXT) tactic/horn_subsume_model_converter$(OBJ_EXT) tactic/goal$(OBJ_EXT) tactic/goal_shared_occs$(OBJ_EXT) tactic/generic_model_converter$(OBJ_EXT) tactic/model_converter$(OBJ_EXT) tactic/dependency_converter$(OBJ_EXT) tactic/replace_proof_converter$(OBJ_EXT) tactic/tactical$(OBJ_EXT) tactic/proof_converter$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/tactic$(LIB_EXT) tactic/sine_filter$(OBJ_EXT) tactic/goal_util$(OBJ_EXT) tactic/tactic$(OBJ_EXT) tactic/equiv_proof_converter$(OBJ_EXT) tactic/probe$(OBJ_EXT) tactic/goal_num_occurs$(OBJ_EXT) tactic/horn_subsume_model_converter$(OBJ_EXT) tactic/goal$(OBJ_EXT) tactic/goal_shared_occs$(OBJ_EXT) tactic/generic_model_converter$(OBJ_EXT) tactic/model_converter$(OBJ_EXT) tactic/dependency_converter$(OBJ_EXT) tactic/replace_proof_converter$(OBJ_EXT) tactic/tactical$(OBJ_EXT) tactic/proof_converter$(OBJ_EXT)
tactic: tactic/tactic$(LIB_EXT)

includes_21 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src
ast/substitution/expr_offset.h.node: ../src/ast/substitution/expr_offset.h ast/ast.h.node
	@echo done > ast/substitution/expr_offset.h.node
ast/substitution/expr_offset_map.h.node: ../src/ast/substitution/expr_offset_map.h ast/substitution/expr_offset.h.node util/vector.h.node
	@echo done > ast/substitution/expr_offset_map.h.node
ast/substitution/var_offset_map.h.node: ../src/ast/substitution/var_offset_map.h ast/ast.h.node util/vector.h.node
	@echo done > ast/substitution/var_offset_map.h.node
ast/substitution/substitution.h.node: ../src/ast/substitution/substitution.h ast/substitution/expr_offset_map.h.node ast/substitution/var_offset_map.h.node ast/ast_pp.h.node
	@echo done > ast/substitution/substitution.h.node
ast/substitution/substitution_tree.h.node: ../src/ast/substitution/substitution_tree.h ast/ast.h.node ast/substitution/substitution.h.node
	@echo done > ast/substitution/substitution_tree.h.node
ast/substitution/substitution_tree$(OBJ_EXT): ../src/ast/substitution/substitution_tree.cpp ast/substitution/substitution_tree.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/substitution/substitution_tree.cpp
	@$(CXX) $(CXXFLAGS) $(includes_21) $(CXX_OUT_FLAG)ast/substitution/substitution_tree$(OBJ_EXT) ../src/ast/substitution/substitution_tree.cpp
ast/substitution/substitution$(OBJ_EXT): ../src/ast/substitution/substitution.cpp ast/substitution/substitution.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/rewriter/rewriter.h.node
	@echo src/ast/substitution/substitution.cpp
	@$(CXX) $(CXXFLAGS) $(includes_21) $(CXX_OUT_FLAG)ast/substitution/substitution$(OBJ_EXT) ../src/ast/substitution/substitution.cpp
ast/substitution/unifier.h.node: ../src/ast/substitution/unifier.h ast/ast.h.node ast/substitution/substitution.h.node
	@echo done > ast/substitution/unifier.h.node
ast/substitution/unifier$(OBJ_EXT): ../src/ast/substitution/unifier.cpp ast/substitution/unifier.h.node ast/ast_pp.h.node
	@echo src/ast/substitution/unifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_21) $(CXX_OUT_FLAG)ast/substitution/unifier$(OBJ_EXT) ../src/ast/substitution/unifier.cpp
ast/substitution/matcher.h.node: ../src/ast/substitution/matcher.h ast/substitution/substitution.h.node util/hashtable.h.node
	@echo done > ast/substitution/matcher.h.node
ast/substitution/matcher$(OBJ_EXT): ../src/ast/substitution/matcher.cpp ast/substitution/matcher.h.node
	@echo src/ast/substitution/matcher.cpp
	@$(CXX) $(CXXFLAGS) $(includes_21) $(CXX_OUT_FLAG)ast/substitution/matcher$(OBJ_EXT) ../src/ast/substitution/matcher.cpp
ast/substitution/substitution$(LIB_EXT): ast/substitution/substitution_tree$(OBJ_EXT) ast/substitution/substitution$(OBJ_EXT) ast/substitution/unifier$(OBJ_EXT) ast/substitution/matcher$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/substitution/substitution$(LIB_EXT) ast/substitution/substitution_tree$(OBJ_EXT) ast/substitution/substitution$(OBJ_EXT) ast/substitution/unifier$(OBJ_EXT) ast/substitution/matcher$(OBJ_EXT)
substitution: ast/substitution/substitution$(LIB_EXT)

includes_22 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src
parsers/util/simple_parser.h.node: ../src/parsers/util/simple_parser.h ast/ast.h.node util/map.h.node
	@echo done > parsers/util/simple_parser.h.node
parsers/util/scanner.h.node: ../src/parsers/util/scanner.h ast/ast.h.node
	@echo done > parsers/util/scanner.h.node
parsers/util/simple_parser$(OBJ_EXT): ../src/parsers/util/simple_parser.cpp parsers/util/simple_parser.h.node util/warning.h.node parsers/util/scanner.h.node
	@echo src/parsers/util/simple_parser.cpp
	@$(CXX) $(CXXFLAGS) $(includes_22) $(CXX_OUT_FLAG)parsers/util/simple_parser$(OBJ_EXT) ../src/parsers/util/simple_parser.cpp
parsers/util/pattern_validation.h.node: ../src/parsers/util/pattern_validation.h ast/ast.h.node util/uint_set.h.node util/vector.h.node
	@echo done > parsers/util/pattern_validation.h.node
parsers/util/pattern_validation$(OBJ_EXT): ../src/parsers/util/pattern_validation.cpp parsers/util/pattern_validation.h.node ast/for_each_expr.h.node util/warning.h.node ast/ast_pp.h.node
	@echo src/parsers/util/pattern_validation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_22) $(CXX_OUT_FLAG)parsers/util/pattern_validation$(OBJ_EXT) ../src/parsers/util/pattern_validation.cpp
parsers/util/cost_parser.h.node: ../src/parsers/util/cost_parser.h parsers/util/simple_parser.h.node ast/arith_decl_plugin.h.node
	@echo done > parsers/util/cost_parser.h.node
parsers/util/cost_parser$(OBJ_EXT): ../src/parsers/util/cost_parser.cpp parsers/util/cost_parser.h.node
	@echo src/parsers/util/cost_parser.cpp
	@$(CXX) $(CXXFLAGS) $(includes_22) $(CXX_OUT_FLAG)parsers/util/cost_parser$(OBJ_EXT) ../src/parsers/util/cost_parser.cpp
parsers/util/scanner$(OBJ_EXT): ../src/parsers/util/scanner.cpp parsers/util/scanner.h.node
	@echo src/parsers/util/scanner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_22) $(CXX_OUT_FLAG)parsers/util/scanner$(OBJ_EXT) ../src/parsers/util/scanner.cpp
parsers/util/parser_util$(LIB_EXT): parsers/util/simple_parser$(OBJ_EXT) parsers/util/pattern_validation$(OBJ_EXT) parsers/util/cost_parser$(OBJ_EXT) parsers/util/scanner$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)parsers/util/parser_util$(LIB_EXT) parsers/util/simple_parser$(OBJ_EXT) parsers/util/pattern_validation$(OBJ_EXT) parsers/util/cost_parser$(OBJ_EXT) parsers/util/scanner$(OBJ_EXT)
parser_util: parsers/util/parser_util$(LIB_EXT)

includes_23 = -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src
ast/proofs/proof_checker.h.node: ../src/ast/proofs/proof_checker.h ast/ast.h.node util/map.h.node
	@echo done > ast/proofs/proof_checker.h.node
ast/proofs/proof_checker$(OBJ_EXT): ../src/ast/proofs/proof_checker.cpp ast/proofs/proof_checker.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_smt_pp.h.node ast/arith_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/var_subst.h.node
	@echo src/ast/proofs/proof_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_23) $(CXX_OUT_FLAG)ast/proofs/proof_checker$(OBJ_EXT) ../src/ast/proofs/proof_checker.cpp
ast/proofs/proof_utils.h.node: ../src/ast/proofs/proof_utils.h ast/ast.h.node ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node ast/proofs/proof_checker.h.node
	@echo done > ast/proofs/proof_utils.h.node
ast/proofs/proof_utils$(OBJ_EXT): ../src/ast/proofs/proof_utils.cpp ast/ast_util.h.node ast/ast_pp.h.node ast/proofs/proof_utils.h.node ast/proofs/proof_checker.h.node ast/rewriter/var_subst.h.node util/container_util.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/proofs/proof_utils.cpp
	@$(CXX) $(CXXFLAGS) $(includes_23) $(CXX_OUT_FLAG)ast/proofs/proof_utils$(OBJ_EXT) ../src/ast/proofs/proof_utils.cpp
ast/proofs/proofs$(LIB_EXT): ast/proofs/proof_checker$(OBJ_EXT) ast/proofs/proof_utils$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/proofs/proofs$(LIB_EXT) ast/proofs/proof_checker$(OBJ_EXT) ast/proofs/proof_utils$(OBJ_EXT)
proofs: ast/proofs/proofs$(LIB_EXT)

includes_24 = -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/tactic -I../src/ast/proofs -I../src
solver/tactic2solver.h.node: ../src/solver/tactic2solver.h util/params.h.node
	@echo done > solver/tactic2solver.h.node
solver/check_sat_result.h.node: ../src/solver/check_sat_result.h model/model.h.node util/lbool.h.node util/statistics.h.node util/event_handler.h.node util/timer.h.node tactic/model_converter.h.node
	@echo done > solver/check_sat_result.h.node
solver/progress_callback.h.node: ../src/solver/progress_callback.h
	@echo done > solver/progress_callback.h.node
solver/solver.h.node: ../src/solver/solver.h solver/check_sat_result.h.node solver/progress_callback.h.node util/params.h.node
	@echo done > solver/solver.h.node
solver/solver_na2as.h.node: ../src/solver/solver_na2as.h solver/solver.h.node
	@echo done > solver/solver_na2as.h.node
solver/mus.h.node: ../src/solver/mus.h
	@echo done > solver/mus.h.node
solver/tactic2solver$(OBJ_EXT): ../src/solver/tactic2solver.cpp ast/ast_translation.h.node ast/ast_pp.h.node tactic/tactic.h.node solver/tactic2solver.h.node solver/solver_na2as.h.node solver/mus.h.node
	@echo src/solver/tactic2solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/tactic2solver$(OBJ_EXT) ../src/solver/tactic2solver.cpp
solver/mus$(OBJ_EXT): ../src/solver/mus.cpp solver/solver.h.node solver/mus.h.node ast/ast_pp.h.node ast/ast_util.h.node model/model_evaluator.h.node
	@echo src/solver/mus.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/mus$(OBJ_EXT) ../src/solver/mus.cpp
solver/check_sat_result$(OBJ_EXT): ../src/solver/check_sat_result.cpp solver/check_sat_result.h.node
	@echo src/solver/check_sat_result.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/check_sat_result$(OBJ_EXT) ../src/solver/check_sat_result.cpp
solver/solver2tactic.h.node: ../src/solver/solver2tactic.h tactic/tactic.h.node tactic/generic_model_converter.h.node
	@echo done > solver/solver2tactic.h.node
solver/parallel_tactic.h.node: ../src/solver/parallel_tactic.h
	@echo done > solver/parallel_tactic.h.node
solver/parallel_params.hpp.node: ../src/solver/parallel_params.hpp util/params.h.node util/gparams.h.node
	@echo done > solver/parallel_params.hpp.node
solver/parallel_tactic$(OBJ_EXT): ../src/solver/parallel_tactic.cpp util/scoped_ptr_vector.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/ast_translation.h.node solver/solver.h.node solver/solver2tactic.h.node tactic/tactic.h.node tactic/tactical.h.node solver/parallel_tactic.h.node solver/parallel_params.hpp.node
	@echo src/solver/parallel_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/parallel_tactic$(OBJ_EXT) ../src/solver/parallel_tactic.cpp
solver/smt_logics.h.node: ../src/solver/smt_logics.h
	@echo done > solver/smt_logics.h.node
solver/smt_logics$(OBJ_EXT): ../src/solver/smt_logics.cpp util/symbol.h.node solver/smt_logics.h.node
	@echo src/solver/smt_logics.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/smt_logics$(OBJ_EXT) ../src/solver/smt_logics.cpp
solver/solver2tactic$(OBJ_EXT): ../src/solver/solver2tactic.cpp solver/solver.h.node tactic/tactic.h.node tactic/generic_model_converter.h.node solver/solver2tactic.h.node ast/ast_util.h.node
	@echo src/solver/solver2tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/solver2tactic$(OBJ_EXT) ../src/solver/solver2tactic.cpp
solver/solver_params.hpp.node: ../src/solver/solver_params.hpp util/params.h.node util/gparams.h.node
	@echo done > solver/solver_params.hpp.node
solver/solver$(OBJ_EXT): ../src/solver/solver.cpp util/common_msgs.h.node util/stopwatch.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_pp_util.h.node ast/display_dimacs.h.node tactic/model_converter.h.node solver/solver.h.node solver/solver_params.hpp.node model/model_evaluator.h.node model/model_params.hpp.node
	@echo src/solver/solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/solver$(OBJ_EXT) ../src/solver/solver.cpp
solver/combined_solver_params.hpp.node: ../src/solver/combined_solver_params.hpp util/params.h.node util/gparams.h.node
	@echo done > solver/combined_solver_params.hpp.node
solver/combined_solver$(OBJ_EXT): ../src/solver/combined_solver.cpp util/scoped_timer.h.node util/common_msgs.h.node ast/ast_pp.h.node solver/solver.h.node solver/combined_solver_params.hpp.node
	@echo src/solver/combined_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/combined_solver$(OBJ_EXT) ../src/solver/combined_solver.cpp
solver/check_logic.h.node: ../src/solver/check_logic.h ast/ast.h.node
	@echo done > solver/check_logic.h.node
solver/check_logic$(OBJ_EXT): ../src/solver/check_logic.cpp solver/check_logic.h.node solver/smt_logics.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/ast_pp.h.node ast/for_each_expr.h.node
	@echo src/solver/check_logic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/check_logic$(OBJ_EXT) ../src/solver/check_logic.cpp
solver/solver_pool.h.node: ../src/solver/solver_pool.h solver/solver.h.node util/stopwatch.h.node
	@echo done > solver/solver_pool.h.node
solver/solver_pool$(OBJ_EXT): ../src/solver/solver_pool.cpp solver/solver_pool.h.node solver/solver_na2as.h.node ast/proofs/proof_utils.h.node ast/ast_util.h.node
	@echo src/solver/solver_pool.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/solver_pool$(OBJ_EXT) ../src/solver/solver_pool.cpp
solver/solver_na2as$(OBJ_EXT): ../src/solver/solver_na2as.cpp solver/solver_na2as.h.node ast/ast_smt2_pp.h.node
	@echo src/solver/solver_na2as.cpp
	@$(CXX) $(CXXFLAGS) $(includes_24) $(CXX_OUT_FLAG)solver/solver_na2as$(OBJ_EXT) ../src/solver/solver_na2as.cpp
solver/solver$(LIB_EXT): solver/tactic2solver$(OBJ_EXT) solver/mus$(OBJ_EXT) solver/check_sat_result$(OBJ_EXT) solver/parallel_tactic$(OBJ_EXT) solver/smt_logics$(OBJ_EXT) solver/solver2tactic$(OBJ_EXT) solver/solver$(OBJ_EXT) solver/combined_solver$(OBJ_EXT) solver/check_logic$(OBJ_EXT) solver/solver_pool$(OBJ_EXT) solver/solver_na2as$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)solver/solver$(LIB_EXT) solver/tactic2solver$(OBJ_EXT) solver/mus$(OBJ_EXT) solver/check_sat_result$(OBJ_EXT) solver/parallel_tactic$(OBJ_EXT) solver/smt_logics$(OBJ_EXT) solver/solver2tactic$(OBJ_EXT) solver/solver$(OBJ_EXT) solver/combined_solver$(OBJ_EXT) solver/check_logic$(OBJ_EXT) solver/solver_pool$(OBJ_EXT) solver/solver_na2as$(OBJ_EXT)
solver: solver/solver$(LIB_EXT)

includes_25 = -I../src/solver -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/tactic -I../src/ast/proofs -I../src
cmd_context/pdecl.h.node: ../src/cmd_context/pdecl.h ast/ast.h.node util/obj_hashtable.h.node util/dictionary.h.node ast/format.h.node ast/datatype_decl_plugin.h.node
	@echo done > cmd_context/pdecl.h.node
cmd_context/pdecl$(OBJ_EXT): ../src/cmd_context/pdecl.cpp cmd_context/pdecl.h.node ast/datatype_decl_plugin.h.node ast/ast_pp.h.node
	@echo src/cmd_context/pdecl.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/pdecl$(OBJ_EXT) ../src/cmd_context/pdecl.cpp
cmd_context/tactic_cmds.h.node: ../src/cmd_context/tactic_cmds.h ast/ast.h.node util/params.h.node util/cmd_context_types.h.node util/ref.h.node
	@echo done > cmd_context/tactic_cmds.h.node
cmd_context/tactic_manager.h.node: ../src/cmd_context/tactic_manager.h cmd_context/tactic_cmds.h.node util/dictionary.h.node
	@echo done > cmd_context/tactic_manager.h.node
cmd_context/cmd_context.h.node: ../src/cmd_context/cmd_context.h util/stopwatch.h.node util/cmd_context_types.h.node util/event_handler.h.node util/sexpr.h.node util/dictionary.h.node util/scoped_ptr_vector.h.node ast/ast.h.node ast/ast_printer.h.node ast/datatype_decl_plugin.h.node ast/recfun_decl_plugin.h.node ast/rewriter/seq_rewriter.h.node tactic/generic_model_converter.h.node solver/solver.h.node solver/check_logic.h.node solver/progress_callback.h.node cmd_context/pdecl.h.node cmd_context/tactic_manager.h.node params/context_params.h.node
	@echo done > cmd_context/cmd_context.h.node
cmd_context/cmd_util.h.node: ../src/cmd_context/cmd_util.h
	@echo done > cmd_context/cmd_util.h.node
cmd_context/simplify_cmd.h.node: ../src/cmd_context/simplify_cmd.h
	@echo done > cmd_context/simplify_cmd.h.node
cmd_context/eval_cmd.h.node: ../src/cmd_context/eval_cmd.h
	@echo done > cmd_context/eval_cmd.h.node
cmd_context/basic_cmds$(OBJ_EXT): ../src/cmd_context/basic_cmds.cpp util/gparams.h.node util/env_params.h.node util/z3_version.h.node ast/ast_smt_pp.h.node ast/ast_smt2_pp.h.node ast/ast_pp_dot.h.node ast/ast_pp.h.node ast/array_decl_plugin.h.node ast/pp.h.node ast/well_sorted.h.node ast/pp_params.hpp.node model/model_smt2_pp.h.node cmd_context/cmd_context.h.node cmd_context/cmd_util.h.node cmd_context/simplify_cmd.h.node cmd_context/eval_cmd.h.node
	@echo src/cmd_context/basic_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/basic_cmds$(OBJ_EXT) ../src/cmd_context/basic_cmds.cpp
cmd_context/cmd_context_to_goal$(OBJ_EXT): ../src/cmd_context/cmd_context_to_goal.cpp cmd_context/cmd_context.h.node tactic/goal.h.node
	@echo src/cmd_context/cmd_context_to_goal.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/cmd_context_to_goal$(OBJ_EXT) ../src/cmd_context/cmd_context_to_goal.cpp
cmd_context/echo_tactic$(OBJ_EXT): ../src/cmd_context/echo_tactic.cpp tactic/tactic.h.node tactic/probe.h.node cmd_context/cmd_context.h.node
	@echo src/cmd_context/echo_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/echo_tactic$(OBJ_EXT) ../src/cmd_context/echo_tactic.cpp
cmd_context/cmd_util$(OBJ_EXT): ../src/cmd_context/cmd_util.cpp cmd_context/cmd_context.h.node
	@echo src/cmd_context/cmd_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/cmd_util$(OBJ_EXT) ../src/cmd_context/cmd_util.cpp
cmd_context/parametric_cmd.h.node: ../src/cmd_context/parametric_cmd.h util/params.h.node util/symbol.h.node util/string_buffer.h.node util/cmd_context_types.h.node
	@echo done > cmd_context/parametric_cmd.h.node
cmd_context/simplify_cmd$(OBJ_EXT): ../src/cmd_context/simplify_cmd.cpp cmd_context/cmd_context.h.node ast/rewriter/th_rewriter.h.node ast/shared_occs.h.node ast/ast_smt_pp.h.node ast/for_each_expr.h.node cmd_context/parametric_cmd.h.node util/scoped_timer.h.node util/scoped_ctrl_c.h.node util/cancel_eh.h.node
	@echo src/cmd_context/simplify_cmd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/simplify_cmd$(OBJ_EXT) ../src/cmd_context/simplify_cmd.cpp
cmd_context/tactic_manager$(OBJ_EXT): ../src/cmd_context/tactic_manager.cpp cmd_context/tactic_manager.h.node
	@echo src/cmd_context/tactic_manager.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/tactic_manager$(OBJ_EXT) ../src/cmd_context/tactic_manager.cpp
cmd_context/parametric_cmd$(OBJ_EXT): ../src/cmd_context/parametric_cmd.cpp cmd_context/cmd_context.h.node cmd_context/parametric_cmd.h.node
	@echo src/cmd_context/parametric_cmd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/parametric_cmd$(OBJ_EXT) ../src/cmd_context/parametric_cmd.cpp
cmd_context/eval_cmd$(OBJ_EXT): ../src/cmd_context/eval_cmd.cpp cmd_context/cmd_context.h.node model/model_evaluator.h.node cmd_context/parametric_cmd.h.node util/scoped_timer.h.node util/scoped_ctrl_c.h.node util/cancel_eh.h.node
	@echo src/cmd_context/eval_cmd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/eval_cmd$(OBJ_EXT) ../src/cmd_context/eval_cmd.cpp
ast/rewriter/recfun_replace.h.node: ../src/ast/rewriter/recfun_replace.h ast/recfun_decl_plugin.h.node ast/rewriter/expr_safe_replace.h.node
	@echo done > ast/rewriter/recfun_replace.h.node
cmd_context/basic_cmds.h.node: ../src/cmd_context/basic_cmds.h
	@echo done > cmd_context/basic_cmds.h.node
cmd_context/cmd_context$(OBJ_EXT): ../src/cmd_context/cmd_context.cpp util/tptr.h.node util/cancel_eh.h.node util/scoped_ctrl_c.h.node util/dec_ref_util.h.node util/scoped_timer.h.node ast/func_decl_dependencies.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/fpa_decl_plugin.h.node ast/special_relations_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/var_subst.h.node ast/pp.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node ast/decl_collector.h.node ast/well_sorted.h.node ast/for_each_expr.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/recfun_replace.h.node model/model_evaluator.h.node model/model_smt2_pp.h.node model/model_v2_pp.h.node model/model_params.hpp.node tactic/tactic_exception.h.node tactic/generic_model_converter.h.node solver/smt_logics.h.node cmd_context/basic_cmds.h.node cmd_context/cmd_context.h.node
	@echo src/cmd_context/cmd_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/cmd_context$(OBJ_EXT) ../src/cmd_context/cmd_context.cpp
cmd_context/cmd_context_to_goal.h.node: ../src/cmd_context/cmd_context_to_goal.h
	@echo done > cmd_context/cmd_context_to_goal.h.node
cmd_context/echo_tactic.h.node: ../src/cmd_context/echo_tactic.h
	@echo done > cmd_context/echo_tactic.h.node
cmd_context/tactic_cmds$(OBJ_EXT): ../src/cmd_context/tactic_cmds.cpp cmd_context/tactic_cmds.h.node cmd_context/cmd_context.h.node cmd_context/cmd_util.h.node cmd_context/parametric_cmd.h.node util/scoped_timer.h.node util/scoped_ctrl_c.h.node util/cancel_eh.h.node model/model_smt2_pp.h.node ast/ast_smt2_pp.h.node tactic/tactic.h.node tactic/tactical.h.node tactic/probe.h.node solver/check_sat_result.h.node cmd_context/cmd_context_to_goal.h.node cmd_context/echo_tactic.h.node
	@echo src/cmd_context/tactic_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_25) $(CXX_OUT_FLAG)cmd_context/tactic_cmds$(OBJ_EXT) ../src/cmd_context/tactic_cmds.cpp
cmd_context/cmd_context$(LIB_EXT): cmd_context/pdecl$(OBJ_EXT) cmd_context/basic_cmds$(OBJ_EXT) cmd_context/cmd_context_to_goal$(OBJ_EXT) cmd_context/echo_tactic$(OBJ_EXT) cmd_context/cmd_util$(OBJ_EXT) cmd_context/simplify_cmd$(OBJ_EXT) cmd_context/tactic_manager$(OBJ_EXT) cmd_context/parametric_cmd$(OBJ_EXT) cmd_context/eval_cmd$(OBJ_EXT) cmd_context/cmd_context$(OBJ_EXT) cmd_context/tactic_cmds$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)cmd_context/cmd_context$(LIB_EXT) cmd_context/pdecl$(OBJ_EXT) cmd_context/basic_cmds$(OBJ_EXT) cmd_context/cmd_context_to_goal$(OBJ_EXT) cmd_context/echo_tactic$(OBJ_EXT) cmd_context/cmd_util$(OBJ_EXT) cmd_context/simplify_cmd$(OBJ_EXT) cmd_context/tactic_manager$(OBJ_EXT) cmd_context/parametric_cmd$(OBJ_EXT) cmd_context/eval_cmd$(OBJ_EXT) cmd_context/cmd_context$(OBJ_EXT) cmd_context/tactic_cmds$(OBJ_EXT)
cmd_context: cmd_context/cmd_context$(LIB_EXT)

includes_26 = -I../src/cmd_context -I../src/solver -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/tactic -I../src/ast/proofs -I../src/parsers/util -I../src
parsers/smt2/smt2parser.h.node: ../src/parsers/smt2/smt2parser.h cmd_context/cmd_context.h.node
	@echo done > parsers/smt2/smt2parser.h.node
parsers/smt2/smt2scanner.h.node: ../src/parsers/smt2/smt2scanner.h util/symbol.h.node util/vector.h.node util/rational.h.node cmd_context/cmd_context.h.node
	@echo done > parsers/smt2/smt2scanner.h.node
parsers/util/parser_params.hpp.node: ../src/parsers/util/parser_params.hpp util/params.h.node util/gparams.h.node
	@echo done > parsers/util/parser_params.hpp.node
parsers/smt2/smt2parser$(OBJ_EXT): ../src/parsers/smt2/smt2parser.cpp util/stack.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/ast_pp.h.node ast/well_sorted.h.node ast/rewriter/rewriter.h.node ast/rewriter/var_subst.h.node ast/has_free_vars.h.node ast/ast_smt2_pp.h.node parsers/smt2/smt2parser.h.node parsers/smt2/smt2scanner.h.node parsers/util/pattern_validation.h.node parsers/util/parser_params.hpp.node
	@echo src/parsers/smt2/smt2parser.cpp
	@$(CXX) $(CXXFLAGS) $(includes_26) $(CXX_OUT_FLAG)parsers/smt2/smt2parser$(OBJ_EXT) ../src/parsers/smt2/smt2parser.cpp
parsers/smt2/smt2scanner$(OBJ_EXT): ../src/parsers/smt2/smt2scanner.cpp parsers/smt2/smt2scanner.h.node parsers/util/parser_params.hpp.node
	@echo src/parsers/smt2/smt2scanner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_26) $(CXX_OUT_FLAG)parsers/smt2/smt2scanner$(OBJ_EXT) ../src/parsers/smt2/smt2scanner.cpp
parsers/smt2/marshal.h.node: ../src/parsers/smt2/marshal.h ast/ast.h.node
	@echo done > parsers/smt2/marshal.h.node
parsers/smt2/marshal$(OBJ_EXT): ../src/parsers/smt2/marshal.cpp parsers/smt2/marshal.h.node cmd_context/cmd_context.h.node parsers/smt2/smt2parser.h.node util/vector.h.node ast/ast_smt_pp.h.node ast/ast_pp.h.node ast/ast_util.h.node
	@echo src/parsers/smt2/marshal.cpp
	@$(CXX) $(CXXFLAGS) $(includes_26) $(CXX_OUT_FLAG)parsers/smt2/marshal$(OBJ_EXT) ../src/parsers/smt2/marshal.cpp
parsers/smt2/smt2parser$(LIB_EXT): parsers/smt2/smt2parser$(OBJ_EXT) parsers/smt2/smt2scanner$(OBJ_EXT) parsers/smt2/marshal$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)parsers/smt2/smt2parser$(LIB_EXT) parsers/smt2/smt2parser$(OBJ_EXT) parsers/smt2/smt2scanner$(OBJ_EXT) parsers/smt2/marshal$(OBJ_EXT)
smt2parser: parsers/smt2/smt2parser$(LIB_EXT)

includes_27 = -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src
tactic/aig/aig.h.node: ../src/tactic/aig/aig.h ast/ast.h.node tactic/tactic_exception.h.node
	@echo done > tactic/aig/aig.h.node
tactic/aig/aig_tactic$(OBJ_EXT): ../src/tactic/aig/aig_tactic.cpp tactic/tactical.h.node tactic/aig/aig.h.node
	@echo src/tactic/aig/aig_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_27) $(CXX_OUT_FLAG)tactic/aig/aig_tactic$(OBJ_EXT) ../src/tactic/aig/aig_tactic.cpp
tactic/aig/aig$(OBJ_EXT): ../src/tactic/aig/aig.cpp tactic/aig/aig.h.node tactic/goal.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/aig/aig.cpp
	@$(CXX) $(CXXFLAGS) $(includes_27) $(CXX_OUT_FLAG)tactic/aig/aig$(OBJ_EXT) ../src/tactic/aig/aig.cpp
tactic/aig/aig_tactic$(LIB_EXT): tactic/aig/aig_tactic$(OBJ_EXT) tactic/aig/aig$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/aig/aig_tactic$(LIB_EXT) tactic/aig/aig_tactic$(OBJ_EXT) tactic/aig/aig$(OBJ_EXT)
aig_tactic: tactic/aig/aig_tactic$(LIB_EXT)

includes_28 = -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src
ackermannization/ackermannize_bv_tactic.h.node: ../src/ackermannization/ackermannize_bv_tactic.h tactic/tactical.h.node
	@echo done > ackermannization/ackermannize_bv_tactic.h.node
ackermannization/ackr_info.h.node: ../src/ackermannization/ackr_info.h util/ref.h.node util/obj_hashtable.h.node ast/ast.h.node ast/rewriter/expr_replacer.h.node ast/ast_translation.h.node
	@echo done > ackermannization/ackr_info.h.node
ackermannization/ackr_helper.h.node: ../src/ackermannization/ackr_helper.h ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node
	@echo done > ackermannization/ackr_helper.h.node
ackermannization/lackr.h.node: ../src/ackermannization/lackr.h util/lbool.h.node util/util.h.node ast/rewriter/th_rewriter.h.node ast/bv_decl_plugin.h.node model/model.h.node solver/solver.h.node tactic/tactic_exception.h.node tactic/goal.h.node ackermannization/ackr_info.h.node ackermannization/ackr_helper.h.node
	@echo done > ackermannization/lackr.h.node
ackermannization/ackermannize_bv_tactic_params.hpp.node: ../src/ackermannization/ackermannize_bv_tactic_params.hpp util/params.h.node util/gparams.h.node
	@echo done > ackermannization/ackermannize_bv_tactic_params.hpp.node
ackermannization/ackermannize_bv_model_converter.h.node: ../src/ackermannization/ackermannize_bv_model_converter.h tactic/model_converter.h.node ackermannization/ackr_info.h.node
	@echo done > ackermannization/ackermannize_bv_model_converter.h.node
ackermannization/ackermannize_bv_tactic$(OBJ_EXT): ../src/ackermannization/ackermannize_bv_tactic.cpp ackermannization/ackermannize_bv_tactic.h.node tactic/tactical.h.node ackermannization/lackr.h.node model/model_smt2_pp.h.node ackermannization/ackermannize_bv_tactic_params.hpp.node ackermannization/ackermannize_bv_model_converter.h.node
	@echo src/ackermannization/ackermannize_bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/ackermannize_bv_tactic$(OBJ_EXT) ../src/ackermannization/ackermannize_bv_tactic.cpp
ackermannization/ackr_model_converter.h.node: ../src/ackermannization/ackr_model_converter.h tactic/model_converter.h.node ackermannization/ackr_info.h.node
	@echo done > ackermannization/ackr_model_converter.h.node
ackermannization/ackr_model_converter$(OBJ_EXT): ../src/ackermannization/ackr_model_converter.cpp ast/ast_smt2_pp.h.node ast/array_decl_plugin.h.node model/model_evaluator.h.node ackermannization/ackr_model_converter.h.node ackermannization/ackr_info.h.node
	@echo src/ackermannization/ackr_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/ackr_model_converter$(OBJ_EXT) ../src/ackermannization/ackr_model_converter.cpp
ackermannization/lackr_model_constructor.h.node: ../src/ackermannization/lackr_model_constructor.h ast/ast.h.node ackermannization/ackr_info.h.node ackermannization/ackr_helper.h.node model/model.h.node
	@echo done > ackermannization/lackr_model_constructor.h.node
ackermannization/lackr_model_constructor$(OBJ_EXT): ../src/ackermannization/lackr_model_constructor.cpp ackermannization/lackr_model_constructor.h.node model/model_evaluator.h.node ast/ast_smt2_pp.h.node ackermannization/ackr_info.h.node ast/for_each_expr.h.node ast/rewriter/bv_rewriter.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/ackermannization/lackr_model_constructor.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/lackr_model_constructor$(OBJ_EXT) ../src/ackermannization/lackr_model_constructor.cpp
ackermannization/ackermannization_params.hpp.node: ../src/ackermannization/ackermannization_params.hpp util/params.h.node util/gparams.h.node
	@echo done > ackermannization/ackermannization_params.hpp.node
ackermannization/lackr$(OBJ_EXT): ../src/ackermannization/lackr.cpp ackermannization/lackr.h.node ackermannization/ackermannization_params.hpp.node tactic/tactic.h.node ackermannization/lackr_model_constructor.h.node ackermannization/ackr_info.h.node ast/for_each_expr.h.node ast/ast_util.h.node model/model_smt2_pp.h.node
	@echo src/ackermannization/lackr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/lackr$(OBJ_EXT) ../src/ackermannization/lackr.cpp
ackermannization/ackr_bound_probe.h.node: ../src/ackermannization/ackr_bound_probe.h tactic/probe.h.node
	@echo done > ackermannization/ackr_bound_probe.h.node
ackermannization/ackr_bound_probe$(OBJ_EXT): ../src/ackermannization/ackr_bound_probe.cpp ast/array_decl_plugin.h.node ast/ast_smt2_pp.h.node ackermannization/ackr_helper.h.node ackermannization/ackr_bound_probe.h.node
	@echo src/ackermannization/ackr_bound_probe.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/ackr_bound_probe$(OBJ_EXT) ../src/ackermannization/ackr_bound_probe.cpp
ackermannization/lackr_model_converter_lazy.h.node: ../src/ackermannization/lackr_model_converter_lazy.h tactic/model_converter.h.node ackermannization/ackr_info.h.node
	@echo done > ackermannization/lackr_model_converter_lazy.h.node
ackermannization/lackr_model_converter_lazy$(OBJ_EXT): ../src/ackermannization/lackr_model_converter_lazy.cpp ackermannization/lackr_model_converter_lazy.h.node model/model_evaluator.h.node ast/ast_smt2_pp.h.node ackermannization/ackr_info.h.node ackermannization/lackr_model_constructor.h.node
	@echo src/ackermannization/lackr_model_converter_lazy.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/lackr_model_converter_lazy$(OBJ_EXT) ../src/ackermannization/lackr_model_converter_lazy.cpp
ackermannization/ackr_helper$(OBJ_EXT): ../src/ackermannization/ackr_helper.cpp ackermannization/ackr_helper.h.node
	@echo src/ackermannization/ackr_helper.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/ackr_helper$(OBJ_EXT) ../src/ackermannization/ackr_helper.cpp
ackermannization/ackermannize_bv_model_converter$(OBJ_EXT): ../src/ackermannization/ackermannize_bv_model_converter.cpp ackermannization/ackr_model_converter.h.node ackermannization/ackermannize_bv_model_converter.h.node
	@echo src/ackermannization/ackermannize_bv_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_28) $(CXX_OUT_FLAG)ackermannization/ackermannize_bv_model_converter$(OBJ_EXT) ../src/ackermannization/ackermannize_bv_model_converter.cpp
ackermannization/ackermannization$(LIB_EXT): ackermannization/ackermannize_bv_tactic$(OBJ_EXT) ackermannization/ackr_model_converter$(OBJ_EXT) ackermannization/lackr_model_constructor$(OBJ_EXT) ackermannization/lackr$(OBJ_EXT) ackermannization/ackr_bound_probe$(OBJ_EXT) ackermannization/lackr_model_converter_lazy$(OBJ_EXT) ackermannization/ackr_helper$(OBJ_EXT) ackermannization/ackermannize_bv_model_converter$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ackermannization/ackermannization$(LIB_EXT) ackermannization/ackermannize_bv_tactic$(OBJ_EXT) ackermannization/ackr_model_converter$(OBJ_EXT) ackermannization/lackr_model_constructor$(OBJ_EXT) ackermannization/lackr$(OBJ_EXT) ackermannization/ackr_bound_probe$(OBJ_EXT) ackermannization/lackr_model_converter_lazy$(OBJ_EXT) ackermannization/ackr_helper$(OBJ_EXT) ackermannization/ackermannize_bv_model_converter$(OBJ_EXT)
ackermannization: ackermannization/ackermannization$(LIB_EXT)

includes_29 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/model -I../src/ast/macros -I../src
ast/fpa/fpa2bv_converter.h.node: ../src/ast/fpa/fpa2bv_converter.h ast/ast.h.node util/obj_hashtable.h.node util/ref_util.h.node ast/fpa_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/dl_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node
	@echo done > ast/fpa/fpa2bv_converter.h.node
ast/fpa/fpa2bv_converter$(OBJ_EXT): ../src/ast/fpa/fpa2bv_converter.cpp ast/ast_smt2_pp.h.node ast/well_sorted.h.node ast/rewriter/th_rewriter.h.node ast/used_vars.h.node ast/rewriter/var_subst.h.node ast/fpa/fpa2bv_converter.h.node ast/rewriter/fpa_rewriter.h.node
	@echo src/ast/fpa/fpa2bv_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_29) $(CXX_OUT_FLAG)ast/fpa/fpa2bv_converter$(OBJ_EXT) ../src/ast/fpa/fpa2bv_converter.cpp
ast/fpa/bv2fpa_converter.h.node: ../src/ast/fpa/bv2fpa_converter.h util/trail.h.node ast/fpa_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/rewriter/th_rewriter.h.node model/model_core.h.node ast/fpa/fpa2bv_converter.h.node
	@echo done > ast/fpa/bv2fpa_converter.h.node
ast/fpa/bv2fpa_converter$(OBJ_EXT): ../src/ast/fpa/bv2fpa_converter.cpp ast/ast_smt2_pp.h.node ast/well_sorted.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/fpa_rewriter.h.node ast/fpa/bv2fpa_converter.h.node
	@echo src/ast/fpa/bv2fpa_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_29) $(CXX_OUT_FLAG)ast/fpa/bv2fpa_converter$(OBJ_EXT) ../src/ast/fpa/bv2fpa_converter.cpp
ast/fpa/fpa2bv_rewriter.h.node: ../src/ast/fpa/fpa2bv_rewriter.h ast/rewriter/rewriter.h.node ast/bv_decl_plugin.h.node ast/fpa/fpa2bv_converter.h.node
	@echo done > ast/fpa/fpa2bv_rewriter.h.node
params/fpa2bv_rewriter_params.hpp.node: ../src/params/fpa2bv_rewriter_params.hpp util/params.h.node util/gparams.h.node
	@echo done > params/fpa2bv_rewriter_params.hpp.node
ast/fpa/fpa2bv_rewriter$(OBJ_EXT): ../src/ast/fpa/fpa2bv_rewriter.cpp ast/rewriter/rewriter_def.h.node ast/fpa/fpa2bv_rewriter.h.node params/fpa2bv_rewriter_params.hpp.node
	@echo src/ast/fpa/fpa2bv_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_29) $(CXX_OUT_FLAG)ast/fpa/fpa2bv_rewriter$(OBJ_EXT) ../src/ast/fpa/fpa2bv_rewriter.cpp
ast/fpa/fpa$(LIB_EXT): ast/fpa/fpa2bv_converter$(OBJ_EXT) ast/fpa/bv2fpa_converter$(OBJ_EXT) ast/fpa/fpa2bv_rewriter$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/fpa/fpa$(LIB_EXT) ast/fpa/fpa2bv_converter$(OBJ_EXT) ast/fpa/bv2fpa_converter$(OBJ_EXT) ast/fpa/fpa2bv_rewriter$(OBJ_EXT)
fpa: ast/fpa/fpa$(LIB_EXT)

includes_30 = -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src
params/bit_blaster_params.h.node: ../src/params/bit_blaster_params.h
	@echo done > params/bit_blaster_params.h.node
ast/rewriter/bit_blaster/bit_blaster_tpl.h.node: ../src/ast/rewriter/bit_blaster/bit_blaster_tpl.h util/rational.h.node
	@echo done > ast/rewriter/bit_blaster/bit_blaster_tpl.h.node
ast/rewriter/bit_blaster/bit_blaster.h.node: ../src/ast/rewriter/bit_blaster/bit_blaster.h util/rational.h.node ast/bv_decl_plugin.h.node params/bit_blaster_params.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/bit_blaster/bit_blaster_tpl.h.node
	@echo done > ast/rewriter/bit_blaster/bit_blaster.h.node
ast/rewriter/bit_blaster/bit_blaster_tpl_def.h.node: ../src/ast/rewriter/bit_blaster/bit_blaster_tpl_def.h util/rational.h.node util/common_msgs.h.node ast/rewriter/bit_blaster/bit_blaster_tpl.h.node ast/ast_pp.h.node ast/rewriter/rewriter_types.h.node
	@echo done > ast/rewriter/bit_blaster/bit_blaster_tpl_def.h.node
ast/rewriter/bit_blaster/bit_blaster$(OBJ_EXT): ../src/ast/rewriter/bit_blaster/bit_blaster.cpp ast/rewriter/bit_blaster/bit_blaster.h.node ast/rewriter/bit_blaster/bit_blaster_tpl_def.h.node ast/ast_pp.h.node ast/bv_decl_plugin.h.node
	@echo src/ast/rewriter/bit_blaster/bit_blaster.cpp
	@$(CXX) $(CXXFLAGS) $(includes_30) $(CXX_OUT_FLAG)ast/rewriter/bit_blaster/bit_blaster$(OBJ_EXT) ../src/ast/rewriter/bit_blaster/bit_blaster.cpp
ast/rewriter/bit_blaster/bit_blaster_rewriter.h.node: ../src/ast/rewriter/bit_blaster/bit_blaster_rewriter.h ast/ast.h.node util/obj_hashtable.h.node util/params.h.node
	@echo done > ast/rewriter/bit_blaster/bit_blaster_rewriter.h.node
ast/rewriter/bit_blaster/bit_blaster_rewriter$(OBJ_EXT): ../src/ast/rewriter/bit_blaster/bit_blaster_rewriter.cpp ast/rewriter/bit_blaster/bit_blaster_rewriter.h.node ast/bv_decl_plugin.h.node ast/rewriter/bit_blaster/bit_blaster_tpl_def.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/bool_rewriter.h.node util/ref_util.h.node ast/ast_smt2_pp.h.node
	@echo src/ast/rewriter/bit_blaster/bit_blaster_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_30) $(CXX_OUT_FLAG)ast/rewriter/bit_blaster/bit_blaster_rewriter$(OBJ_EXT) ../src/ast/rewriter/bit_blaster/bit_blaster_rewriter.cpp
ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT): ast/rewriter/bit_blaster/bit_blaster$(OBJ_EXT) ast/rewriter/bit_blaster/bit_blaster_rewriter$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(OBJ_EXT) ast/rewriter/bit_blaster/bit_blaster_rewriter$(OBJ_EXT)
bit_blaster: ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT)

includes_31 = -I../src/ast/normal_forms -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/model -I../src/ast/macros -I../src/tactic -I../src/ast/proofs -I../src/parsers/util -I../src
util/nat_set.h.node: ../src/util/nat_set.h util/vector.h.node
	@echo done > util/nat_set.h.node
ast/pattern/expr_pattern_match.h.node: ../src/ast/pattern/expr_pattern_match.h ast/ast.h.node util/map.h.node
	@echo done > ast/pattern/expr_pattern_match.h.node
ast/pattern/pattern_inference.h.node: ../src/ast/pattern/pattern_inference.h ast/ast.h.node ast/rewriter/rewriter.h.node params/pattern_inference_params.h.node util/vector.h.node util/uint_set.h.node util/nat_set.h.node util/obj_hashtable.h.node util/obj_pair_hashtable.h.node util/map.h.node ast/pattern/expr_pattern_match.h.node
	@echo done > ast/pattern/pattern_inference.h.node
ast/pattern/database.h.node: ../src/ast/pattern/database.h
	@echo done > ast/pattern/database.h.node
ast/pattern/pattern_inference$(OBJ_EXT): ../src/ast/pattern/pattern_inference.cpp util/warning.h.node ast/pattern/pattern_inference.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/arith_decl_plugin.h.node ast/normal_forms/pull_quant.h.node ast/well_sorted.h.node ast/for_each_expr.h.node ast/rewriter/rewriter_def.h.node ast/pattern/database.h.node
	@echo src/ast/pattern/pattern_inference.cpp
	@$(CXX) $(CXXFLAGS) $(includes_31) $(CXX_OUT_FLAG)ast/pattern/pattern_inference$(OBJ_EXT) ../src/ast/pattern/pattern_inference.cpp
ast/pattern/expr_pattern_match$(OBJ_EXT): ../src/ast/pattern/expr_pattern_match.cpp ast/ast.h.node ast/pattern/expr_pattern_match.h.node ast/for_each_ast.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node cmd_context/cmd_context.h.node parsers/smt2/smt2parser.h.node
	@echo src/ast/pattern/expr_pattern_match.cpp
	@$(CXX) $(CXXFLAGS) $(includes_31) $(CXX_OUT_FLAG)ast/pattern/expr_pattern_match$(OBJ_EXT) ../src/ast/pattern/expr_pattern_match.cpp
ast/pattern/pattern$(LIB_EXT): ast/pattern/pattern_inference$(OBJ_EXT) ast/pattern/expr_pattern_match$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)ast/pattern/pattern$(LIB_EXT) ast/pattern/pattern_inference$(OBJ_EXT) ast/pattern/expr_pattern_match$(OBJ_EXT)
pattern: ast/pattern/pattern$(LIB_EXT)

includes_32 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src/params -I../src
smt/params/theory_pb_params.h.node: ../src/smt/params/theory_pb_params.h util/params.h.node
	@echo done > smt/params/theory_pb_params.h.node
smt/params/smt_params_helper.hpp.node: ../src/smt/params/smt_params_helper.hpp util/params.h.node util/gparams.h.node
	@echo done > smt/params/smt_params_helper.hpp.node
smt/params/theory_pb_params$(OBJ_EXT): ../src/smt/params/theory_pb_params.cpp smt/params/theory_pb_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/theory_pb_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/theory_pb_params$(OBJ_EXT) ../src/smt/params/theory_pb_params.cpp
smt/params/theory_arith_params.h.node: ../src/smt/params/theory_arith_params.h util/params.h.node
	@echo done > smt/params/theory_arith_params.h.node
smt/params/theory_arith_params$(OBJ_EXT): ../src/smt/params/theory_arith_params.cpp smt/params/theory_arith_params.h.node smt/params/smt_params_helper.hpp.node params/arith_rewriter_params.hpp.node
	@echo src/smt/params/theory_arith_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/theory_arith_params$(OBJ_EXT) ../src/smt/params/theory_arith_params.cpp
smt/params/preprocessor_params.h.node: ../src/smt/params/preprocessor_params.h params/pattern_inference_params.h.node params/bit_blaster_params.h.node
	@echo done > smt/params/preprocessor_params.h.node
smt/params/preprocessor_params$(OBJ_EXT): ../src/smt/params/preprocessor_params.cpp smt/params/preprocessor_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/preprocessor_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/preprocessor_params$(OBJ_EXT) ../src/smt/params/preprocessor_params.cpp
smt/params/dyn_ack_params.h.node: ../src/smt/params/dyn_ack_params.h util/params.h.node
	@echo done > smt/params/dyn_ack_params.h.node
smt/params/qi_params.h.node: ../src/smt/params/qi_params.h util/util.h.node util/params.h.node
	@echo done > smt/params/qi_params.h.node
smt/params/theory_array_params.h.node: ../src/smt/params/theory_array_params.h util/params.h.node
	@echo done > smt/params/theory_array_params.h.node
smt/params/theory_bv_params.h.node: ../src/smt/params/theory_bv_params.h util/params.h.node
	@echo done > smt/params/theory_bv_params.h.node
smt/params/theory_str_params.h.node: ../src/smt/params/theory_str_params.h util/params.h.node
	@echo done > smt/params/theory_str_params.h.node
smt/params/theory_seq_params.h.node: ../src/smt/params/theory_seq_params.h util/params.h.node
	@echo done > smt/params/theory_seq_params.h.node
smt/params/theory_datatype_params.h.node: ../src/smt/params/theory_datatype_params.h smt/params/smt_params_helper.hpp.node
	@echo done > smt/params/theory_datatype_params.h.node
smt/params/smt_params.h.node: ../src/smt/params/smt_params.h smt/params/dyn_ack_params.h.node smt/params/qi_params.h.node smt/params/theory_arith_params.h.node smt/params/theory_array_params.h.node smt/params/theory_bv_params.h.node smt/params/theory_str_params.h.node smt/params/theory_seq_params.h.node smt/params/theory_pb_params.h.node smt/params/theory_datatype_params.h.node smt/params/preprocessor_params.h.node params/context_params.h.node
	@echo done > smt/params/smt_params.h.node
smt/params/smt_params$(OBJ_EXT): ../src/smt/params/smt_params.cpp smt/params/smt_params.h.node smt/params/smt_params_helper.hpp.node util/gparams.h.node
	@echo src/smt/params/smt_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/smt_params$(OBJ_EXT) ../src/smt/params/smt_params.cpp
smt/params/theory_str_params$(OBJ_EXT): ../src/smt/params/theory_str_params.cpp smt/params/theory_str_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/theory_str_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/theory_str_params$(OBJ_EXT) ../src/smt/params/theory_str_params.cpp
smt/params/theory_seq_params$(OBJ_EXT): ../src/smt/params/theory_seq_params.cpp smt/params/theory_seq_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/theory_seq_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/theory_seq_params$(OBJ_EXT) ../src/smt/params/theory_seq_params.cpp
smt/params/dyn_ack_params$(OBJ_EXT): ../src/smt/params/dyn_ack_params.cpp smt/params/dyn_ack_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/dyn_ack_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/dyn_ack_params$(OBJ_EXT) ../src/smt/params/dyn_ack_params.cpp
smt/params/theory_bv_params$(OBJ_EXT): ../src/smt/params/theory_bv_params.cpp smt/params/theory_bv_params.h.node smt/params/smt_params_helper.hpp.node params/bv_rewriter_params.hpp.node
	@echo src/smt/params/theory_bv_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/theory_bv_params$(OBJ_EXT) ../src/smt/params/theory_bv_params.cpp
smt/params/theory_array_params$(OBJ_EXT): ../src/smt/params/theory_array_params.cpp smt/params/theory_array_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/theory_array_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/theory_array_params$(OBJ_EXT) ../src/smt/params/theory_array_params.cpp
smt/params/qi_params$(OBJ_EXT): ../src/smt/params/qi_params.cpp smt/params/qi_params.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/params/qi_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_32) $(CXX_OUT_FLAG)smt/params/qi_params$(OBJ_EXT) ../src/smt/params/qi_params.cpp
smt/params/smt_params$(LIB_EXT): smt/params/theory_pb_params$(OBJ_EXT) smt/params/theory_arith_params$(OBJ_EXT) smt/params/preprocessor_params$(OBJ_EXT) smt/params/smt_params$(OBJ_EXT) smt/params/theory_str_params$(OBJ_EXT) smt/params/theory_seq_params$(OBJ_EXT) smt/params/dyn_ack_params$(OBJ_EXT) smt/params/theory_bv_params$(OBJ_EXT) smt/params/theory_array_params$(OBJ_EXT) smt/params/qi_params$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)smt/params/smt_params$(LIB_EXT) smt/params/theory_pb_params$(OBJ_EXT) smt/params/theory_arith_params$(OBJ_EXT) smt/params/preprocessor_params$(OBJ_EXT) smt/params/smt_params$(OBJ_EXT) smt/params/theory_str_params$(OBJ_EXT) smt/params/theory_seq_params$(OBJ_EXT) smt/params/dyn_ack_params$(OBJ_EXT) smt/params/theory_bv_params$(OBJ_EXT) smt/params/theory_array_params$(OBJ_EXT) smt/params/qi_params$(OBJ_EXT)
smt_params: smt/params/smt_params$(LIB_EXT)

includes_33 = -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src
tactic/core/reduce_invertible_tactic.h.node: ../src/tactic/core/reduce_invertible_tactic.h util/params.h.node
	@echo done > tactic/core/reduce_invertible_tactic.h.node
tactic/core/collect_occs.h.node: ../src/tactic/core/collect_occs.h
	@echo done > tactic/core/collect_occs.h.node
tactic/core/reduce_invertible_tactic$(OBJ_EXT): ../src/tactic/core/reduce_invertible_tactic.cpp ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/th_rewriter.h.node tactic/tactic.h.node tactic/core/reduce_invertible_tactic.h.node tactic/core/collect_occs.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/core/reduce_invertible_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/reduce_invertible_tactic$(OBJ_EXT) ../src/tactic/core/reduce_invertible_tactic.cpp
tactic/core/special_relations_tactic.h.node: ../src/tactic/core/special_relations_tactic.h tactic/tactic.h.node tactic/tactical.h.node ast/special_relations_decl_plugin.h.node ast/pattern/expr_pattern_match.h.node
	@echo done > tactic/core/special_relations_tactic.h.node
tactic/core/special_relations_tactic$(OBJ_EXT): ../src/tactic/core/special_relations_tactic.cpp tactic/core/special_relations_tactic.h.node ast/rewriter/func_decl_replace.h.node ast/ast_util.h.node ast/ast_pp.h.node
	@echo src/tactic/core/special_relations_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/special_relations_tactic$(OBJ_EXT) ../src/tactic/core/special_relations_tactic.cpp
tactic/core/simplify_tactic.h.node: ../src/tactic/core/simplify_tactic.h tactic/tactic.h.node tactic/tactical.h.node
	@echo done > tactic/core/simplify_tactic.h.node
tactic/core/tseitin_cnf_tactic$(OBJ_EXT): ../src/tactic/core/tseitin_cnf_tactic.cpp ast/ast_pp.h.node tactic/tactical.h.node tactic/goal_shared_occs.h.node tactic/generic_model_converter.h.node ast/rewriter/bool_rewriter.h.node tactic/core/simplify_tactic.h.node
	@echo src/tactic/core/tseitin_cnf_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/tseitin_cnf_tactic$(OBJ_EXT) ../src/tactic/core/tseitin_cnf_tactic.cpp
tactic/core/dom_simplify_tactic.h.node: ../src/tactic/core/dom_simplify_tactic.h ast/ast.h.node ast/expr_substitution.h.node tactic/tactic.h.node tactic/tactical.h.node util/obj_pair_hashtable.h.node
	@echo done > tactic/core/dom_simplify_tactic.h.node
tactic/core/dom_simplify_tactic$(OBJ_EXT): ../src/tactic/core/dom_simplify_tactic.cpp ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node tactic/core/dom_simplify_tactic.h.node
	@echo src/tactic/core/dom_simplify_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/dom_simplify_tactic$(OBJ_EXT) ../src/tactic/core/dom_simplify_tactic.cpp
tactic/core/distribute_forall_tactic$(OBJ_EXT): ../src/tactic/core/distribute_forall_tactic.cpp tactic/tactical.h.node ast/ast_util.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node
	@echo src/tactic/core/distribute_forall_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/distribute_forall_tactic$(OBJ_EXT) ../src/tactic/core/distribute_forall_tactic.cpp
tactic/core/cofactor_elim_term_ite.h.node: ../src/tactic/core/cofactor_elim_term_ite.h ast/ast.h.node util/params.h.node
	@echo done > tactic/core/cofactor_elim_term_ite.h.node
tactic/core/cofactor_term_ite_tactic$(OBJ_EXT): ../src/tactic/core/cofactor_term_ite_tactic.cpp tactic/tactical.h.node tactic/core/cofactor_elim_term_ite.h.node
	@echo src/tactic/core/cofactor_term_ite_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/cofactor_term_ite_tactic$(OBJ_EXT) ../src/tactic/core/cofactor_term_ite_tactic.cpp
tactic/core/collect_occs$(OBJ_EXT): ../src/tactic/core/collect_occs.cpp ast/ast.h.node tactic/goal.h.node util/hashtable.h.node tactic/core/collect_occs.h.node
	@echo src/tactic/core/collect_occs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/collect_occs$(OBJ_EXT) ../src/tactic/core/collect_occs.cpp
tactic/core/pb_preprocess_tactic.h.node: ../src/tactic/core/pb_preprocess_tactic.h util/params.h.node
	@echo done > tactic/core/pb_preprocess_tactic.h.node
tactic/core/pb_preprocess_tactic$(OBJ_EXT): ../src/tactic/core/pb_preprocess_tactic.cpp tactic/core/pb_preprocess_tactic.h.node tactic/tactical.h.node tactic/generic_model_converter.h.node ast/for_each_expr.h.node ast/pb_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/expr_substitution.h.node ast/ast_pp.h.node
	@echo src/tactic/core/pb_preprocess_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/pb_preprocess_tactic$(OBJ_EXT) ../src/tactic/core/pb_preprocess_tactic.cpp
tactic/core/symmetry_reduce_tactic$(OBJ_EXT): ../src/tactic/core/symmetry_reduce_tactic.cpp tactic/tactical.h.node ast/for_each_expr.h.node util/map.h.node ast/rewriter/expr_replacer.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node
	@echo src/tactic/core/symmetry_reduce_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/symmetry_reduce_tactic$(OBJ_EXT) ../src/tactic/core/symmetry_reduce_tactic.cpp
tactic/core/split_clause_tactic.h.node: ../src/tactic/core/split_clause_tactic.h util/params.h.node
	@echo done > tactic/core/split_clause_tactic.h.node
tactic/core/split_clause_tactic$(OBJ_EXT): ../src/tactic/core/split_clause_tactic.cpp tactic/tactical.h.node tactic/core/split_clause_tactic.h.node
	@echo src/tactic/core/split_clause_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/split_clause_tactic$(OBJ_EXT) ../src/tactic/core/split_clause_tactic.cpp
tactic/core/der_tactic$(OBJ_EXT): ../src/tactic/core/der_tactic.cpp ast/rewriter/der.h.node tactic/tactical.h.node
	@echo src/tactic/core/der_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/der_tactic$(OBJ_EXT) ../src/tactic/core/der_tactic.cpp
tactic/core/nnf_tactic$(OBJ_EXT): ../src/tactic/core/nnf_tactic.cpp ast/normal_forms/nnf.h.node tactic/tactical.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/core/nnf_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/nnf_tactic$(OBJ_EXT) ../src/tactic/core/nnf_tactic.cpp
tactic/core/ctx_simplify_tactic.h.node: ../src/tactic/core/ctx_simplify_tactic.h tactic/tactical.h.node tactic/goal_num_occurs.h.node
	@echo done > tactic/core/ctx_simplify_tactic.h.node
tactic/core/ctx_simplify_tactic$(OBJ_EXT): ../src/tactic/core/ctx_simplify_tactic.cpp tactic/core/ctx_simplify_tactic.h.node ast/rewriter/mk_simplified_app.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node
	@echo src/tactic/core/ctx_simplify_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/ctx_simplify_tactic$(OBJ_EXT) ../src/tactic/core/ctx_simplify_tactic.cpp
util/stats.h.node: ../src/util/stats.h
	@echo done > util/stats.h.node
tactic/core/collect_statistics_tactic.h.node: ../src/tactic/core/collect_statistics_tactic.h util/params.h.node
	@echo done > tactic/core/collect_statistics_tactic.h.node
tactic/core/collect_statistics_tactic$(OBJ_EXT): ../src/tactic/core/collect_statistics_tactic.cpp ast/ast.h.node util/params.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/fpa_decl_plugin.h.node tactic/tactical.h.node util/stats.h.node tactic/core/collect_statistics_tactic.h.node
	@echo src/tactic/core/collect_statistics_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/collect_statistics_tactic$(OBJ_EXT) ../src/tactic/core/collect_statistics_tactic.cpp
tactic/core/elim_term_ite_tactic$(OBJ_EXT): ../src/tactic/core/elim_term_ite_tactic.cpp tactic/tactical.h.node ast/normal_forms/defined_names.h.node ast/rewriter/rewriter_def.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/core/elim_term_ite_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/elim_term_ite_tactic$(OBJ_EXT) ../src/tactic/core/elim_term_ite_tactic.cpp
tactic/core/propagate_values_tactic.h.node: ../src/tactic/core/propagate_values_tactic.h util/params.h.node
	@echo done > tactic/core/propagate_values_tactic.h.node
tactic/tactic_params.hpp.node: ../src/tactic/tactic_params.hpp util/params.h.node util/gparams.h.node
	@echo done > tactic/tactic_params.hpp.node
tactic/core/propagate_values_tactic$(OBJ_EXT): ../src/tactic/core/propagate_values_tactic.cpp tactic/tactical.h.node tactic/core/propagate_values_tactic.h.node ast/rewriter/th_rewriter.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/expr_substitution.h.node tactic/goal_shared_occs.h.node tactic/tactic_params.hpp.node
	@echo src/tactic/core/propagate_values_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/propagate_values_tactic$(OBJ_EXT) ../src/tactic/core/propagate_values_tactic.cpp
tactic/core/cofactor_elim_term_ite$(OBJ_EXT): ../src/tactic/core/cofactor_elim_term_ite.cpp tactic/core/cofactor_elim_term_ite.h.node ast/rewriter/mk_simplified_app.h.node ast/rewriter/rewriter_def.h.node ast/for_each_expr.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node tactic/tactic.h.node
	@echo src/tactic/core/cofactor_elim_term_ite.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/cofactor_elim_term_ite$(OBJ_EXT) ../src/tactic/core/cofactor_elim_term_ite.cpp
tactic/core/elim_uncnstr_tactic$(OBJ_EXT): ../src/tactic/core/elim_uncnstr_tactic.cpp tactic/tactical.h.node tactic/generic_model_converter.h.node ast/rewriter/rewriter_def.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node tactic/core/collect_occs.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node
	@echo src/tactic/core/elim_uncnstr_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/elim_uncnstr_tactic$(OBJ_EXT) ../src/tactic/core/elim_uncnstr_tactic.cpp
tactic/core/solve_eqs_tactic$(OBJ_EXT): ../src/tactic/core/solve_eqs_tactic.cpp ast/rewriter/expr_replacer.h.node ast/occurs.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/pb_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/hoist_rewriter.h.node tactic/goal_shared_occs.h.node tactic/tactical.h.node tactic/generic_model_converter.h.node tactic/tactic_params.hpp.node
	@echo src/tactic/core/solve_eqs_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/solve_eqs_tactic$(OBJ_EXT) ../src/tactic/core/solve_eqs_tactic.cpp
tactic/core/reduce_args_tactic$(OBJ_EXT): ../src/tactic/core/reduce_args_tactic.cpp tactic/tactical.h.node ast/ast_smt2_pp.h.node ast/array_decl_plugin.h.node ast/has_free_vars.h.node util/map.h.node ast/rewriter/rewriter_def.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/core/reduce_args_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/reduce_args_tactic$(OBJ_EXT) ../src/tactic/core/reduce_args_tactic.cpp
tactic/core/injectivity_tactic.h.node: ../src/tactic/core/injectivity_tactic.h util/params.h.node
	@echo done > tactic/core/injectivity_tactic.h.node
tactic/core/injectivity_tactic$(OBJ_EXT): ../src/tactic/core/injectivity_tactic.cpp tactic/tactical.h.node ast/rewriter/rewriter_def.h.node tactic/core/injectivity_tactic.h.node util/dec_ref_util.h.node
	@echo src/tactic/core/injectivity_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/injectivity_tactic$(OBJ_EXT) ../src/tactic/core/injectivity_tactic.cpp
tactic/core/occf_tactic.h.node: ../src/tactic/core/occf_tactic.h util/params.h.node
	@echo done > tactic/core/occf_tactic.h.node
tactic/core/occf_tactic$(OBJ_EXT): ../src/tactic/core/occf_tactic.cpp tactic/tactical.h.node tactic/core/occf_tactic.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/core/occf_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/occf_tactic$(OBJ_EXT) ../src/tactic/core/occf_tactic.cpp
tactic/core/simplify_tactic$(OBJ_EXT): ../src/tactic/core/simplify_tactic.cpp tactic/core/simplify_tactic.h.node ast/rewriter/th_rewriter.h.node ast/ast_pp.h.node
	@echo src/tactic/core/simplify_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/simplify_tactic$(OBJ_EXT) ../src/tactic/core/simplify_tactic.cpp
tactic/core/blast_term_ite_tactic$(OBJ_EXT): ../src/tactic/core/blast_term_ite_tactic.cpp ast/normal_forms/defined_names.h.node ast/rewriter/rewriter_def.h.node ast/scoped_proof.h.node tactic/tactical.h.node tactic/tactic_params.hpp.node
	@echo src/tactic/core/blast_term_ite_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_33) $(CXX_OUT_FLAG)tactic/core/blast_term_ite_tactic$(OBJ_EXT) ../src/tactic/core/blast_term_ite_tactic.cpp
tactic/core/core_tactics$(LIB_EXT): tactic/core/reduce_invertible_tactic$(OBJ_EXT) tactic/core/special_relations_tactic$(OBJ_EXT) tactic/core/tseitin_cnf_tactic$(OBJ_EXT) tactic/core/dom_simplify_tactic$(OBJ_EXT) tactic/core/distribute_forall_tactic$(OBJ_EXT) tactic/core/cofactor_term_ite_tactic$(OBJ_EXT) tactic/core/collect_occs$(OBJ_EXT) tactic/core/pb_preprocess_tactic$(OBJ_EXT) tactic/core/symmetry_reduce_tactic$(OBJ_EXT) tactic/core/split_clause_tactic$(OBJ_EXT) tactic/core/der_tactic$(OBJ_EXT) tactic/core/nnf_tactic$(OBJ_EXT) tactic/core/ctx_simplify_tactic$(OBJ_EXT) tactic/core/collect_statistics_tactic$(OBJ_EXT) tactic/core/elim_term_ite_tactic$(OBJ_EXT) tactic/core/propagate_values_tactic$(OBJ_EXT) tactic/core/cofactor_elim_term_ite$(OBJ_EXT) tactic/core/elim_uncnstr_tactic$(OBJ_EXT) tactic/core/solve_eqs_tactic$(OBJ_EXT) tactic/core/reduce_args_tactic$(OBJ_EXT) tactic/core/injectivity_tactic$(OBJ_EXT) tactic/core/occf_tactic$(OBJ_EXT) tactic/core/simplify_tactic$(OBJ_EXT) tactic/core/blast_term_ite_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/core/core_tactics$(LIB_EXT) tactic/core/reduce_invertible_tactic$(OBJ_EXT) tactic/core/special_relations_tactic$(OBJ_EXT) tactic/core/tseitin_cnf_tactic$(OBJ_EXT) tactic/core/dom_simplify_tactic$(OBJ_EXT) tactic/core/distribute_forall_tactic$(OBJ_EXT) tactic/core/cofactor_term_ite_tactic$(OBJ_EXT) tactic/core/collect_occs$(OBJ_EXT) tactic/core/pb_preprocess_tactic$(OBJ_EXT) tactic/core/symmetry_reduce_tactic$(OBJ_EXT) tactic/core/split_clause_tactic$(OBJ_EXT) tactic/core/der_tactic$(OBJ_EXT) tactic/core/nnf_tactic$(OBJ_EXT) tactic/core/ctx_simplify_tactic$(OBJ_EXT) tactic/core/collect_statistics_tactic$(OBJ_EXT) tactic/core/elim_term_ite_tactic$(OBJ_EXT) tactic/core/propagate_values_tactic$(OBJ_EXT) tactic/core/cofactor_elim_term_ite$(OBJ_EXT) tactic/core/elim_uncnstr_tactic$(OBJ_EXT) tactic/core/solve_eqs_tactic$(OBJ_EXT) tactic/core/reduce_args_tactic$(OBJ_EXT) tactic/core/injectivity_tactic$(OBJ_EXT) tactic/core/occf_tactic$(OBJ_EXT) tactic/core/simplify_tactic$(OBJ_EXT) tactic/core/blast_term_ite_tactic$(OBJ_EXT)
core_tactics: tactic/core/core_tactics$(LIB_EXT)

includes_34 = -I../src/tactic/core -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src
tactic/arith/card2bv_tactic.h.node: ../src/tactic/arith/card2bv_tactic.h util/params.h.node ast/pb_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/rewriter.h.node util/sorting_network.h.node
	@echo done > tactic/arith/card2bv_tactic.h.node
tactic/arith/card2bv_tactic$(OBJ_EXT): ../src/tactic/arith/card2bv_tactic.cpp tactic/tactical.h.node ast/ast_smt2_pp.h.node tactic/arith/card2bv_tactic.h.node ast/rewriter/pb2bv_rewriter.h.node ast/ast_util.h.node ast/ast_pp.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/arith/card2bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/card2bv_tactic$(OBJ_EXT) ../src/tactic/arith/card2bv_tactic.cpp
tactic/arith/arith_bounds_tactic.h.node: ../src/tactic/arith/arith_bounds_tactic.h tactic/tactic.h.node
	@echo done > tactic/arith/arith_bounds_tactic.h.node
tactic/arith/arith_bounds_tactic$(OBJ_EXT): ../src/tactic/arith/arith_bounds_tactic.cpp tactic/arith/arith_bounds_tactic.h.node ast/arith_decl_plugin.h.node
	@echo src/tactic/arith/arith_bounds_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/arith_bounds_tactic$(OBJ_EXT) ../src/tactic/arith/arith_bounds_tactic.cpp
tactic/arith/fm_tactic.h.node: ../src/tactic/arith/fm_tactic.h util/params.h.node
	@echo done > tactic/arith/fm_tactic.h.node
tactic/arith/fm_tactic$(OBJ_EXT): ../src/tactic/arith/fm_tactic.cpp tactic/arith/fm_tactic.h.node tactic/tactical.h.node ast/arith_decl_plugin.h.node ast/for_each_expr.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node util/id_gen.h.node model/model_evaluator.h.node model/model_v2_pp.h.node tactic/core/simplify_tactic.h.node
	@echo src/tactic/arith/fm_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/fm_tactic$(OBJ_EXT) ../src/tactic/arith/fm_tactic.cpp
tactic/arith/degree_shift_tactic$(OBJ_EXT): ../src/tactic/arith/degree_shift_tactic.cpp tactic/tactical.h.node tactic/generic_model_converter.h.node ast/arith_decl_plugin.h.node tactic/core/simplify_tactic.h.node ast/ast_smt2_pp.h.node ast/rewriter/rewriter_def.h.node
	@echo src/tactic/arith/degree_shift_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/degree_shift_tactic$(OBJ_EXT) ../src/tactic/arith/degree_shift_tactic.cpp
tactic/arith/bound_manager.h.node: ../src/tactic/arith/bound_manager.h ast/ast.h.node ast/arith_decl_plugin.h.node
	@echo done > tactic/arith/bound_manager.h.node
tactic/arith/lia2card_tactic$(OBJ_EXT): ../src/tactic/arith/lia2card_tactic.cpp ast/ast_pp.h.node ast/pb_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/expr_safe_replace.h.node ast/ast_util.h.node ast/ast_pp_util.h.node tactic/tactical.h.node tactic/arith/bound_manager.h.node tactic/generic_model_converter.h.node
	@echo src/tactic/arith/lia2card_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/lia2card_tactic$(OBJ_EXT) ../src/tactic/arith/lia2card_tactic.cpp
util/numeral_buffer.h.node: ../src/util/numeral_buffer.h util/vector.h.node
	@echo done > util/numeral_buffer.h.node
util/double_manager.h.node: ../src/util/double_manager.h util/util.h.node util/debug.h.node util/hash.h.node util/params.h.node
	@echo done > util/double_manager.h.node
tactic/arith/linear_equation.h.node: ../src/tactic/arith/linear_equation.h util/mpq.h.node util/small_object_allocator.h.node util/numeral_buffer.h.node util/double_manager.h.node
	@echo done > tactic/arith/linear_equation.h.node
tactic/arith/linear_equation$(OBJ_EXT): ../src/tactic/arith/linear_equation.cpp tactic/arith/linear_equation.h.node
	@echo src/tactic/arith/linear_equation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/linear_equation$(OBJ_EXT) ../src/tactic/arith/linear_equation.cpp
tactic/arith/bv2int_rewriter.h.node: ../src/tactic/arith/bv2int_rewriter.h ast/ast.h.node ast/rewriter/rewriter.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node util/params.h.node tactic/goal.h.node
	@echo done > tactic/arith/bv2int_rewriter.h.node
tactic/arith/bv2int_rewriter$(OBJ_EXT): ../src/tactic/arith/bv2int_rewriter.cpp tactic/arith/bv2int_rewriter.h.node tactic/tactic_exception.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node ast/ast_util.h.node
	@echo src/tactic/arith/bv2int_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/bv2int_rewriter$(OBJ_EXT) ../src/tactic/arith/bv2int_rewriter.cpp
tactic/arith/pb2bv_model_converter.h.node: ../src/tactic/arith/pb2bv_model_converter.h tactic/model_converter.h.node tactic/arith/bound_manager.h.node
	@echo done > tactic/arith/pb2bv_model_converter.h.node
tactic/arith/pb2bv_tactic.h.node: ../src/tactic/arith/pb2bv_tactic.h util/params.h.node
	@echo done > tactic/arith/pb2bv_tactic.h.node
tactic/arith/pb2bv_tactic$(OBJ_EXT): ../src/tactic/arith/pb2bv_tactic.cpp util/ref_util.h.node util/trace.h.node util/statistics.h.node ast/arith_decl_plugin.h.node ast/ast_smt2_pp.h.node ast/expr_substitution.h.node ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/pb2bv_rewriter.h.node tactic/tactical.h.node tactic/arith/bound_manager.h.node tactic/generic_model_converter.h.node tactic/arith/pb2bv_model_converter.h.node tactic/arith/pb2bv_tactic.h.node
	@echo src/tactic/arith/pb2bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/pb2bv_tactic$(OBJ_EXT) ../src/tactic/arith/pb2bv_tactic.cpp
tactic/arith/bv2real_rewriter.h.node: ../src/tactic/arith/bv2real_rewriter.h ast/ast.h.node ast/rewriter/rewriter.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node
	@echo done > tactic/arith/bv2real_rewriter.h.node
tactic/arith/bv2real_rewriter$(OBJ_EXT): ../src/tactic/arith/bv2real_rewriter.cpp tactic/arith/bv2real_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node ast/for_each_expr.h.node
	@echo src/tactic/arith/bv2real_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/bv2real_rewriter$(OBJ_EXT) ../src/tactic/arith/bv2real_rewriter.cpp
tactic/arith/diff_neq_tactic$(OBJ_EXT): ../src/tactic/arith/diff_neq_tactic.cpp tactic/tactical.h.node ast/arith_decl_plugin.h.node ast/ast_smt2_pp.h.node model/model.h.node
	@echo src/tactic/arith/diff_neq_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/diff_neq_tactic$(OBJ_EXT) ../src/tactic/arith/diff_neq_tactic.cpp
tactic/arith/bound_propagator.h.node: ../src/tactic/arith/bound_propagator.h util/mpq.h.node util/vector.h.node util/params.h.node util/statistics.h.node util/numeral_buffer.h.node tactic/arith/linear_equation.h.node
	@echo done > tactic/arith/bound_propagator.h.node
tactic/arith/bound_propagator$(OBJ_EXT): ../src/tactic/arith/bound_propagator.cpp tactic/arith/bound_propagator.h.node
	@echo src/tactic/arith/bound_propagator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/bound_propagator$(OBJ_EXT) ../src/tactic/arith/bound_propagator.cpp
tactic/core/nnf_tactic.h.node: ../src/tactic/core/nnf_tactic.h util/params.h.node
	@echo done > tactic/core/nnf_tactic.h.node
tactic/arith/purify_arith_tactic$(OBJ_EXT): ../src/tactic/arith/purify_arith_tactic.cpp tactic/tactical.h.node ast/rewriter/rewriter_def.h.node ast/arith_decl_plugin.h.node math/polynomial/algebraic_numbers.h.node tactic/core/nnf_tactic.h.node tactic/core/simplify_tactic.h.node ast/rewriter/th_rewriter.h.node tactic/generic_model_converter.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/rewriter/expr_replacer.h.node
	@echo src/tactic/arith/purify_arith_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/purify_arith_tactic$(OBJ_EXT) ../src/tactic/arith/purify_arith_tactic.cpp
tactic/arith/recover_01_tactic$(OBJ_EXT): ../src/tactic/arith/recover_01_tactic.cpp tactic/tactical.h.node ast/rewriter/th_rewriter.h.node tactic/generic_model_converter.h.node ast/arith_decl_plugin.h.node ast/expr_substitution.h.node util/dec_ref_util.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/arith/recover_01_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/recover_01_tactic$(OBJ_EXT) ../src/tactic/arith/recover_01_tactic.cpp
tactic/arith/bound_manager$(OBJ_EXT): ../src/tactic/arith/bound_manager.cpp tactic/arith/bound_manager.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node tactic/goal.h.node
	@echo src/tactic/arith/bound_manager.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/bound_manager$(OBJ_EXT) ../src/tactic/arith/bound_manager.cpp
tactic/arith/lia2pb_tactic$(OBJ_EXT): ../src/tactic/arith/lia2pb_tactic.cpp tactic/tactical.h.node tactic/arith/bound_manager.h.node ast/rewriter/th_rewriter.h.node ast/for_each_expr.h.node tactic/generic_model_converter.h.node ast/arith_decl_plugin.h.node ast/expr_substitution.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/arith/lia2pb_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/lia2pb_tactic$(OBJ_EXT) ../src/tactic/arith/lia2pb_tactic.cpp
tactic/arith/normalize_bounds_tactic$(OBJ_EXT): ../src/tactic/arith/normalize_bounds_tactic.cpp tactic/tactical.h.node tactic/arith/bound_manager.h.node ast/rewriter/th_rewriter.h.node tactic/generic_model_converter.h.node ast/arith_decl_plugin.h.node ast/expr_substitution.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/arith/normalize_bounds_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/normalize_bounds_tactic$(OBJ_EXT) ../src/tactic/arith/normalize_bounds_tactic.cpp
tactic/arith/propagate_ineqs_tactic$(OBJ_EXT): ../src/tactic/arith/propagate_ineqs_tactic.cpp tactic/tactical.h.node tactic/arith/bound_propagator.h.node ast/arith_decl_plugin.h.node tactic/core/simplify_tactic.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/arith/propagate_ineqs_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/propagate_ineqs_tactic$(OBJ_EXT) ../src/tactic/arith/propagate_ineqs_tactic.cpp
tactic/arith/nla2bv_tactic$(OBJ_EXT): ../src/tactic/arith/nla2bv_tactic.cpp tactic/tactical.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/for_each_expr.h.node ast/rewriter/expr_replacer.h.node util/optional.h.node tactic/arith/bv2int_rewriter.h.node tactic/arith/bv2real_rewriter.h.node tactic/generic_model_converter.h.node tactic/arith/bound_manager.h.node util/obj_pair_hashtable.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/arith/nla2bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/nla2bv_tactic$(OBJ_EXT) ../src/tactic/arith/nla2bv_tactic.cpp
tactic/arith/probe_arith$(OBJ_EXT): ../src/tactic/arith/probe_arith.cpp tactic/probe.h.node ast/expr2polynomial.h.node ast/for_each_expr.h.node ast/arith_decl_plugin.h.node tactic/goal_util.h.node
	@echo src/tactic/arith/probe_arith.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/probe_arith$(OBJ_EXT) ../src/tactic/arith/probe_arith.cpp
tactic/arith/fix_dl_var_tactic$(OBJ_EXT): ../src/tactic/arith/fix_dl_var_tactic.cpp tactic/tactical.h.node ast/rewriter/th_rewriter.h.node tactic/generic_model_converter.h.node ast/arith_decl_plugin.h.node ast/expr_substitution.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node
	@echo src/tactic/arith/fix_dl_var_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/fix_dl_var_tactic$(OBJ_EXT) ../src/tactic/arith/fix_dl_var_tactic.cpp
tactic/arith/factor_tactic$(OBJ_EXT): ../src/tactic/arith/factor_tactic.cpp tactic/tactical.h.node ast/expr2polynomial.h.node ast/rewriter/rewriter_def.h.node
	@echo src/tactic/arith/factor_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/factor_tactic$(OBJ_EXT) ../src/tactic/arith/factor_tactic.cpp
tactic/arith/eq2bv_tactic$(OBJ_EXT): ../src/tactic/arith/eq2bv_tactic.cpp tactic/tactical.h.node tactic/arith/bound_manager.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter_def.h.node ast/ast_util.h.node ast/ast_pp_util.h.node
	@echo src/tactic/arith/eq2bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/eq2bv_tactic$(OBJ_EXT) ../src/tactic/arith/eq2bv_tactic.cpp
tactic/arith/add_bounds_tactic$(OBJ_EXT): ../src/tactic/arith/add_bounds_tactic.cpp tactic/tactical.h.node ast/arith_decl_plugin.h.node ast/ast_smt2_pp.h.node tactic/arith/bound_manager.h.node
	@echo src/tactic/arith/add_bounds_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/add_bounds_tactic$(OBJ_EXT) ../src/tactic/arith/add_bounds_tactic.cpp
tactic/arith/pb2bv_model_converter$(OBJ_EXT): ../src/tactic/arith/pb2bv_model_converter.cpp util/trace.h.node ast/arith_decl_plugin.h.node model/model_v2_pp.h.node tactic/arith/pb2bv_model_converter.h.node
	@echo src/tactic/arith/pb2bv_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_34) $(CXX_OUT_FLAG)tactic/arith/pb2bv_model_converter$(OBJ_EXT) ../src/tactic/arith/pb2bv_model_converter.cpp
tactic/arith/arith_tactics$(LIB_EXT): tactic/arith/card2bv_tactic$(OBJ_EXT) tactic/arith/arith_bounds_tactic$(OBJ_EXT) tactic/arith/fm_tactic$(OBJ_EXT) tactic/arith/degree_shift_tactic$(OBJ_EXT) tactic/arith/lia2card_tactic$(OBJ_EXT) tactic/arith/linear_equation$(OBJ_EXT) tactic/arith/bv2int_rewriter$(OBJ_EXT) tactic/arith/pb2bv_tactic$(OBJ_EXT) tactic/arith/bv2real_rewriter$(OBJ_EXT) tactic/arith/diff_neq_tactic$(OBJ_EXT) tactic/arith/bound_propagator$(OBJ_EXT) tactic/arith/purify_arith_tactic$(OBJ_EXT) tactic/arith/recover_01_tactic$(OBJ_EXT) tactic/arith/bound_manager$(OBJ_EXT) tactic/arith/lia2pb_tactic$(OBJ_EXT) tactic/arith/normalize_bounds_tactic$(OBJ_EXT) tactic/arith/propagate_ineqs_tactic$(OBJ_EXT) tactic/arith/nla2bv_tactic$(OBJ_EXT) tactic/arith/probe_arith$(OBJ_EXT) tactic/arith/fix_dl_var_tactic$(OBJ_EXT) tactic/arith/factor_tactic$(OBJ_EXT) tactic/arith/eq2bv_tactic$(OBJ_EXT) tactic/arith/add_bounds_tactic$(OBJ_EXT) tactic/arith/pb2bv_model_converter$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/arith/arith_tactics$(LIB_EXT) tactic/arith/card2bv_tactic$(OBJ_EXT) tactic/arith/arith_bounds_tactic$(OBJ_EXT) tactic/arith/fm_tactic$(OBJ_EXT) tactic/arith/degree_shift_tactic$(OBJ_EXT) tactic/arith/lia2card_tactic$(OBJ_EXT) tactic/arith/linear_equation$(OBJ_EXT) tactic/arith/bv2int_rewriter$(OBJ_EXT) tactic/arith/pb2bv_tactic$(OBJ_EXT) tactic/arith/bv2real_rewriter$(OBJ_EXT) tactic/arith/diff_neq_tactic$(OBJ_EXT) tactic/arith/bound_propagator$(OBJ_EXT) tactic/arith/purify_arith_tactic$(OBJ_EXT) tactic/arith/recover_01_tactic$(OBJ_EXT) tactic/arith/bound_manager$(OBJ_EXT) tactic/arith/lia2pb_tactic$(OBJ_EXT) tactic/arith/normalize_bounds_tactic$(OBJ_EXT) tactic/arith/propagate_ineqs_tactic$(OBJ_EXT) tactic/arith/nla2bv_tactic$(OBJ_EXT) tactic/arith/probe_arith$(OBJ_EXT) tactic/arith/fix_dl_var_tactic$(OBJ_EXT) tactic/arith/factor_tactic$(OBJ_EXT) tactic/arith/eq2bv_tactic$(OBJ_EXT) tactic/arith/add_bounds_tactic$(OBJ_EXT) tactic/arith/pb2bv_model_converter$(OBJ_EXT)
arith_tactics: tactic/arith/arith_tactics$(LIB_EXT)

includes_35 = -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/math/simplex -I../src
qe/mbp/mbp_plugin.h.node: ../src/qe/mbp/mbp_plugin.h ast/ast.h.node util/params.h.node model/model.h.node math/simplex/model_based_opt.h.node
	@echo done > qe/mbp/mbp_plugin.h.node
qe/mbp/mbp_plugin$(OBJ_EXT): ../src/qe/mbp/mbp_plugin.cpp ast/rewriter/expr_safe_replace.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/occurs.h.node ast/rewriter/th_rewriter.h.node ast/expr_functors.h.node ast/for_each_expr.h.node ast/scoped_proof.h.node qe/mbp/mbp_plugin.h.node model/model_pp.h.node model/model_evaluator.h.node
	@echo src/qe/mbp/mbp_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_35) $(CXX_OUT_FLAG)qe/mbp/mbp_plugin$(OBJ_EXT) ../src/qe/mbp/mbp_plugin.cpp
ast/is_variable_test.h.node: ../src/ast/is_variable_test.h ast/ast.h.node util/uint_set.h.node
	@echo done > ast/is_variable_test.h.node
qe/mbp/mbp_solve_plugin.h.node: ../src/qe/mbp/mbp_solve_plugin.h ast/ast.h.node ast/is_variable_test.h.node util/plugin_manager.h.node
	@echo done > qe/mbp/mbp_solve_plugin.h.node
qe/mbp/mbp_term_graph.h.node: ../src/qe/mbp/mbp_term_graph.h ast/ast.h.node ast/is_variable_test.h.node util/plugin_manager.h.node qe/mbp/mbp_solve_plugin.h.node model/model.h.node
	@echo done > qe/mbp/mbp_term_graph.h.node
qe/mbp/mbp_term_graph$(OBJ_EXT): ../src/qe/mbp/mbp_term_graph.cpp util/util.h.node util/uint_set.h.node util/obj_pair_hashtable.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/for_each_expr.h.node ast/occurs.h.node model/model_evaluator.h.node qe/mbp/mbp_term_graph.h.node
	@echo src/qe/mbp/mbp_term_graph.cpp
	@$(CXX) $(CXXFLAGS) $(includes_35) $(CXX_OUT_FLAG)qe/mbp/mbp_term_graph$(OBJ_EXT) ../src/qe/mbp/mbp_term_graph.cpp
qe/mbp/mbp_arrays.h.node: ../src/qe/mbp/mbp_arrays.h ast/array_decl_plugin.h.node qe/mbp/mbp_plugin.h.node
	@echo done > qe/mbp/mbp_arrays.h.node
qe/mbp/mbp_arrays$(OBJ_EXT): ../src/qe/mbp/mbp_arrays.cpp util/lbool.h.node ast/rewriter/rewriter_def.h.node ast/expr_functors.h.node ast/for_each_expr.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/th_rewriter.h.node ast/ast_util.h.node ast/ast_pp.h.node model/model_evaluator.h.node qe/mbp/mbp_arrays.h.node qe/mbp/mbp_term_graph.h.node
	@echo src/qe/mbp/mbp_arrays.cpp
	@$(CXX) $(CXXFLAGS) $(includes_35) $(CXX_OUT_FLAG)qe/mbp/mbp_arrays$(OBJ_EXT) ../src/qe/mbp/mbp_arrays.cpp
qe/mbp/mbp_arith.h.node: ../src/qe/mbp/mbp_arith.h ast/arith_decl_plugin.h.node model/model.h.node qe/mbp/mbp_plugin.h.node
	@echo done > qe/mbp/mbp_arith.h.node
qe/mbp/mbp_arith$(OBJ_EXT): ../src/qe/mbp/mbp_arith.cpp qe/mbp/mbp_arith.h.node ast/ast_util.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/th_rewriter.h.node ast/expr_functors.h.node ast/rewriter/expr_safe_replace.h.node math/simplex/model_based_opt.h.node model/model_evaluator.h.node model/model_smt2_pp.h.node model/model_v2_pp.h.node
	@echo src/qe/mbp/mbp_arith.cpp
	@$(CXX) $(CXXFLAGS) $(includes_35) $(CXX_OUT_FLAG)qe/mbp/mbp_arith$(OBJ_EXT) ../src/qe/mbp/mbp_arith.cpp
qe/mbp/mbp_solve_plugin$(OBJ_EXT): ../src/qe/mbp/mbp_solve_plugin.cpp ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/ast_util.h.node ast/ast_pp.h.node qe/mbp/mbp_solve_plugin.h.node
	@echo src/qe/mbp/mbp_solve_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_35) $(CXX_OUT_FLAG)qe/mbp/mbp_solve_plugin$(OBJ_EXT) ../src/qe/mbp/mbp_solve_plugin.cpp
qe/mbp/mbp_datatypes.h.node: ../src/qe/mbp/mbp_datatypes.h ast/datatype_decl_plugin.h.node qe/mbp/mbp_plugin.h.node
	@echo done > qe/mbp/mbp_datatypes.h.node
qe/mbp/mbp_datatypes$(OBJ_EXT): ../src/qe/mbp/mbp_datatypes.cpp ast/ast_pp.h.node ast/rewriter/th_rewriter.h.node ast/expr_functors.h.node model/model_v2_pp.h.node ast/rewriter/expr_safe_replace.h.node util/obj_pair_hashtable.h.node qe/mbp/mbp_datatypes.h.node
	@echo src/qe/mbp/mbp_datatypes.cpp
	@$(CXX) $(CXXFLAGS) $(includes_35) $(CXX_OUT_FLAG)qe/mbp/mbp_datatypes$(OBJ_EXT) ../src/qe/mbp/mbp_datatypes.cpp
qe/mbp/mbp$(LIB_EXT): qe/mbp/mbp_plugin$(OBJ_EXT) qe/mbp/mbp_term_graph$(OBJ_EXT) qe/mbp/mbp_arrays$(OBJ_EXT) qe/mbp/mbp_arith$(OBJ_EXT) qe/mbp/mbp_solve_plugin$(OBJ_EXT) qe/mbp/mbp_datatypes$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)qe/mbp/mbp$(LIB_EXT) qe/mbp/mbp_plugin$(OBJ_EXT) qe/mbp/mbp_term_graph$(OBJ_EXT) qe/mbp/mbp_arrays$(OBJ_EXT) qe/mbp/mbp_arith$(OBJ_EXT) qe/mbp/mbp_solve_plugin$(OBJ_EXT) qe/mbp/mbp_datatypes$(OBJ_EXT)
mbp: qe/mbp/mbp$(LIB_EXT)

includes_36 = -I../src/sat -I../src/util -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/ast -I../src/math/polynomial -I../src/math/simplex -I../src/ast/euf -I../src/tactic -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/ast/proofs -I../src/smt/params -I../src/ast/rewriter/bit_blaster -I../src/ast/fpa -I../src/qe/mbp -I../src/ast/normal_forms -I../src/math/lp -I../src/nlsat -I../src
sat/smt/atom2bool_var.h.node: ../src/sat/smt/atom2bool_var.h ast/expr2var.h.node sat/sat_types.h.node
	@echo done > sat/smt/atom2bool_var.h.node
sat/smt/atom2bool_var$(OBJ_EXT): ../src/sat/smt/atom2bool_var.cpp util/ref_util.h.node ast/ast_smt2_pp.h.node tactic/goal.h.node sat/smt/atom2bool_var.h.node
	@echo src/sat/smt/atom2bool_var.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/atom2bool_var$(OBJ_EXT) ../src/sat/smt/atom2bool_var.cpp
sat/smt/sat_smt.h.node: ../src/sat/smt/sat_smt.h ast/ast.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node sat/sat_solver.h.node
	@echo done > sat/smt/sat_smt.h.node
sat/smt/sat_th.h.node: ../src/sat/smt/sat_th.h util/top_sort.h.node sat/smt/sat_smt.h.node ast/euf/euf_egraph.h.node model/model.h.node smt/params/smt_params.h.node
	@echo done > sat/smt/sat_th.h.node
sat/smt/q_model_fixer.h.node: ../src/sat/smt/q_model_fixer.h sat/smt/sat_th.h.node solver/solver.h.node model/model_macro_solver.h.node qe/mbp/mbp_plugin.h.node
	@echo done > sat/smt/q_model_fixer.h.node
ast/ast_trail.h.node: ../src/ast/ast_trail.h ast/ast.h.node util/trail.h.node
	@echo done > ast/ast_trail.h.node
sat/smt/q_mbi.h.node: ../src/sat/smt/q_mbi.h solver/solver.h.node qe/mbp/mbp_plugin.h.node sat/smt/sat_th.h.node sat/smt/q_model_fixer.h.node
	@echo done > sat/smt/q_mbi.h.node
sat/smt/q_solver.h.node: ../src/sat/smt/q_solver.h util/obj_hashtable.h.node ast/ast_trail.h.node sat/smt/sat_th.h.node sat/smt/q_mbi.h.node
	@echo done > sat/smt/q_solver.h.node
util/lim_vector.h.node: ../src/util/lim_vector.h util/vector.h.node
	@echo done > util/lim_vector.h.node
sat/smt/sat_dual_solver.h.node: ../src/sat/smt/sat_dual_solver.h util/lim_vector.h.node sat/sat_solver.h.node
	@echo done > sat/smt/sat_dual_solver.h.node
util/dlist.h.node: ../src/util/dlist.h
	@echo done > util/dlist.h.node
sat/smt/euf_ackerman.h.node: ../src/sat/smt/euf_ackerman.h util/dlist.h.node ast/euf/euf_egraph.h.node sat/smt/atom2bool_var.h.node sat/smt/sat_th.h.node
	@echo done > sat/smt/euf_ackerman.h.node
sat/smt/user_solver.h.node: ../src/sat/smt/user_solver.h sat/smt/sat_th.h.node solver/solver.h.node
	@echo done > sat/smt/user_solver.h.node
sat/smt/euf_solver.h.node: ../src/sat/smt/euf_solver.h util/scoped_ptr_vector.h.node util/trail.h.node ast/ast_translation.h.node ast/euf/euf_egraph.h.node ast/rewriter/th_rewriter.h.node tactic/model_converter.h.node sat/sat_extension.h.node sat/smt/atom2bool_var.h.node sat/smt/sat_th.h.node sat/smt/sat_dual_solver.h.node sat/smt/euf_ackerman.h.node sat/smt/user_solver.h.node smt/params/smt_params.h.node
	@echo done > sat/smt/euf_solver.h.node
sat/smt/q_model_fixer$(OBJ_EXT): ../src/sat/smt/q_model_fixer.cpp ast/for_each_expr.h.node ast/ast_util.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node model/model_macro_solver.h.node sat/smt/q_model_fixer.h.node sat/smt/q_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/q_model_fixer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/q_model_fixer$(OBJ_EXT) ../src/sat/smt/q_model_fixer.cpp
sat/smt/euf_model$(OBJ_EXT): ../src/sat/smt/euf_model.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node sat/smt/euf_solver.h.node model/value_factory.h.node
	@echo src/sat/smt/euf_model.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_model$(OBJ_EXT) ../src/sat/smt/euf_model.cpp
sat/smt/bv_ackerman.h.node: ../src/sat/smt/bv_ackerman.h util/dlist.h.node sat/smt/atom2bool_var.h.node sat/smt/sat_th.h.node
	@echo done > sat/smt/bv_ackerman.h.node
sat/smt/bv_solver.h.node: ../src/sat/smt/bv_solver.h sat/smt/sat_th.h.node sat/smt/bv_ackerman.h.node ast/rewriter/bit_blaster/bit_blaster.h.node
	@echo done > sat/smt/bv_solver.h.node
sat/smt/bv_solver$(OBJ_EXT): ../src/sat/smt/bv_solver.cpp ast/ast_ll_pp.h.node sat/smt/bv_solver.h.node sat/smt/euf_solver.h.node sat/smt/sat_th.h.node tactic/tactic_exception.h.node
	@echo src/sat/smt/bv_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/bv_solver$(OBJ_EXT) ../src/sat/smt/bv_solver.cpp
sat/smt/dt_solver.h.node: ../src/sat/smt/dt_solver.h sat/smt/sat_th.h.node ast/datatype_decl_plugin.h.node ast/array_decl_plugin.h.node
	@echo done > sat/smt/dt_solver.h.node
sat/smt/array_solver.h.node: ../src/sat/smt/array_solver.h ast/ast_trail.h.node sat/smt/sat_th.h.node ast/array_decl_plugin.h.node
	@echo done > sat/smt/array_solver.h.node
sat/smt/dt_solver$(OBJ_EXT): ../src/sat/smt/dt_solver.cpp sat/smt/dt_solver.h.node sat/smt/euf_solver.h.node sat/smt/array_solver.h.node
	@echo src/sat/smt/dt_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/dt_solver$(OBJ_EXT) ../src/sat/smt/dt_solver.cpp
sat/smt/ba_solver_interface.h.node: ../src/sat/smt/ba_solver_interface.h sat/sat_types.h.node sat/sat_solver.h.node sat/smt/sat_smt.h.node
	@echo done > sat/smt/ba_solver_interface.h.node
sat/smt/ba_constraint.h.node: ../src/sat/smt/ba_constraint.h sat/smt/ba_solver_interface.h.node
	@echo done > sat/smt/ba_constraint.h.node
sat/smt/ba_card.h.node: ../src/sat/smt/ba_card.h sat/sat_types.h.node sat/smt/ba_constraint.h.node
	@echo done > sat/smt/ba_card.h.node
sat/smt/ba_pb.h.node: ../src/sat/smt/ba_pb.h sat/sat_types.h.node sat/smt/ba_card.h.node
	@echo done > sat/smt/ba_pb.h.node
sat/smt/ba_xor.h.node: ../src/sat/smt/ba_xor.h sat/sat_types.h.node sat/smt/ba_constraint.h.node
	@echo done > sat/smt/ba_xor.h.node
sat/smt/ba_solver.h.node: ../src/sat/smt/ba_solver.h sat/sat_extension.h.node sat/sat_solver.h.node sat/sat_lookahead.h.node sat/sat_big.h.node sat/smt/sat_smt.h.node sat/smt/sat_th.h.node sat/smt/ba_constraint.h.node sat/smt/ba_card.h.node sat/smt/ba_pb.h.node sat/smt/ba_xor.h.node util/small_object_allocator.h.node util/scoped_ptr_vector.h.node util/sorting_network.h.node ast/pb_decl_plugin.h.node
	@echo done > sat/smt/ba_solver.h.node
sat/smt/ba_solver$(OBJ_EXT): ../src/sat/smt/ba_solver.cpp util/mpz.h.node sat/sat_types.h.node sat/smt/ba_solver.h.node sat/smt/euf_solver.h.node sat/sat_simplifier_params.hpp.node sat/sat_xor_finder.h.node
	@echo src/sat/smt/ba_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/ba_solver$(OBJ_EXT) ../src/sat/smt/ba_solver.cpp
sat/smt/q_mbi$(OBJ_EXT): ../src/sat/smt/q_mbi.cpp ast/ast_trail.h.node ast/ast_util.h.node ast/for_each_expr.h.node ast/rewriter/var_subst.h.node ast/rewriter/expr_safe_replace.h.node qe/mbp/mbp_arith.h.node qe/mbp/mbp_arrays.h.node qe/mbp/mbp_datatypes.h.node sat/smt/sat_th.h.node sat/smt/q_mbi.h.node sat/smt/q_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/q_mbi.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/q_mbi$(OBJ_EXT) ../src/sat/smt/q_mbi.cpp
sat/smt/ba_xor$(OBJ_EXT): ../src/sat/smt/ba_xor.cpp sat/smt/ba_xor.h.node sat/smt/ba_solver.h.node
	@echo src/sat/smt/ba_xor.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/ba_xor$(OBJ_EXT) ../src/sat/smt/ba_xor.cpp
sat/smt/array_axioms$(OBJ_EXT): ../src/sat/smt/array_axioms.cpp ast/ast_trail.h.node ast/ast_ll_pp.h.node sat/smt/array_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/array_axioms.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/array_axioms$(OBJ_EXT) ../src/sat/smt/array_axioms.cpp
sat/smt/euf_proof$(OBJ_EXT): ../src/sat/smt/euf_proof.cpp sat/smt/euf_solver.h.node
	@echo src/sat/smt/euf_proof.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_proof$(OBJ_EXT) ../src/sat/smt/euf_proof.cpp
sat/smt/ba_constraint$(OBJ_EXT): ../src/sat/smt/ba_constraint.cpp sat/smt/ba_constraint.h.node
	@echo src/sat/smt/ba_constraint.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/ba_constraint$(OBJ_EXT) ../src/sat/smt/ba_constraint.cpp
sat/smt/ba_internalize$(OBJ_EXT): ../src/sat/smt/ba_internalize.cpp sat/smt/ba_solver.h.node ast/pb_decl_plugin.h.node
	@echo src/sat/smt/ba_internalize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/ba_internalize$(OBJ_EXT) ../src/sat/smt/ba_internalize.cpp
sat/smt/array_solver$(OBJ_EXT): ../src/sat/smt/array_solver.cpp ast/ast_ll_pp.h.node sat/smt/array_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/array_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/array_solver$(OBJ_EXT) ../src/sat/smt/array_solver.cpp
sat/smt/bv_invariant$(OBJ_EXT): ../src/sat/smt/bv_invariant.cpp sat/smt/bv_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/bv_invariant.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/bv_invariant$(OBJ_EXT) ../src/sat/smt/bv_invariant.cpp
util/obj_pair_set.h.node: ../src/util/obj_pair_set.h util/chashtable.h.node
	@echo done > util/obj_pair_set.h.node
math/lp/lp_api.h.node: ../src/math/lp/lp_api.h util/inf_rational.h.node util/optional.h.node
	@echo done > math/lp/lp_api.h.node
sat/smt/arith_solver.h.node: ../src/sat/smt/arith_solver.h util/obj_pair_set.h.node ast/ast_trail.h.node ast/arith_decl_plugin.h.node math/lp/lp_solver.h.node math/lp/lp_primal_simplex.h.node math/lp/lp_dual_simplex.h.node math/lp/indexed_value.h.node math/lp/lar_solver.h.node math/lp/nla_solver.h.node math/lp/lp_types.h.node math/lp/lp_api.h.node math/polynomial/algebraic_numbers.h.node math/polynomial/polynomial.h.node sat/smt/sat_th.h.node
	@echo done > sat/smt/arith_solver.h.node
sat/smt/fpa_solver.h.node: ../src/sat/smt/fpa_solver.h sat/smt/euf_solver.h.node ast/fpa/fpa2bv_converter.h.node ast/fpa/fpa2bv_rewriter.h.node
	@echo done > sat/smt/fpa_solver.h.node
sat/smt/euf_solver$(OBJ_EXT): ../src/sat/smt/euf_solver.cpp ast/pb_decl_plugin.h.node ast/ast_ll_pp.h.node sat/sat_solver.h.node sat/smt/sat_smt.h.node sat/smt/ba_solver.h.node sat/smt/bv_solver.h.node sat/smt/euf_solver.h.node sat/smt/array_solver.h.node sat/smt/arith_solver.h.node sat/smt/q_solver.h.node sat/smt/fpa_solver.h.node sat/smt/dt_solver.h.node
	@echo src/sat/smt/euf_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_solver$(OBJ_EXT) ../src/sat/smt/euf_solver.cpp
sat/smt/euf_relevancy$(OBJ_EXT): ../src/sat/smt/euf_relevancy.cpp sat/smt/euf_solver.h.node
	@echo src/sat/smt/euf_relevancy.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_relevancy$(OBJ_EXT) ../src/sat/smt/euf_relevancy.cpp
sat/smt/arith_solver$(OBJ_EXT): ../src/sat/smt/arith_solver.cpp sat/smt/euf_solver.h.node sat/smt/arith_solver.h.node
	@echo src/sat/smt/arith_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/arith_solver$(OBJ_EXT) ../src/sat/smt/arith_solver.cpp
sat/smt/euf_internalize$(OBJ_EXT): ../src/sat/smt/euf_internalize.cpp ast/pb_decl_plugin.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/euf_internalize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_internalize$(OBJ_EXT) ../src/sat/smt/euf_internalize.cpp
sat/smt/sat_dual_solver$(OBJ_EXT): ../src/sat/smt/sat_dual_solver.cpp sat/smt/sat_dual_solver.h.node
	@echo src/sat/smt/sat_dual_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/sat_dual_solver$(OBJ_EXT) ../src/sat/smt/sat_dual_solver.cpp
sat/smt/ba_card$(OBJ_EXT): ../src/sat/smt/ba_card.cpp sat/smt/ba_card.h.node sat/smt/ba_solver.h.node sat/sat_simplifier.h.node
	@echo src/sat/smt/ba_card.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/ba_card$(OBJ_EXT) ../src/sat/smt/ba_card.cpp
sat/smt/user_solver$(OBJ_EXT): ../src/sat/smt/user_solver.cpp sat/smt/user_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/user_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/user_solver$(OBJ_EXT) ../src/sat/smt/user_solver.cpp
sat/smt/xor_solver$(OBJ_EXT): ../src/sat/smt/xor_solver.cpp sat/sat_types.h.node sat/smt/ba_solver.h.node sat/sat_simplifier_params.hpp.node sat/sat_xor_finder.h.node
	@echo src/sat/smt/xor_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/xor_solver$(OBJ_EXT) ../src/sat/smt/xor_solver.cpp
sat/smt/q_solver$(OBJ_EXT): ../src/sat/smt/q_solver.cpp ast/rewriter/var_subst.h.node sat/smt/q_solver.h.node sat/smt/euf_solver.h.node sat/smt/sat_th.h.node ast/normal_forms/pull_quant.h.node ast/well_sorted.h.node
	@echo src/sat/smt/q_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/q_solver$(OBJ_EXT) ../src/sat/smt/q_solver.cpp
sat/smt/bv_internalize$(OBJ_EXT): ../src/sat/smt/bv_internalize.cpp params/bv_rewriter_params.hpp.node sat/smt/bv_solver.h.node sat/smt/euf_solver.h.node sat/smt/sat_th.h.node tactic/tactic_exception.h.node
	@echo src/sat/smt/bv_internalize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/bv_internalize$(OBJ_EXT) ../src/sat/smt/bv_internalize.cpp
sat/smt/bv_delay_internalize$(OBJ_EXT): ../src/sat/smt/bv_delay_internalize.cpp sat/smt/bv_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/bv_delay_internalize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/bv_delay_internalize$(OBJ_EXT) ../src/sat/smt/bv_delay_internalize.cpp
sat/smt/arith_diagnostics$(OBJ_EXT): ../src/sat/smt/arith_diagnostics.cpp sat/smt/euf_solver.h.node sat/smt/arith_solver.h.node
	@echo src/sat/smt/arith_diagnostics.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/arith_diagnostics$(OBJ_EXT) ../src/sat/smt/arith_diagnostics.cpp
sat/smt/sat_th$(OBJ_EXT): ../src/sat/smt/sat_th.cpp sat/smt/sat_th.h.node sat/smt/euf_solver.h.node tactic/tactic_exception.h.node
	@echo src/sat/smt/sat_th.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/sat_th$(OBJ_EXT) ../src/sat/smt/sat_th.cpp
sat/smt/array_model$(OBJ_EXT): ../src/sat/smt/array_model.cpp model/array_factory.h.node sat/smt/array_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/array_model.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/array_model$(OBJ_EXT) ../src/sat/smt/array_model.cpp
sat/smt/bv_ackerman$(OBJ_EXT): ../src/sat/smt/bv_ackerman.cpp sat/smt/bv_solver.h.node sat/smt/bv_ackerman.h.node
	@echo src/sat/smt/bv_ackerman.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/bv_ackerman$(OBJ_EXT) ../src/sat/smt/bv_ackerman.cpp
sat/smt/ba_pb$(OBJ_EXT): ../src/sat/smt/ba_pb.cpp sat/smt/ba_pb.h.node
	@echo src/sat/smt/ba_pb.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/ba_pb$(OBJ_EXT) ../src/sat/smt/ba_pb.cpp
sat/smt/arith_internalize$(OBJ_EXT): ../src/sat/smt/arith_internalize.cpp sat/smt/euf_solver.h.node sat/smt/arith_solver.h.node
	@echo src/sat/smt/arith_internalize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/arith_internalize$(OBJ_EXT) ../src/sat/smt/arith_internalize.cpp
sat/smt/euf_invariant$(OBJ_EXT): ../src/sat/smt/euf_invariant.cpp sat/smt/euf_solver.h.node
	@echo src/sat/smt/euf_invariant.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_invariant$(OBJ_EXT) ../src/sat/smt/euf_invariant.cpp
sat/smt/arith_axioms$(OBJ_EXT): ../src/sat/smt/arith_axioms.cpp sat/smt/euf_solver.h.node sat/smt/arith_solver.h.node
	@echo src/sat/smt/arith_axioms.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/arith_axioms$(OBJ_EXT) ../src/sat/smt/arith_axioms.cpp
sat/smt/array_internalize$(OBJ_EXT): ../src/sat/smt/array_internalize.cpp sat/smt/array_solver.h.node sat/smt/euf_solver.h.node
	@echo src/sat/smt/array_internalize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/array_internalize$(OBJ_EXT) ../src/sat/smt/array_internalize.cpp
sat/smt/euf_ackerman$(OBJ_EXT): ../src/sat/smt/euf_ackerman.cpp sat/smt/euf_solver.h.node sat/smt/euf_ackerman.h.node
	@echo src/sat/smt/euf_ackerman.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/euf_ackerman$(OBJ_EXT) ../src/sat/smt/euf_ackerman.cpp
sat/smt/fpa_solver$(OBJ_EXT): ../src/sat/smt/fpa_solver.cpp sat/smt/fpa_solver.h.node ast/fpa/bv2fpa_converter.h.node
	@echo src/sat/smt/fpa_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_36) $(CXX_OUT_FLAG)sat/smt/fpa_solver$(OBJ_EXT) ../src/sat/smt/fpa_solver.cpp
sat/smt/sat_smt$(LIB_EXT): sat/smt/atom2bool_var$(OBJ_EXT) sat/smt/q_model_fixer$(OBJ_EXT) sat/smt/euf_model$(OBJ_EXT) sat/smt/bv_solver$(OBJ_EXT) sat/smt/dt_solver$(OBJ_EXT) sat/smt/ba_solver$(OBJ_EXT) sat/smt/q_mbi$(OBJ_EXT) sat/smt/ba_xor$(OBJ_EXT) sat/smt/array_axioms$(OBJ_EXT) sat/smt/euf_proof$(OBJ_EXT) sat/smt/ba_constraint$(OBJ_EXT) sat/smt/ba_internalize$(OBJ_EXT) sat/smt/array_solver$(OBJ_EXT) sat/smt/bv_invariant$(OBJ_EXT) sat/smt/euf_solver$(OBJ_EXT) sat/smt/euf_relevancy$(OBJ_EXT) sat/smt/arith_solver$(OBJ_EXT) sat/smt/euf_internalize$(OBJ_EXT) sat/smt/sat_dual_solver$(OBJ_EXT) sat/smt/ba_card$(OBJ_EXT) sat/smt/user_solver$(OBJ_EXT) sat/smt/xor_solver$(OBJ_EXT) sat/smt/q_solver$(OBJ_EXT) sat/smt/bv_internalize$(OBJ_EXT) sat/smt/bv_delay_internalize$(OBJ_EXT) sat/smt/arith_diagnostics$(OBJ_EXT) sat/smt/sat_th$(OBJ_EXT) sat/smt/array_model$(OBJ_EXT) sat/smt/bv_ackerman$(OBJ_EXT) sat/smt/ba_pb$(OBJ_EXT) sat/smt/arith_internalize$(OBJ_EXT) sat/smt/euf_invariant$(OBJ_EXT) sat/smt/arith_axioms$(OBJ_EXT) sat/smt/array_internalize$(OBJ_EXT) sat/smt/euf_ackerman$(OBJ_EXT) sat/smt/fpa_solver$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)sat/smt/sat_smt$(LIB_EXT) sat/smt/atom2bool_var$(OBJ_EXT) sat/smt/q_model_fixer$(OBJ_EXT) sat/smt/euf_model$(OBJ_EXT) sat/smt/bv_solver$(OBJ_EXT) sat/smt/dt_solver$(OBJ_EXT) sat/smt/ba_solver$(OBJ_EXT) sat/smt/q_mbi$(OBJ_EXT) sat/smt/ba_xor$(OBJ_EXT) sat/smt/array_axioms$(OBJ_EXT) sat/smt/euf_proof$(OBJ_EXT) sat/smt/ba_constraint$(OBJ_EXT) sat/smt/ba_internalize$(OBJ_EXT) sat/smt/array_solver$(OBJ_EXT) sat/smt/bv_invariant$(OBJ_EXT) sat/smt/euf_solver$(OBJ_EXT) sat/smt/euf_relevancy$(OBJ_EXT) sat/smt/arith_solver$(OBJ_EXT) sat/smt/euf_internalize$(OBJ_EXT) sat/smt/sat_dual_solver$(OBJ_EXT) sat/smt/ba_card$(OBJ_EXT) sat/smt/user_solver$(OBJ_EXT) sat/smt/xor_solver$(OBJ_EXT) sat/smt/q_solver$(OBJ_EXT) sat/smt/bv_internalize$(OBJ_EXT) sat/smt/bv_delay_internalize$(OBJ_EXT) sat/smt/arith_diagnostics$(OBJ_EXT) sat/smt/sat_th$(OBJ_EXT) sat/smt/array_model$(OBJ_EXT) sat/smt/bv_ackerman$(OBJ_EXT) sat/smt/ba_pb$(OBJ_EXT) sat/smt/arith_internalize$(OBJ_EXT) sat/smt/euf_invariant$(OBJ_EXT) sat/smt/arith_axioms$(OBJ_EXT) sat/smt/array_internalize$(OBJ_EXT) sat/smt/euf_ackerman$(OBJ_EXT) sat/smt/fpa_solver$(OBJ_EXT)
sat_smt: sat/smt/sat_smt$(LIB_EXT)

includes_37 = -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/solver -I../src/ast/proofs -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/rewriter/bit_blaster -I../src/ast/fpa -I../src/qe/mbp -I../src/ast/normal_forms -I../src/math/lp -I../src/nlsat -I../src
sat/tactic/goal2sat.h.node: ../src/sat/tactic/goal2sat.h tactic/goal.h.node sat/sat_solver.h.node tactic/model_converter.h.node tactic/generic_model_converter.h.node sat/smt/atom2bool_var.h.node sat/smt/sat_smt.h.node
	@echo done > sat/tactic/goal2sat.h.node
sat/tactic/goal2sat$(OBJ_EXT): ../src/sat/tactic/goal2sat.cpp util/ref_util.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node ast/pb_decl_plugin.h.node ast/ast_util.h.node ast/for_each_expr.h.node model/model_evaluator.h.node model/model_v2_pp.h.node tactic/tactic.h.node tactic/generic_model_converter.h.node sat/sat_cut_simplifier.h.node sat/sat_drat.h.node sat/tactic/goal2sat.h.node sat/smt/ba_solver.h.node sat/smt/euf_solver.h.node sat/smt/sat_th.h.node sat/sat_params.hpp.node
	@echo src/sat/tactic/goal2sat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_37) $(CXX_OUT_FLAG)sat/tactic/goal2sat$(OBJ_EXT) ../src/sat/tactic/goal2sat.cpp
sat/tactic/sat_tactic$(OBJ_EXT): ../src/sat/tactic/sat_tactic.cpp ast/ast_pp.h.node model/model_v2_pp.h.node tactic/tactical.h.node sat/tactic/goal2sat.h.node sat/sat_solver.h.node sat/sat_params.hpp.node
	@echo src/sat/tactic/sat_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_37) $(CXX_OUT_FLAG)sat/tactic/sat_tactic$(OBJ_EXT) ../src/sat/tactic/sat_tactic.cpp
sat/tactic/sat_tactic$(LIB_EXT): sat/tactic/goal2sat$(OBJ_EXT) sat/tactic/sat_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)sat/tactic/sat_tactic$(LIB_EXT) sat/tactic/goal2sat$(OBJ_EXT) sat/tactic/sat_tactic$(OBJ_EXT)
sat_tactic: sat/tactic/sat_tactic$(LIB_EXT)

includes_38 = -I../src/nlsat -I../src/math/polynomial -I../src/util -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/ast -I../src/math/simplex -I../src/sat/tactic -I../src/tactic -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/ast/proofs -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/rewriter/bit_blaster -I../src/ast/fpa -I../src/qe/mbp -I../src/ast/normal_forms -I../src/math/lp -I../src/tactic/arith -I../src/tactic/core -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src
nlsat/tactic/goal2nlsat.h.node: ../src/nlsat/tactic/goal2nlsat.h nlsat/nlsat_types.h.node tactic/model_converter.h.node
	@echo done > nlsat/tactic/goal2nlsat.h.node
nlsat/tactic/goal2nlsat$(OBJ_EXT): ../src/nlsat/tactic/goal2nlsat.cpp nlsat/tactic/goal2nlsat.h.node tactic/goal.h.node tactic/goal_util.h.node nlsat/nlsat_solver.h.node ast/expr2polynomial.h.node ast/expr2var.h.node ast/arith_decl_plugin.h.node tactic/tactic.h.node ast/ast_pp.h.node math/polynomial/polynomial.h.node math/polynomial/algebraic_numbers.h.node
	@echo src/nlsat/tactic/goal2nlsat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_38) $(CXX_OUT_FLAG)nlsat/tactic/goal2nlsat$(OBJ_EXT) ../src/nlsat/tactic/goal2nlsat.cpp
nlsat/tactic/nlsat_tactic$(OBJ_EXT): ../src/nlsat/tactic/nlsat_tactic.cpp tactic/tactical.h.node nlsat/tactic/goal2nlsat.h.node nlsat/nlsat_solver.h.node model/model.h.node ast/expr2var.h.node ast/arith_decl_plugin.h.node ast/ast_smt2_pp.h.node util/z3_exception.h.node math/polynomial/algebraic_numbers.h.node ast/ast_pp.h.node
	@echo src/nlsat/tactic/nlsat_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_38) $(CXX_OUT_FLAG)nlsat/tactic/nlsat_tactic$(OBJ_EXT) ../src/nlsat/tactic/nlsat_tactic.cpp
tactic/core/tseitin_cnf_tactic.h.node: ../src/tactic/core/tseitin_cnf_tactic.h util/params.h.node
	@echo done > tactic/core/tseitin_cnf_tactic.h.node
tactic/arith/degree_shift_tactic.h.node: ../src/tactic/arith/degree_shift_tactic.h util/params.h.node
	@echo done > tactic/arith/degree_shift_tactic.h.node
tactic/arith/purify_arith_tactic.h.node: ../src/tactic/arith/purify_arith_tactic.h util/params.h.node
	@echo done > tactic/arith/purify_arith_tactic.h.node
nlsat/tactic/nlsat_tactic.h.node: ../src/nlsat/tactic/nlsat_tactic.h util/params.h.node
	@echo done > nlsat/tactic/nlsat_tactic.h.node
tactic/arith/factor_tactic.h.node: ../src/tactic/arith/factor_tactic.h util/params.h.node
	@echo done > tactic/arith/factor_tactic.h.node
tactic/core/elim_uncnstr_tactic.h.node: ../src/tactic/core/elim_uncnstr_tactic.h util/params.h.node
	@echo done > tactic/core/elim_uncnstr_tactic.h.node
tactic/core/solve_eqs_tactic.h.node: ../src/tactic/core/solve_eqs_tactic.h util/params.h.node
	@echo done > tactic/core/solve_eqs_tactic.h.node
tactic/core/elim_term_ite_tactic.h.node: ../src/tactic/core/elim_term_ite_tactic.h util/params.h.node
	@echo done > tactic/core/elim_term_ite_tactic.h.node
nlsat/tactic/qfnra_nlsat_tactic$(OBJ_EXT): ../src/nlsat/tactic/qfnra_nlsat_tactic.cpp tactic/tactical.h.node tactic/core/tseitin_cnf_tactic.h.node tactic/arith/degree_shift_tactic.h.node tactic/arith/purify_arith_tactic.h.node nlsat/tactic/nlsat_tactic.h.node tactic/arith/factor_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_term_ite_tactic.h.node
	@echo src/nlsat/tactic/qfnra_nlsat_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_38) $(CXX_OUT_FLAG)nlsat/tactic/qfnra_nlsat_tactic$(OBJ_EXT) ../src/nlsat/tactic/qfnra_nlsat_tactic.cpp
nlsat/tactic/nlsat_tactic$(LIB_EXT): nlsat/tactic/goal2nlsat$(OBJ_EXT) nlsat/tactic/nlsat_tactic$(OBJ_EXT) nlsat/tactic/qfnra_nlsat_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)nlsat/tactic/nlsat_tactic$(LIB_EXT) nlsat/tactic/goal2nlsat$(OBJ_EXT) nlsat/tactic/nlsat_tactic$(OBJ_EXT) nlsat/tactic/qfnra_nlsat_tactic$(OBJ_EXT)
nlsat_tactic: nlsat/tactic/nlsat_tactic$(LIB_EXT)

includes_39 = -I../src/tactic/core -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src/math/subpaving -I../src/math/interval -I../src
math/subpaving/tactic/expr2subpaving.h.node: ../src/math/subpaving/tactic/expr2subpaving.h ast/ast.h.node math/subpaving/subpaving.h.node
	@echo done > math/subpaving/tactic/expr2subpaving.h.node
math/subpaving/tactic/expr2subpaving$(OBJ_EXT): ../src/math/subpaving/tactic/expr2subpaving.cpp math/subpaving/tactic/expr2subpaving.h.node ast/expr2var.h.node util/ref_util.h.node util/z3_exception.h.node ast/arith_decl_plugin.h.node util/scoped_numeral_buffer.h.node util/common_msgs.h.node
	@echo src/math/subpaving/tactic/expr2subpaving.cpp
	@$(CXX) $(CXXFLAGS) $(includes_39) $(CXX_OUT_FLAG)math/subpaving/tactic/expr2subpaving$(OBJ_EXT) ../src/math/subpaving/tactic/expr2subpaving.cpp
math/subpaving/tactic/subpaving_tactic$(OBJ_EXT): ../src/math/subpaving/tactic/subpaving_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node math/subpaving/tactic/expr2subpaving.h.node ast/expr2var.h.node ast/arith_decl_plugin.h.node ast/ast_smt2_pp.h.node util/hwf.h.node util/mpff.h.node util/mpfx.h.node util/f2n.h.node
	@echo src/math/subpaving/tactic/subpaving_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_39) $(CXX_OUT_FLAG)math/subpaving/tactic/subpaving_tactic$(OBJ_EXT) ../src/math/subpaving/tactic/subpaving_tactic.cpp
math/subpaving/tactic/subpaving_tactic$(LIB_EXT): math/subpaving/tactic/expr2subpaving$(OBJ_EXT) math/subpaving/tactic/subpaving_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)math/subpaving/tactic/subpaving_tactic$(LIB_EXT) math/subpaving/tactic/expr2subpaving$(OBJ_EXT) math/subpaving/tactic/subpaving_tactic$(OBJ_EXT)
subpaving_tactic: math/subpaving/tactic/subpaving_tactic$(LIB_EXT)

includes_40 = -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/smt/params -I../src
smt/proto_model/proto_model.h.node: ../src/smt/proto_model/proto_model.h model/model_core.h.node model/model_evaluator.h.node model/value_factory.h.node ast/arith_decl_plugin.h.node ast/func_decl_dependencies.h.node model/model.h.node util/params.h.node ast/rewriter/th_rewriter.h.node
	@echo done > smt/proto_model/proto_model.h.node
smt/proto_model/proto_model$(OBJ_EXT): ../src/smt/proto_model/proto_model.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/well_sorted.h.node ast/array_decl_plugin.h.node ast/used_symbols.h.node ast/for_each_expr.h.node ast/rewriter/var_subst.h.node model/model_params.hpp.node model/model_v2_pp.h.node smt/proto_model/proto_model.h.node
	@echo src/smt/proto_model/proto_model.cpp
	@$(CXX) $(CXXFLAGS) $(includes_40) $(CXX_OUT_FLAG)smt/proto_model/proto_model$(OBJ_EXT) ../src/smt/proto_model/proto_model.cpp
smt/proto_model/proto_model$(LIB_EXT): smt/proto_model/proto_model$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)smt/proto_model/proto_model$(LIB_EXT) smt/proto_model/proto_model$(OBJ_EXT)
proto_model: smt/proto_model/proto_model$(LIB_EXT)

includes_41 = -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src
util/list.h.node: ../src/util/list.h util/buffer.h.node util/region.h.node
	@echo done > util/list.h.node
smt/smt_types.h.node: ../src/smt/smt_types.h util/list.h.node util/vector.h.node util/hashtable.h.node util/lbool.h.node
	@echo done > smt/smt_types.h.node
smt/smt_literal.h.node: ../src/smt/smt_literal.h ast/ast.h.node smt/smt_types.h.node util/approx_set.h.node
	@echo done > smt/smt_literal.h.node
smt/smt_eq_justification.h.node: ../src/smt/smt_eq_justification.h smt/smt_literal.h.node util/tptr.h.node
	@echo done > smt/smt_eq_justification.h.node
smt/smt_justification.h.node: ../src/smt/smt_justification.h ast/ast.h.node smt/smt_types.h.node smt/smt_literal.h.node smt/smt_eq_justification.h.node
	@echo done > smt/smt_justification.h.node
smt/smt_clause.h.node: ../src/smt/smt_clause.h ast/ast.h.node smt/smt_literal.h.node util/tptr.h.node util/obj_hashtable.h.node smt/smt_justification.h.node
	@echo done > smt/smt_clause.h.node
smt/smt_setup.h.node: ../src/smt/smt_setup.h ast/ast.h.node smt/params/smt_params.h.node
	@echo done > smt/smt_setup.h.node
smt/smt_enode.h.node: ../src/smt/smt_enode.h util/id_var_list.h.node util/approx_set.h.node ast/ast.h.node smt/smt_types.h.node smt/smt_eq_justification.h.node
	@echo done > smt/smt_enode.h.node
smt/smt_cg_table.h.node: ../src/smt/smt_cg_table.h smt/smt_enode.h.node util/hashtable.h.node util/chashtable.h.node
	@echo done > smt/smt_cg_table.h.node
smt/smt_b_justification.h.node: ../src/smt/smt_b_justification.h smt/smt_literal.h.node smt/smt_clause.h.node
	@echo done > smt/smt_b_justification.h.node
smt/smt_bool_var_data.h.node: ../src/smt/smt_bool_var_data.h smt/smt_b_justification.h.node
	@echo done > smt/smt_bool_var_data.h.node
smt/smt_quantifier.h.node: ../src/smt/smt_quantifier.h ast/ast.h.node util/statistics.h.node util/params.h.node smt/smt_types.h.node
	@echo done > smt/smt_quantifier.h.node
smt/smt_theory.h.node: ../src/smt/smt_theory.h ast/ast_pp.h.node smt/smt_enode.h.node smt/smt_quantifier.h.node util/obj_hashtable.h.node util/statistics.h.node
	@echo done > smt/smt_theory.h.node
smt/smt_clause_proof.h.node: ../src/smt/smt_clause_proof.h smt/smt_theory.h.node smt/smt_clause.h.node
	@echo done > smt/smt_clause_proof.h.node
smt/smt_quantifier_stat.h.node: ../src/smt/smt_quantifier_stat.h ast/ast.h.node util/obj_hashtable.h.node util/approx_nat.h.node util/region.h.node
	@echo done > smt/smt_quantifier_stat.h.node
smt/smt_statistics.h.node: ../src/smt/smt_statistics.h
	@echo done > smt/smt_statistics.h.node
util/obj_triple_hashtable.h.node: ../src/util/obj_triple_hashtable.h util/hashtable.h.node
	@echo done > util/obj_triple_hashtable.h.node
smt/dyn_ack.h.node: ../src/smt/dyn_ack.h ast/ast.h.node smt/params/dyn_ack_params.h.node util/obj_hashtable.h.node util/obj_pair_hashtable.h.node util/obj_triple_hashtable.h.node smt/smt_clause.h.node
	@echo done > smt/dyn_ack.h.node
smt/watch_list.h.node: ../src/smt/watch_list.h smt/smt_clause.h.node util/memory_manager.h.node
	@echo done > smt/watch_list.h.node
smt/smt_conflict_resolution.h.node: ../src/smt/smt_conflict_resolution.h smt/smt_literal.h.node smt/smt_bool_var_data.h.node smt/smt_justification.h.node smt/smt_enode.h.node smt/dyn_ack.h.node util/obj_pair_hashtable.h.node smt/params/smt_params.h.node util/obj_pair_hashtable.h.node util/map.h.node smt/watch_list.h.node util/obj_pair_set.h.node
	@echo done > smt/smt_conflict_resolution.h.node
smt/smt_relevancy.h.node: ../src/smt/smt_relevancy.h ast/ast.h.node
	@echo done > smt/smt_relevancy.h.node
smt/smt_induction.h.node: ../src/smt/smt_induction.h util/hash.h.node util/hashtable.h.node smt/smt_types.h.node ast/rewriter/value_sweep.h.node ast/datatype_decl_plugin.h.node ast/arith_decl_plugin.h.node
	@echo done > smt/smt_induction.h.node
smt/smt_case_split_queue.h.node: ../src/smt/smt_case_split_queue.h smt/smt_types.h.node util/heap.h.node smt/params/smt_params.h.node
	@echo done > smt/smt_case_split_queue.h.node
smt/smt_almost_cg_table.h.node: ../src/smt/smt_almost_cg_table.h smt/smt_enode.h.node util/map.h.node
	@echo done > smt/smt_almost_cg_table.h.node
smt/smt_failure.h.node: ../src/smt/smt_failure.h
	@echo done > smt/smt_failure.h.node
smt/asserted_formulas.h.node: ../src/smt/asserted_formulas.h util/statistics.h.node ast/static_features.h.node ast/expr_substitution.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/bit2int.h.node ast/rewriter/maximize_ac_sharing.h.node ast/rewriter/distribute_forall.h.node ast/rewriter/push_app_ite.h.node ast/rewriter/inj_axiom.h.node ast/rewriter/bv_elim.h.node ast/rewriter/der.h.node ast/rewriter/elim_bounds.h.node ast/macros/macro_manager.h.node ast/macros/macro_finder.h.node ast/normal_forms/defined_names.h.node ast/normal_forms/pull_quant.h.node ast/normal_forms/elim_term_ite.h.node ast/pattern/pattern_inference.h.node smt/params/smt_params.h.node
	@echo done > smt/asserted_formulas.h.node
smt/fingerprints.h.node: ../src/smt/fingerprints.h smt/smt_enode.h.node util/util.h.node
	@echo done > smt/fingerprints.h.node
smt/user_propagator.h.node: ../src/smt/user_propagator.h util/uint_set.h.node smt/smt_theory.h.node solver/solver.h.node
	@echo done > smt/user_propagator.h.node
smt/smt_context.h.node: ../src/smt/smt_context.h smt/smt_clause.h.node smt/smt_setup.h.node smt/smt_enode.h.node smt/smt_cg_table.h.node smt/smt_b_justification.h.node smt/smt_eq_justification.h.node smt/smt_justification.h.node smt/smt_bool_var_data.h.node smt/smt_clause_proof.h.node smt/smt_theory.h.node smt/smt_quantifier.h.node smt/smt_quantifier_stat.h.node smt/smt_statistics.h.node smt/smt_conflict_resolution.h.node smt/smt_relevancy.h.node smt/smt_induction.h.node smt/smt_case_split_queue.h.node smt/smt_almost_cg_table.h.node smt/smt_failure.h.node smt/asserted_formulas.h.node smt/smt_types.h.node smt/dyn_ack.h.node ast/ast_smt_pp.h.node smt/watch_list.h.node util/trail.h.node util/ref.h.node util/timer.h.node util/statistics.h.node smt/fingerprints.h.node smt/proto_model/proto_model.h.node smt/user_propagator.h.node model/model.h.node solver/progress_callback.h.node
	@echo done > smt/smt_context.h.node
smt/arith_eq_adapter.h.node: ../src/smt/arith_eq_adapter.h smt/smt_theory.h.node util/obj_pair_hashtable.h.node ast/arith_decl_plugin.h.node util/statistics.h.node
	@echo done > smt/arith_eq_adapter.h.node
smt/old_interval.h.node: ../src/smt/old_interval.h util/rational.h.node util/dependency.h.node
	@echo done > smt/old_interval.h.node
smt/arith_eq_solver.h.node: ../src/smt/arith_eq_solver.h ast/arith_decl_plugin.h.node ast/rewriter/arith_rewriter.h.node
	@echo done > smt/arith_eq_solver.h.node
smt/theory_opt.h.node: ../src/smt/theory_opt.h util/inf_rational.h.node util/inf_eps_rational.h.node ast/arith_decl_plugin.h.node
	@echo done > smt/theory_opt.h.node
smt/theory_arith.h.node: ../src/smt/theory_arith.h util/map.h.node util/heap.h.node util/nat_set.h.node util/inf_rational.h.node util/s_integer.h.node util/inf_s_integer.h.node util/obj_pair_hashtable.h.node util/uint_set.h.node ast/arith_decl_plugin.h.node model/numeral_factory.h.node smt/smt_theory.h.node smt/params/theory_arith_params.h.node smt/arith_eq_adapter.h.node smt/smt_context.h.node smt/old_interval.h.node smt/arith_eq_solver.h.node smt/theory_opt.h.node math/grobner/grobner.h.node
	@echo done > smt/theory_arith.h.node
smt/diff_logic.h.node: ../src/smt/diff_logic.h util/vector.h.node util/heap.h.node util/statistics.h.node util/trace.h.node util/warning.h.node util/uint_set.h.node
	@echo done > smt/diff_logic.h.node
smt/smt_model_generator.h.node: ../src/smt/smt_model_generator.h ast/ast.h.node smt/smt_types.h.node util/obj_hashtable.h.node util/map.h.node util/ref.h.node
	@echo done > smt/smt_model_generator.h.node
smt/theory_diff_logic.h.node: ../src/smt/theory_diff_logic.h util/rational.h.node util/inf_rational.h.node util/inf_int_rational.h.node util/s_integer.h.node util/inf_s_integer.h.node util/map.h.node ast/arith_decl_plugin.h.node model/numeral_factory.h.node smt/smt_theory.h.node smt/diff_logic.h.node smt/smt_justification.h.node smt/params/smt_params.h.node smt/arith_eq_adapter.h.node smt/smt_model_generator.h.node smt/smt_clause.h.node smt/theory_opt.h.node math/simplex/simplex.h.node math/simplex/simplex_def.h.node
	@echo done > smt/theory_diff_logic.h.node
smt/theory_special_relations.h.node: ../src/smt/theory_special_relations.h ast/special_relations_decl_plugin.h.node smt/smt_theory.h.node smt/theory_diff_logic.h.node util/union_find.h.node util/rational.h.node
	@echo done > smt/theory_special_relations.h.node
smt/theory_special_relations$(OBJ_EXT): ../src/smt/theory_special_relations.cpp ast/reg_decl_plugins.h.node ast/datatype_decl_plugin.h.node ast/recfun_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/recfun_replace.h.node smt/smt_context.h.node smt/theory_arith.h.node smt/theory_special_relations.h.node
	@echo src/smt/theory_special_relations.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_special_relations$(OBJ_EXT) ../src/smt/theory_special_relations.cpp
smt/theory_array_bapa.h.node: ../src/smt/theory_array_bapa.h ast/ast.h.node smt/smt_theory.h.node
	@echo done > smt/theory_array_bapa.h.node
smt/theory_array_base.h.node: ../src/smt/theory_array_base.h smt/smt_theory.h.node smt/theory_array_bapa.h.node ast/array_decl_plugin.h.node model/array_factory.h.node
	@echo done > smt/theory_array_base.h.node
smt/theory_array.h.node: ../src/smt/theory_array.h smt/theory_array_base.h.node smt/params/theory_array_params.h.node util/union_find.h.node
	@echo done > smt/theory_array.h.node
smt/theory_array_full.h.node: ../src/smt/theory_array_full.h smt/theory_array.h.node ast/ast_trail.h.node
	@echo done > smt/theory_array_full.h.node
smt/theory_array_full$(OBJ_EXT): ../src/smt/theory_array_full.cpp util/stats.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/ast_smt2_pp.h.node smt/smt_context.h.node smt/theory_array_full.h.node
	@echo src/smt/theory_array_full.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_array_full$(OBJ_EXT) ../src/smt/theory_array_full.cpp
util/scoped_vector.h.node: ../src/util/scoped_vector.h util/vector.h.node
	@echo done > util/scoped_vector.h.node
util/obj_ref_hashtable.h.node: ../src/util/obj_ref_hashtable.h util/obj_hashtable.h.node
	@echo done > util/obj_ref_hashtable.h.node
smt/theory_lra.h.node: ../src/smt/theory_lra.h smt/theory_opt.h.node
	@echo done > smt/theory_lra.h.node
smt/theory_bv.h.node: ../src/smt/theory_bv.h ast/rewriter/bit_blaster/bit_blaster.h.node util/trail.h.node util/union_find.h.node ast/arith_decl_plugin.h.node model/numeral_factory.h.node smt/smt_theory.h.node smt/params/theory_bv_params.h.node
	@echo done > smt/theory_bv.h.node
smt/smt_arith_value.h.node: ../src/smt/smt_arith_value.h ast/arith_decl_plugin.h.node smt/smt_context.h.node smt/theory_lra.h.node smt/theory_arith.h.node smt/theory_bv.h.node
	@echo done > smt/smt_arith_value.h.node
smt/theory_seq_empty.h.node: ../src/smt/theory_seq_empty.h ast/seq_decl_plugin.h.node model/seq_factory.h.node smt/smt_theory.h.node smt/smt_model_generator.h.node
	@echo done > smt/theory_seq_empty.h.node
smt/seq_skolem.h.node: ../src/smt/seq_skolem.h ast/seq_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/rewriter/th_rewriter.h.node
	@echo done > smt/seq_skolem.h.node
smt/seq_axioms.h.node: ../src/smt/seq_axioms.h ast/seq_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/rewriter/th_rewriter.h.node smt/smt_theory.h.node smt/seq_skolem.h.node
	@echo done > smt/seq_axioms.h.node
smt/seq_unicode.h.node: ../src/smt/seq_unicode.h util/s_integer.h.node ast/seq_decl_plugin.h.node smt/smt_theory.h.node smt/diff_logic.h.node
	@echo done > smt/seq_unicode.h.node
smt/seq_regex.h.node: ../src/smt/seq_regex.h util/scoped_vector.h.node util/state_graph.h.node ast/seq_decl_plugin.h.node ast/rewriter/seq_rewriter.h.node smt/smt_context.h.node smt/seq_skolem.h.node
	@echo done > smt/seq_regex.h.node
smt/seq_offset_eq.h.node: ../src/smt/seq_offset_eq.h util/obj_pair_hashtable.h.node ast/seq_decl_plugin.h.node ast/arith_decl_plugin.h.node smt/smt_theory.h.node
	@echo done > smt/seq_offset_eq.h.node
smt/theory_seq.h.node: ../src/smt/theory_seq.h ast/seq_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/ast_trail.h.node util/scoped_vector.h.node util/scoped_ptr_vector.h.node ast/rewriter/seq_rewriter.h.node util/union_find.h.node util/obj_ref_hashtable.h.node smt/smt_theory.h.node smt/smt_arith_value.h.node smt/theory_seq_empty.h.node smt/seq_skolem.h.node smt/seq_axioms.h.node smt/seq_unicode.h.node smt/seq_regex.h.node smt/seq_offset_eq.h.node
	@echo done > smt/theory_seq.h.node
smt/seq_ne_solver$(OBJ_EXT): ../src/smt/seq_ne_solver.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_trail.h.node ast/for_each_expr.h.node smt/smt_context.h.node smt/theory_seq.h.node smt/theory_arith.h.node
	@echo src/smt/seq_ne_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_ne_solver$(OBJ_EXT) ../src/smt/seq_ne_solver.cpp
smt/smt_farkas_util.h.node: ../src/smt/smt_farkas_util.h ast/arith_decl_plugin.h.node
	@echo done > smt/smt_farkas_util.h.node
smt/smt_farkas_util$(OBJ_EXT): ../src/smt/smt_farkas_util.cpp smt/smt_farkas_util.h.node ast/ast_pp.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/smt/smt_farkas_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_farkas_util$(OBJ_EXT) ../src/smt/smt_farkas_util.cpp
smt/smt_justification$(OBJ_EXT): ../src/smt/smt_justification.cpp smt/smt_context.h.node smt/smt_conflict_resolution.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_justification.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_justification$(OBJ_EXT) ../src/smt/smt_justification.cpp
smt/theory_dense_diff_logic.h.node: ../src/smt/theory_dense_diff_logic.h smt/theory_arith.h.node smt/params/theory_arith_params.h.node ast/arith_decl_plugin.h.node smt/arith_eq_adapter.h.node smt/theory_opt.h.node
	@echo done > smt/theory_dense_diff_logic.h.node
smt/theory_dense_diff_logic_def.h.node: ../src/smt/theory_dense_diff_logic_def.h smt/smt_context.h.node smt/theory_dense_diff_logic.h.node ast/ast_pp.h.node smt/smt_model_generator.h.node math/simplex/simplex.h.node math/simplex/simplex_def.h.node
	@echo done > smt/theory_dense_diff_logic_def.h.node
smt/theory_dense_diff_logic$(OBJ_EXT): ../src/smt/theory_dense_diff_logic.cpp smt/theory_dense_diff_logic_def.h.node
	@echo src/smt/theory_dense_diff_logic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_dense_diff_logic$(OBJ_EXT) ../src/smt/theory_dense_diff_logic.cpp
smt/watch_list$(OBJ_EXT): ../src/smt/watch_list.cpp smt/watch_list.h.node
	@echo src/smt/watch_list.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/watch_list$(OBJ_EXT) ../src/smt/watch_list.cpp
smt/smt_quantifier_stat$(OBJ_EXT): ../src/smt/smt_quantifier_stat.cpp smt/smt_quantifier_stat.h.node
	@echo src/smt/smt_quantifier_stat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_quantifier_stat$(OBJ_EXT) ../src/smt/smt_quantifier_stat.cpp
smt/smt_for_each_relevant_expr.h.node: ../src/smt/smt_for_each_relevant_expr.h ast/ast.h.node util/obj_hashtable.h.node util/vector.h.node
	@echo done > smt/smt_for_each_relevant_expr.h.node
smt/smt_for_each_relevant_expr$(OBJ_EXT): ../src/smt/smt_for_each_relevant_expr.cpp smt/smt_context.h.node smt/smt_for_each_relevant_expr.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_for_each_relevant_expr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_for_each_relevant_expr$(OBJ_EXT) ../src/smt/smt_for_each_relevant_expr.cpp
smt/smt_conflict_resolution$(OBJ_EXT): ../src/smt/smt_conflict_resolution.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node smt/smt_context.h.node smt/smt_conflict_resolution.h.node
	@echo src/smt/smt_conflict_resolution.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_conflict_resolution$(OBJ_EXT) ../src/smt/smt_conflict_resolution.cpp
util/pool.h.node: ../src/util/pool.h util/util.h.node util/vector.h.node
	@echo done > util/pool.h.node
smt/mam.h.node: ../src/smt/mam.h ast/ast.h.node smt/smt_types.h.node
	@echo done > smt/mam.h.node
smt/mam$(OBJ_EXT): ../src/smt/mam.cpp util/pool.h.node util/trail.h.node util/stopwatch.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node smt/mam.h.node smt/smt_context.h.node
	@echo src/smt/mam.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/mam$(OBJ_EXT) ../src/smt/mam.cpp
smt/smt_induction$(OBJ_EXT): ../src/smt/smt_induction.cpp ast/ast_pp.h.node ast/ast_util.h.node ast/for_each_expr.h.node ast/recfun_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/rewriter/value_sweep.h.node ast/rewriter/expr_safe_replace.h.node smt/smt_context.h.node smt/smt_induction.h.node
	@echo src/smt/smt_induction.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_induction$(OBJ_EXT) ../src/smt/smt_induction.cpp
smt/smt_quick_checker.h.node: ../src/smt/smt_quick_checker.h ast/ast.h.node ast/rewriter/th_rewriter.h.node util/obj_hashtable.h.node
	@echo done > smt/smt_quick_checker.h.node
smt/smt_quick_checker$(OBJ_EXT): ../src/smt/smt_quick_checker.cpp smt/smt_context.h.node smt/smt_quick_checker.h.node ast/ast_pp.h.node
	@echo src/smt/smt_quick_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_quick_checker$(OBJ_EXT) ../src/smt/smt_quick_checker.cpp
smt/smt_relevancy$(OBJ_EXT): ../src/smt/smt_relevancy.cpp smt/smt_context.h.node smt/smt_relevancy.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/smt/smt_relevancy.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_relevancy$(OBJ_EXT) ../src/smt/smt_relevancy.cpp
smt/seq_axioms$(OBJ_EXT): ../src/smt/seq_axioms.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node smt/seq_axioms.h.node smt/smt_context.h.node
	@echo src/smt/seq_axioms.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_axioms$(OBJ_EXT) ../src/smt/seq_axioms.cpp
smt/seq_skolem$(OBJ_EXT): ../src/smt/seq_skolem.cpp smt/seq_skolem.h.node ast/ast_pp.h.node
	@echo src/smt/seq_skolem.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_skolem$(OBJ_EXT) ../src/smt/seq_skolem.cpp
smt/theory_utvpi.h.node: ../src/smt/theory_utvpi.h smt/theory_diff_logic.h.node
	@echo done > smt/theory_utvpi.h.node
smt/theory_utvpi_def.h.node: ../src/smt/theory_utvpi_def.h smt/theory_utvpi.h.node util/heap.h.node ast/ast_pp.h.node smt/smt_context.h.node
	@echo done > smt/theory_utvpi_def.h.node
smt/theory_utvpi$(OBJ_EXT): ../src/smt/theory_utvpi.cpp smt/theory_utvpi.h.node smt/theory_utvpi_def.h.node
	@echo src/smt/theory_utvpi.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_utvpi$(OBJ_EXT) ../src/smt/theory_utvpi.cpp
smt/seq_offset_eq$(OBJ_EXT): ../src/smt/seq_offset_eq.cpp ast/ast_pp.h.node smt/seq_offset_eq.h.node smt/smt_context.h.node
	@echo src/smt/seq_offset_eq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_offset_eq$(OBJ_EXT) ../src/smt/seq_offset_eq.cpp
smt/seq_regex$(OBJ_EXT): ../src/smt/seq_regex.cpp smt/seq_regex.h.node smt/theory_seq.h.node ast/expr_abstract.h.node
	@echo src/smt/seq_regex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_regex$(OBJ_EXT) ../src/smt/seq_regex.cpp
smt/theory_array_bapa$(OBJ_EXT): ../src/smt/theory_array_bapa.cpp ast/ast_util.h.node ast/ast_pp.h.node ast/rewriter/array_rewriter.h.node smt/smt_context.h.node smt/smt_arith_value.h.node smt/theory_array_full.h.node smt/theory_array_bapa.h.node
	@echo src/smt/theory_array_bapa.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_array_bapa$(OBJ_EXT) ../src/smt/theory_array_bapa.cpp
smt/smt_almost_cg_table$(OBJ_EXT): ../src/smt/smt_almost_cg_table.cpp smt/smt_almost_cg_table.h.node
	@echo src/smt/smt_almost_cg_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_almost_cg_table$(OBJ_EXT) ../src/smt/smt_almost_cg_table.cpp
smt/cached_var_subst.h.node: ../src/smt/cached_var_subst.h ast/rewriter/var_subst.h.node util/map.h.node smt/smt_enode.h.node
	@echo done > smt/cached_var_subst.h.node
smt/cached_var_subst$(OBJ_EXT): ../src/smt/cached_var_subst.cpp smt/cached_var_subst.h.node ast/rewriter/rewriter_def.h.node
	@echo src/smt/cached_var_subst.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/cached_var_subst$(OBJ_EXT) ../src/smt/cached_var_subst.cpp
smt/uses_theory.h.node: ../src/smt/uses_theory.h ast/ast.h.node
	@echo done > smt/uses_theory.h.node
smt/uses_theory$(OBJ_EXT): ../src/smt/uses_theory.cpp smt/uses_theory.h.node ast/for_each_expr.h.node
	@echo src/smt/uses_theory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/uses_theory$(OBJ_EXT) ../src/smt/uses_theory.cpp
smt/smt_value_sort.h.node: ../src/smt/smt_value_sort.h ast/ast.h.node
	@echo done > smt/smt_value_sort.h.node
smt/smt_value_sort$(OBJ_EXT): ../src/smt/smt_value_sort.cpp smt/smt_value_sort.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/datatype_decl_plugin.h.node
	@echo src/smt/smt_value_sort.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_value_sort$(OBJ_EXT) ../src/smt/smt_value_sort.cpp
smt/theory_fpa.h.node: ../src/smt/theory_fpa.h smt/smt_theory.h.node util/trail.h.node ast/fpa/fpa2bv_converter.h.node ast/fpa/fpa2bv_rewriter.h.node ast/rewriter/th_rewriter.h.node model/fpa_factory.h.node smt/smt_model_generator.h.node
	@echo done > smt/theory_fpa.h.node
smt/theory_fpa$(OBJ_EXT): ../src/smt/theory_fpa.cpp ast/ast_smt2_pp.h.node smt/smt_context.h.node smt/theory_fpa.h.node smt/theory_bv.h.node smt/smt_model_generator.h.node ast/fpa/bv2fpa_converter.h.node
	@echo src/smt/theory_fpa.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_fpa$(OBJ_EXT) ../src/smt/theory_fpa.cpp
smt/smt_statistics$(OBJ_EXT): ../src/smt/smt_statistics.cpp smt/smt_statistics.h.node
	@echo src/smt/smt_statistics.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_statistics$(OBJ_EXT) ../src/smt/smt_statistics.cpp
smt/seq_eq_solver$(OBJ_EXT): ../src/smt/seq_eq_solver.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_trail.h.node ast/for_each_expr.h.node smt/smt_context.h.node smt/theory_seq.h.node smt/theory_arith.h.node
	@echo src/smt/seq_eq_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_eq_solver$(OBJ_EXT) ../src/smt/seq_eq_solver.cpp
smt/smt_kernel.h.node: ../src/smt/smt_kernel.h util/params.h.node util/lbool.h.node util/statistics.h.node ast/ast.h.node model/model.h.node solver/solver.h.node smt/smt_failure.h.node
	@echo done > smt/smt_kernel.h.node
smt/theory_str.h.node: ../src/smt/theory_str.h util/trail.h.node util/union_find.h.node util/scoped_ptr_vector.h.node util/hashtable.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/seq_rewriter.h.node ast/seq_decl_plugin.h.node model/value_factory.h.node smt/smt_theory.h.node smt/params/theory_str_params.h.node smt/smt_model_generator.h.node smt/smt_arith_value.h.node smt/smt_kernel.h.node
	@echo done > smt/theory_str.h.node
smt/theory_str$(OBJ_EXT): ../src/smt/theory_str.cpp ast/ast_smt2_pp.h.node smt/smt_context.h.node smt/theory_str.h.node smt/smt_model_generator.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node smt/theory_seq_empty.h.node smt/theory_arith.h.node ast/ast_util.h.node ast/rewriter/seq_rewriter.h.node ast/rewriter/expr_replacer.h.node ast/rewriter/var_subst.h.node smt/smt_kernel.h.node model/model_smt2_pp.h.node
	@echo src/smt/theory_str.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_str$(OBJ_EXT) ../src/smt/theory_str.cpp
smt/expr_context_simplifier.h.node: ../src/smt/expr_context_simplifier.h ast/ast.h.node util/obj_hashtable.h.node smt/params/smt_params.h.node smt/smt_kernel.h.node ast/arith_decl_plugin.h.node ast/rewriter/bool_rewriter.h.node
	@echo done > smt/expr_context_simplifier.h.node
smt/expr_context_simplifier$(OBJ_EXT): ../src/smt/expr_context_simplifier.cpp smt/expr_context_simplifier.h.node ast/ast_pp.h.node util/obj_hashtable.h.node smt/smt_kernel.h.node ast/for_each_expr.h.node
	@echo src/smt/expr_context_simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/expr_context_simplifier$(OBJ_EXT) ../src/smt/expr_context_simplifier.cpp
smt/smt_consequences$(OBJ_EXT): ../src/smt/smt_consequences.cpp util/max_cliques.h.node util/stopwatch.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/datatype_decl_plugin.h.node model/model_pp.h.node smt/smt_context.h.node
	@echo src/smt/smt_consequences.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_consequences$(OBJ_EXT) ../src/smt/smt_consequences.cpp
smt/smt_cg_table$(OBJ_EXT): ../src/smt/smt_cg_table.cpp smt/smt_cg_table.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_cg_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_cg_table$(OBJ_EXT) ../src/smt/smt_cg_table.cpp
smt/theory_datatype.h.node: ../src/smt/theory_datatype.h util/union_find.h.node ast/array_decl_plugin.h.node ast/datatype_decl_plugin.h.node model/datatype_factory.h.node smt/smt_theory.h.node smt/params/theory_datatype_params.h.node
	@echo done > smt/theory_datatype.h.node
smt/theory_datatype$(OBJ_EXT): ../src/smt/theory_datatype.cpp util/stats.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node smt/smt_context.h.node smt/theory_datatype.h.node smt/theory_array.h.node smt/smt_model_generator.h.node
	@echo src/smt/theory_datatype.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_datatype$(OBJ_EXT) ../src/smt/theory_datatype.cpp
smt/theory_str_mc$(OBJ_EXT): ../src/smt/theory_str_mc.cpp ast/ast_smt2_pp.h.node smt/smt_context.h.node smt/theory_str.h.node smt/smt_model_generator.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node smt/theory_seq_empty.h.node smt/theory_arith.h.node ast/ast_util.h.node ast/rewriter/seq_rewriter.h.node ast/rewriter/expr_replacer.h.node smt/smt_kernel.h.node model/model_smt2_pp.h.node
	@echo src/smt/theory_str_mc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_str_mc$(OBJ_EXT) ../src/smt/theory_str_mc.cpp
smt/smt_model_checker.h.node: ../src/smt/smt_model_checker.h util/obj_hashtable.h.node ast/ast.h.node ast/array_decl_plugin.h.node ast/normal_forms/defined_names.h.node smt/params/qi_params.h.node smt/params/smt_params.h.node
	@echo done > smt/smt_model_checker.h.node
smt/smt_model_finder.h.node: ../src/smt/smt_model_finder.h ast/ast.h.node ast/func_decl_dependencies.h.node model/model_macro_solver.h.node smt/proto_model/proto_model.h.node tactic/tactic_exception.h.node
	@echo done > smt/smt_model_finder.h.node
smt/smt_parallel.h.node: ../src/smt/smt_parallel.h smt/smt_context.h.node
	@echo done > smt/smt_parallel.h.node
smt/smt_context$(OBJ_EXT): ../src/smt/smt_context.cpp util/luby.h.node util/warning.h.node util/timeit.h.node util/union_find.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node ast/ast_translation.h.node ast/recfun_decl_plugin.h.node ast/proofs/proof_checker.h.node ast/ast_util.h.node ast/well_sorted.h.node model/model.h.node model/model_pp.h.node smt/smt_context.h.node smt/smt_quick_checker.h.node smt/uses_theory.h.node smt/smt_for_each_relevant_expr.h.node smt/smt_model_generator.h.node smt/smt_model_checker.h.node smt/smt_model_finder.h.node smt/smt_parallel.h.node smt/smt_arith_value.h.node
	@echo src/smt/smt_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_context$(OBJ_EXT) ../src/smt/smt_context.cpp
smt/smt_lookahead.h.node: ../src/smt/smt_lookahead.h ast/ast.h.node
	@echo done > smt/smt_lookahead.h.node
smt/smt_kernel$(OBJ_EXT): ../src/smt/smt_kernel.cpp smt/smt_kernel.h.node smt/smt_context.h.node smt/smt_lookahead.h.node ast/ast_smt2_pp.h.node smt/params/smt_params_helper.hpp.node
	@echo src/smt/smt_kernel.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_kernel$(OBJ_EXT) ../src/smt/smt_kernel.cpp
smt/smt_model_finder$(OBJ_EXT): ../src/smt/smt_model_finder.cpp util/backtrackable_set.h.node ast/ast_util.h.node ast/macros/macro_util.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/normal_forms/pull_quant.h.node ast/rewriter/var_subst.h.node ast/macros/cond_macro.h.node ast/macros/quantifier_macro_info.h.node ast/for_each_expr.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/well_sorted.h.node ast/ast_smt2_pp.h.node model/model_pp.h.node model/model_macro_solver.h.node smt/smt_model_finder.h.node smt/smt_context.h.node tactic/tactic_exception.h.node
	@echo src/smt/smt_model_finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_model_finder$(OBJ_EXT) ../src/smt/smt_model_finder.cpp
smt/smt_checker.h.node: ../src/smt/smt_checker.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > smt/smt_checker.h.node
smt/smt_checker$(OBJ_EXT): ../src/smt/smt_checker.cpp smt/smt_context.h.node smt/smt_checker.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_checker$(OBJ_EXT) ../src/smt/smt_checker.cpp
smt/smt_clause_proof$(OBJ_EXT): ../src/smt/smt_clause_proof.cpp smt/smt_clause_proof.h.node smt/smt_context.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_clause_proof.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_clause_proof$(OBJ_EXT) ../src/smt/smt_clause_proof.cpp
smt/theory_array$(OBJ_EXT): ../src/smt/theory_array.cpp smt/smt_context.h.node smt/theory_array.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node util/stats.h.node
	@echo src/smt/theory_array.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_array$(OBJ_EXT) ../src/smt/theory_array.cpp
smt/smt_parallel$(OBJ_EXT): ../src/smt/smt_parallel.cpp util/scoped_ptr_vector.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_translation.h.node smt/smt_parallel.h.node smt/smt_lookahead.h.node
	@echo src/smt/smt_parallel.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_parallel$(OBJ_EXT) ../src/smt/smt_parallel.cpp
smt/smt_solver$(OBJ_EXT): ../src/smt/smt_solver.cpp solver/solver_na2as.h.node smt/smt_kernel.h.node ast/reg_decl_plugins.h.node smt/params/smt_params.h.node smt/params/smt_params_helper.hpp.node solver/mus.h.node ast/for_each_expr.h.node ast/ast_smt2_pp.h.node ast/func_decl_dependencies.h.node util/dec_ref_util.h.node
	@echo src/smt/smt_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_solver$(OBJ_EXT) ../src/smt/smt_solver.cpp
smt/fingerprints$(OBJ_EXT): ../src/smt/fingerprints.cpp smt/fingerprints.h.node
	@echo src/smt/fingerprints.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/fingerprints$(OBJ_EXT) ../src/smt/fingerprints.cpp
smt/theory_opt$(OBJ_EXT): ../src/smt/theory_opt.cpp ast/arith_decl_plugin.h.node smt/smt_types.h.node smt/smt_theory.h.node smt/theory_opt.h.node
	@echo src/smt/theory_opt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_opt$(OBJ_EXT) ../src/smt/theory_opt.cpp
smt/cost_evaluator.h.node: ../src/smt/cost_evaluator.h ast/ast.h.node ast/arith_decl_plugin.h.node
	@echo done > smt/cost_evaluator.h.node
smt/qi_queue.h.node: ../src/smt/qi_queue.h ast/ast.h.node smt/smt_quantifier_stat.h.node smt/smt_checker.h.node smt/smt_quantifier.h.node smt/fingerprints.h.node smt/params/qi_params.h.node parsers/util/cost_parser.h.node smt/cost_evaluator.h.node smt/cached_var_subst.h.node util/statistics.h.node
	@echo done > smt/qi_queue.h.node
smt/qi_queue$(OBJ_EXT): ../src/smt/qi_queue.cpp util/warning.h.node util/stats.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/rewriter/var_subst.h.node smt/smt_context.h.node smt/qi_queue.h.node
	@echo src/smt/qi_queue.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/qi_queue$(OBJ_EXT) ../src/smt/qi_queue.cpp
smt/smt_solver.h.node: ../src/smt/smt_solver.h ast/ast.h.node util/params.h.node
	@echo done > smt/smt_solver.h.node
smt/smt_implied_equalities.h.node: ../src/smt/smt_implied_equalities.h smt/smt_solver.h.node util/lbool.h.node ast/ast.h.node
	@echo done > smt/smt_implied_equalities.h.node
smt/smt_implied_equalities$(OBJ_EXT): ../src/smt/smt_implied_equalities.cpp smt/smt_implied_equalities.h.node util/union_find.h.node ast/ast_pp.h.node ast/array_decl_plugin.h.node util/uint_set.h.node smt/smt_value_sort.h.node model/model_smt2_pp.h.node util/stopwatch.h.node model/model.h.node solver/solver.h.node
	@echo src/smt/smt_implied_equalities.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_implied_equalities$(OBJ_EXT) ../src/smt/smt_implied_equalities.cpp
smt/seq_unicode$(OBJ_EXT): ../src/smt/seq_unicode.cpp smt/seq_unicode.h.node smt/smt_context.h.node smt/smt_arith_value.h.node util/trail.h.node
	@echo src/smt/seq_unicode.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/seq_unicode$(OBJ_EXT) ../src/smt/seq_unicode.cpp
smt/theory_pb.h.node: ../src/smt/theory_pb.h smt/smt_theory.h.node ast/pb_decl_plugin.h.node smt/smt_clause.h.node smt/smt_b_justification.h.node smt/params/theory_pb_params.h.node math/simplex/simplex.h.node
	@echo done > smt/theory_pb.h.node
smt/theory_pb$(OBJ_EXT): ../src/smt/theory_pb.cpp smt/theory_pb.h.node smt/smt_context.h.node smt/smt_kernel.h.node ast/ast_pp.h.node util/sorting_network.h.node util/uint_set.h.node smt/smt_model_generator.h.node ast/rewriter/pb_rewriter_def.h.node math/simplex/sparse_matrix_def.h.node math/simplex/simplex_def.h.node
	@echo src/smt/theory_pb.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_pb$(OBJ_EXT) ../src/smt/theory_pb.cpp
smt/theory_array_base$(OBJ_EXT): ../src/smt/theory_array_base.cpp smt/smt_context.h.node smt/theory_array_base.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node smt/smt_model_generator.h.node model/func_interp.h.node ast/ast_smt2_pp.h.node
	@echo src/smt/theory_array_base.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_array_base$(OBJ_EXT) ../src/smt/theory_array_base.cpp
smt/smt_model_checker$(OBJ_EXT): ../src/smt/smt_model_checker.cpp ast/normal_forms/pull_quant.h.node ast/for_each_expr.h.node ast/rewriter/var_subst.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node ast/array_decl_plugin.h.node ast/ast_smt2_pp.h.node smt/smt_model_checker.h.node smt/smt_context.h.node smt/smt_model_finder.h.node model/model_pp.h.node
	@echo src/smt/smt_model_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_model_checker$(OBJ_EXT) ../src/smt/smt_model_checker.cpp
smt/smt_internalizer$(OBJ_EXT): ../src/smt/smt_internalizer.cpp smt/smt_context.h.node ast/expr_stat.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node smt/smt_model_finder.h.node ast/for_each_expr.h.node
	@echo src/smt/smt_internalizer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_internalizer$(OBJ_EXT) ../src/smt/smt_internalizer.cpp
smt/smt_context_stat$(OBJ_EXT): ../src/smt/smt_context_stat.cpp smt/smt_context.h.node ast/ast_pp.h.node
	@echo src/smt/smt_context_stat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_context_stat$(OBJ_EXT) ../src/smt/smt_context_stat.cpp
smt/smt_arith_value$(OBJ_EXT): ../src/smt/smt_arith_value.cpp ast/ast_pp.h.node smt/smt_arith_value.h.node
	@echo src/smt/smt_arith_value.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_arith_value$(OBJ_EXT) ../src/smt/smt_arith_value.cpp
smt/theory_dl$(OBJ_EXT): ../src/smt/theory_dl.cpp ast/dl_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/ast_pp.h.node model/value_factory.h.node smt/smt_theory.h.node smt/smt_model_generator.h.node smt/theory_bv.h.node smt/smt_context.h.node
	@echo src/smt/theory_dl.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_dl$(OBJ_EXT) ../src/smt/theory_dl.cpp
smt/theory_bv$(OBJ_EXT): ../src/smt/theory_bv.cpp smt/smt_context.h.node smt/theory_bv.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/bv_decl_plugin.h.node smt/smt_model_generator.h.node util/stats.h.node
	@echo src/smt/theory_bv.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_bv$(OBJ_EXT) ../src/smt/theory_bv.cpp
smt/theory_arith_core.h.node: ../src/smt/theory_arith_core.h ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node smt/smt_context.h.node smt/theory_arith.h.node smt/smt_model_generator.h.node
	@echo done > smt/theory_arith_core.h.node
smt/theory_arith_aux.h.node: ../src/smt/theory_arith_aux.h util/inf_eps_rational.h.node smt/theory_arith.h.node smt/smt_farkas_util.h.node ast/rewriter/th_rewriter.h.node tactic/generic_model_converter.h.node
	@echo done > smt/theory_arith_aux.h.node
smt/theory_arith_inv.h.node: ../src/smt/theory_arith_inv.h smt/theory_arith.h.node ast/ast_pp.h.node
	@echo done > smt/theory_arith_inv.h.node
smt/theory_arith_pp.h.node: ../src/smt/theory_arith_pp.h smt/theory_arith.h.node ast/ast_smt_pp.h.node util/stats.h.node
	@echo done > smt/theory_arith_pp.h.node
smt/theory_arith_int.h.node: ../src/smt/theory_arith_int.h util/numeral_buffer.h.node ast/ast_ll_pp.h.node ast/well_sorted.h.node ast/ast_smt2_pp.h.node
	@echo done > smt/theory_arith_int.h.node
smt/theory_arith_eq.h.node: ../src/smt/theory_arith_eq.h util/stopwatch.h.node
	@echo done > smt/theory_arith_eq.h.node
smt/theory_arith_nl.h.node: ../src/smt/theory_arith_nl.h ast/ast_smt2_pp.h.node
	@echo done > smt/theory_arith_nl.h.node
smt/theory_arith_def.h.node: ../src/smt/theory_arith_def.h smt/theory_arith.h.node smt/theory_arith_core.h.node smt/theory_arith_aux.h.node smt/theory_arith_inv.h.node smt/theory_arith_pp.h.node smt/theory_arith_int.h.node smt/theory_arith_eq.h.node smt/theory_arith_nl.h.node
	@echo done > smt/theory_arith_def.h.node
smt/theory_arith$(OBJ_EXT): ../src/smt/theory_arith.cpp smt/theory_arith_def.h.node
	@echo src/smt/theory_arith.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_arith$(OBJ_EXT) ../src/smt/theory_arith.cpp
smt/theory_recfun.h.node: ../src/smt/theory_recfun.h smt/smt_theory.h.node smt/smt_context.h.node ast/ast_pp.h.node ast/recfun_decl_plugin.h.node
	@echo done > smt/theory_recfun.h.node
smt/theory_recfun$(OBJ_EXT): ../src/smt/theory_recfun.cpp util/stats.h.node ast/ast_util.h.node ast/for_each_expr.h.node smt/theory_recfun.h.node
	@echo src/smt/theory_recfun.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_recfun$(OBJ_EXT) ../src/smt/theory_recfun.cpp
smt/arith_eq_adapter$(OBJ_EXT): ../src/smt/arith_eq_adapter.cpp smt/smt_context.h.node smt/arith_eq_adapter.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node util/stats.h.node ast/ast_smt2_pp.h.node
	@echo src/smt/arith_eq_adapter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/arith_eq_adapter$(OBJ_EXT) ../src/smt/arith_eq_adapter.cpp
smt/theory_dummy.h.node: ../src/smt/theory_dummy.h smt/smt_theory.h.node
	@echo done > smt/theory_dummy.h.node
smt/theory_dl.h.node: ../src/smt/theory_dl.h
	@echo done > smt/theory_dl.h.node
smt/smt_setup$(OBJ_EXT): ../src/smt/smt_setup.cpp smt/smt_context.h.node smt/smt_setup.h.node ast/static_features.h.node smt/theory_arith.h.node smt/theory_lra.h.node smt/theory_dense_diff_logic.h.node smt/theory_diff_logic.h.node smt/theory_utvpi.h.node smt/theory_array.h.node smt/theory_array_full.h.node smt/theory_bv.h.node smt/theory_datatype.h.node smt/theory_recfun.h.node smt/theory_dummy.h.node smt/theory_dl.h.node smt/theory_seq_empty.h.node smt/theory_seq.h.node smt/theory_special_relations.h.node smt/theory_pb.h.node smt/theory_fpa.h.node smt/theory_str.h.node
	@echo src/smt/smt_setup.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_setup$(OBJ_EXT) ../src/smt/smt_setup.cpp
smt/smt_enode$(OBJ_EXT): ../src/smt/smt_enode.cpp ast/ast_pp.h.node smt/smt_context.h.node smt/smt_enode.h.node
	@echo src/smt/smt_enode.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_enode$(OBJ_EXT) ../src/smt/smt_enode.cpp
smt/asserted_formulas$(OBJ_EXT): ../src/smt/asserted_formulas.cpp util/warning.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/for_each_expr.h.node ast/well_sorted.h.node ast/rewriter/rewriter_def.h.node ast/normal_forms/nnf.h.node ast/pattern/pattern_inference.h.node ast/macros/quasi_macros.h.node ast/occurs.h.node smt/asserted_formulas.h.node
	@echo src/smt/asserted_formulas.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/asserted_formulas$(OBJ_EXT) ../src/smt/asserted_formulas.cpp
smt/arith_eq_solver$(OBJ_EXT): ../src/smt/arith_eq_solver.cpp smt/arith_eq_solver.h.node
	@echo src/smt/arith_eq_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/arith_eq_solver$(OBJ_EXT) ../src/smt/arith_eq_solver.cpp
smt/theory_lra$(OBJ_EXT): ../src/smt/theory_lra.cpp util/stopwatch.h.node math/lp/lp_solver.h.node math/lp/lp_primal_simplex.h.node math/lp/lp_dual_simplex.h.node math/lp/indexed_value.h.node math/lp/lar_solver.h.node math/lp/nla_solver.h.node math/lp/lp_types.h.node math/lp/lp_api.h.node math/polynomial/algebraic_numbers.h.node math/polynomial/polynomial.h.node util/nat_set.h.node util/optional.h.node util/inf_rational.h.node util/cancel_eh.h.node util/scoped_timer.h.node util/nat_set.h.node ast/ast_pp.h.node model/numeral_factory.h.node smt/smt_theory.h.node smt/smt_context.h.node smt/theory_lra.h.node smt/smt_model_generator.h.node smt/arith_eq_adapter.h.node util/nat_set.h.node tactic/generic_model_converter.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node util/cancel_eh.h.node util/scoped_timer.h.node
	@echo src/smt/theory_lra.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_lra$(OBJ_EXT) ../src/smt/theory_lra.cpp
smt/theory_str_regex$(OBJ_EXT): ../src/smt/theory_str_regex.cpp smt/theory_str.h.node
	@echo src/smt/theory_str_regex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_str_regex$(OBJ_EXT) ../src/smt/theory_str_regex.cpp
smt/smt_context_inv$(OBJ_EXT): ../src/smt/smt_context_inv.cpp smt/smt_context.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node
	@echo src/smt/smt_context_inv.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_context_inv$(OBJ_EXT) ../src/smt/smt_context_inv.cpp
smt/user_propagator$(OBJ_EXT): ../src/smt/user_propagator.cpp ast/ast_pp.h.node smt/user_propagator.h.node smt/smt_context.h.node
	@echo src/smt/user_propagator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/user_propagator$(OBJ_EXT) ../src/smt/user_propagator.cpp
smt/theory_dummy$(OBJ_EXT): ../src/smt/theory_dummy.cpp smt/smt_context.h.node smt/theory_dummy.h.node
	@echo src/smt/theory_dummy.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_dummy$(OBJ_EXT) ../src/smt/theory_dummy.cpp
smt/smt_quantifier$(OBJ_EXT): ../src/smt/smt_quantifier.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node smt/smt_quantifier.h.node smt/smt_context.h.node smt/smt_quantifier_stat.h.node smt/smt_model_finder.h.node smt/smt_model_checker.h.node smt/smt_quick_checker.h.node smt/mam.h.node smt/qi_queue.h.node util/obj_hashtable.h.node
	@echo src/smt/smt_quantifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_quantifier$(OBJ_EXT) ../src/smt/smt_quantifier.cpp
smt/theory_diff_logic_def.h.node: ../src/smt/theory_diff_logic_def.h util/map.h.node util/warning.h.node ast/ast_pp.h.node smt/theory_diff_logic.h.node smt/smt_context.h.node smt/smt_model_generator.h.node model/model_implicant.h.node
	@echo done > smt/theory_diff_logic_def.h.node
smt/theory_diff_logic$(OBJ_EXT): ../src/smt/theory_diff_logic.cpp smt/theory_diff_logic.h.node util/rational.h.node smt/theory_diff_logic_def.h.node math/simplex/sparse_matrix_def.h.node
	@echo src/smt/theory_diff_logic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_diff_logic$(OBJ_EXT) ../src/smt/theory_diff_logic.cpp
smt/smt_context_pp$(OBJ_EXT): ../src/smt/smt_context_pp.cpp smt/smt_context.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_pp_util.h.node util/stats.h.node
	@echo src/smt/smt_context_pp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_context_pp$(OBJ_EXT) ../src/smt/smt_context_pp.cpp
smt/smt_model_generator$(OBJ_EXT): ../src/smt/smt_model_generator.cpp util/ref_util.h.node ast/for_each_expr.h.node ast/ast_pp.h.node ast/ast_smt2_pp.h.node ast/array_decl_plugin.h.node smt/smt_context.h.node smt/smt_model_generator.h.node smt/proto_model/proto_model.h.node model/model_v2_pp.h.node
	@echo src/smt/smt_model_generator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_model_generator$(OBJ_EXT) ../src/smt/smt_model_generator.cpp
smt/theory_seq$(OBJ_EXT): ../src/smt/theory_seq.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_trail.h.node ast/for_each_expr.h.node model/value_factory.h.node smt/smt_context.h.node smt/theory_seq.h.node smt/theory_arith.h.node smt/theory_lra.h.node smt/smt_kernel.h.node
	@echo src/smt/theory_seq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_seq$(OBJ_EXT) ../src/smt/theory_seq.cpp
smt/smt_clause$(OBJ_EXT): ../src/smt/smt_clause.cpp smt/smt_clause.h.node smt/smt_justification.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node
	@echo src/smt/smt_clause.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_clause$(OBJ_EXT) ../src/smt/smt_clause.cpp
smt/smt2_extra_cmds.h.node: ../src/smt/smt2_extra_cmds.h
	@echo done > smt/smt2_extra_cmds.h.node
smt/smt2_extra_cmds$(OBJ_EXT): ../src/smt/smt2_extra_cmds.cpp cmd_context/cmd_context.h.node parsers/smt2/smt2parser.h.node smt/smt2_extra_cmds.h.node
	@echo src/smt/smt2_extra_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt2_extra_cmds$(OBJ_EXT) ../src/smt/smt2_extra_cmds.cpp
smt/old_interval$(OBJ_EXT): ../src/smt/old_interval.cpp smt/old_interval.h.node
	@echo src/smt/old_interval.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/old_interval$(OBJ_EXT) ../src/smt/old_interval.cpp
smt/smt_literal$(OBJ_EXT): ../src/smt/smt_literal.cpp smt/smt_literal.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_literal.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_literal$(OBJ_EXT) ../src/smt/smt_literal.cpp
smt/smt_lookahead$(OBJ_EXT): ../src/smt/smt_lookahead.cpp ast/ast_pp.h.node ast/ast_ll_pp.h.node smt/smt_lookahead.h.node smt/smt_context.h.node
	@echo src/smt/smt_lookahead.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_lookahead$(OBJ_EXT) ../src/smt/smt_lookahead.cpp
smt/smt_case_split_queue$(OBJ_EXT): ../src/smt/smt_case_split_queue.cpp smt/smt_context.h.node smt/smt_case_split_queue.h.node util/warning.h.node util/stopwatch.h.node ast/for_each_expr.h.node ast/ast_pp.h.node util/map.h.node util/hashtable.h.node
	@echo src/smt/smt_case_split_queue.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_case_split_queue$(OBJ_EXT) ../src/smt/smt_case_split_queue.cpp
smt/theory_wmaxsat.h.node: ../src/smt/theory_wmaxsat.h smt/smt_theory.h.node smt/smt_clause.h.node tactic/generic_model_converter.h.node
	@echo done > smt/theory_wmaxsat.h.node
smt/theory_wmaxsat$(OBJ_EXT): ../src/smt/theory_wmaxsat.cpp smt/smt_context.h.node ast/ast_pp.h.node smt/theory_wmaxsat.h.node smt/smt_justification.h.node
	@echo src/smt/theory_wmaxsat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/theory_wmaxsat$(OBJ_EXT) ../src/smt/theory_wmaxsat.cpp
smt/dyn_ack$(OBJ_EXT): ../src/smt/dyn_ack.cpp smt/smt_context.h.node smt/dyn_ack.h.node ast/ast_pp.h.node
	@echo src/smt/dyn_ack.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/dyn_ack$(OBJ_EXT) ../src/smt/dyn_ack.cpp
smt/smt_theory$(OBJ_EXT): ../src/smt/smt_theory.cpp smt/smt_context.h.node util/buffer.h.node ast/ast_ll_pp.h.node
	@echo src/smt/smt_theory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/smt_theory$(OBJ_EXT) ../src/smt/smt_theory.cpp
smt/cost_evaluator$(OBJ_EXT): ../src/smt/cost_evaluator.cpp smt/cost_evaluator.h.node util/warning.h.node
	@echo src/smt/cost_evaluator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_41) $(CXX_OUT_FLAG)smt/cost_evaluator$(OBJ_EXT) ../src/smt/cost_evaluator.cpp
smt/smt$(LIB_EXT): smt/theory_special_relations$(OBJ_EXT) smt/theory_array_full$(OBJ_EXT) smt/seq_ne_solver$(OBJ_EXT) smt/smt_farkas_util$(OBJ_EXT) smt/smt_justification$(OBJ_EXT) smt/theory_dense_diff_logic$(OBJ_EXT) smt/watch_list$(OBJ_EXT) smt/smt_quantifier_stat$(OBJ_EXT) smt/smt_for_each_relevant_expr$(OBJ_EXT) smt/smt_conflict_resolution$(OBJ_EXT) smt/mam$(OBJ_EXT) smt/smt_induction$(OBJ_EXT) smt/smt_quick_checker$(OBJ_EXT) smt/smt_relevancy$(OBJ_EXT) smt/seq_axioms$(OBJ_EXT) smt/seq_skolem$(OBJ_EXT) smt/theory_utvpi$(OBJ_EXT) smt/seq_offset_eq$(OBJ_EXT) smt/seq_regex$(OBJ_EXT) smt/theory_array_bapa$(OBJ_EXT) smt/smt_almost_cg_table$(OBJ_EXT) smt/cached_var_subst$(OBJ_EXT) smt/uses_theory$(OBJ_EXT) smt/smt_value_sort$(OBJ_EXT) smt/theory_fpa$(OBJ_EXT) smt/smt_statistics$(OBJ_EXT) smt/seq_eq_solver$(OBJ_EXT) smt/theory_str$(OBJ_EXT) smt/expr_context_simplifier$(OBJ_EXT) smt/smt_consequences$(OBJ_EXT) smt/smt_cg_table$(OBJ_EXT) smt/theory_datatype$(OBJ_EXT) smt/theory_str_mc$(OBJ_EXT) smt/smt_context$(OBJ_EXT) smt/smt_kernel$(OBJ_EXT) smt/smt_model_finder$(OBJ_EXT) smt/smt_checker$(OBJ_EXT) smt/smt_clause_proof$(OBJ_EXT) smt/theory_array$(OBJ_EXT) smt/smt_parallel$(OBJ_EXT) smt/smt_solver$(OBJ_EXT) smt/fingerprints$(OBJ_EXT) smt/theory_opt$(OBJ_EXT) smt/qi_queue$(OBJ_EXT) smt/smt_implied_equalities$(OBJ_EXT) smt/seq_unicode$(OBJ_EXT) smt/theory_pb$(OBJ_EXT) smt/theory_array_base$(OBJ_EXT) smt/smt_model_checker$(OBJ_EXT) smt/smt_internalizer$(OBJ_EXT) smt/smt_context_stat$(OBJ_EXT) smt/smt_arith_value$(OBJ_EXT) smt/theory_dl$(OBJ_EXT) smt/theory_bv$(OBJ_EXT) smt/theory_arith$(OBJ_EXT) smt/theory_recfun$(OBJ_EXT) smt/arith_eq_adapter$(OBJ_EXT) smt/smt_setup$(OBJ_EXT) smt/smt_enode$(OBJ_EXT) smt/asserted_formulas$(OBJ_EXT) smt/arith_eq_solver$(OBJ_EXT) smt/theory_lra$(OBJ_EXT) smt/theory_str_regex$(OBJ_EXT) smt/smt_context_inv$(OBJ_EXT) smt/user_propagator$(OBJ_EXT) smt/theory_dummy$(OBJ_EXT) smt/smt_quantifier$(OBJ_EXT) smt/theory_diff_logic$(OBJ_EXT) smt/smt_context_pp$(OBJ_EXT) smt/smt_model_generator$(OBJ_EXT) smt/theory_seq$(OBJ_EXT) smt/smt_clause$(OBJ_EXT) smt/smt2_extra_cmds$(OBJ_EXT) smt/old_interval$(OBJ_EXT) smt/smt_literal$(OBJ_EXT) smt/smt_lookahead$(OBJ_EXT) smt/smt_case_split_queue$(OBJ_EXT) smt/theory_wmaxsat$(OBJ_EXT) smt/dyn_ack$(OBJ_EXT) smt/smt_theory$(OBJ_EXT) smt/cost_evaluator$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)smt/smt$(LIB_EXT) smt/theory_special_relations$(OBJ_EXT) smt/theory_array_full$(OBJ_EXT) smt/seq_ne_solver$(OBJ_EXT) smt/smt_farkas_util$(OBJ_EXT) smt/smt_justification$(OBJ_EXT) smt/theory_dense_diff_logic$(OBJ_EXT) smt/watch_list$(OBJ_EXT) smt/smt_quantifier_stat$(OBJ_EXT) smt/smt_for_each_relevant_expr$(OBJ_EXT) smt/smt_conflict_resolution$(OBJ_EXT) smt/mam$(OBJ_EXT) smt/smt_induction$(OBJ_EXT) smt/smt_quick_checker$(OBJ_EXT) smt/smt_relevancy$(OBJ_EXT) smt/seq_axioms$(OBJ_EXT) smt/seq_skolem$(OBJ_EXT) smt/theory_utvpi$(OBJ_EXT) smt/seq_offset_eq$(OBJ_EXT) smt/seq_regex$(OBJ_EXT) smt/theory_array_bapa$(OBJ_EXT) smt/smt_almost_cg_table$(OBJ_EXT) smt/cached_var_subst$(OBJ_EXT) smt/uses_theory$(OBJ_EXT) smt/smt_value_sort$(OBJ_EXT) smt/theory_fpa$(OBJ_EXT) smt/smt_statistics$(OBJ_EXT) smt/seq_eq_solver$(OBJ_EXT) smt/theory_str$(OBJ_EXT) smt/expr_context_simplifier$(OBJ_EXT) smt/smt_consequences$(OBJ_EXT) smt/smt_cg_table$(OBJ_EXT) smt/theory_datatype$(OBJ_EXT) smt/theory_str_mc$(OBJ_EXT) smt/smt_context$(OBJ_EXT) smt/smt_kernel$(OBJ_EXT) smt/smt_model_finder$(OBJ_EXT) smt/smt_checker$(OBJ_EXT) smt/smt_clause_proof$(OBJ_EXT) smt/theory_array$(OBJ_EXT) smt/smt_parallel$(OBJ_EXT) smt/smt_solver$(OBJ_EXT) smt/fingerprints$(OBJ_EXT) smt/theory_opt$(OBJ_EXT) smt/qi_queue$(OBJ_EXT) smt/smt_implied_equalities$(OBJ_EXT) smt/seq_unicode$(OBJ_EXT) smt/theory_pb$(OBJ_EXT) smt/theory_array_base$(OBJ_EXT) smt/smt_model_checker$(OBJ_EXT) smt/smt_internalizer$(OBJ_EXT) smt/smt_context_stat$(OBJ_EXT) smt/smt_arith_value$(OBJ_EXT) smt/theory_dl$(OBJ_EXT) smt/theory_bv$(OBJ_EXT) smt/theory_arith$(OBJ_EXT) smt/theory_recfun$(OBJ_EXT) smt/arith_eq_adapter$(OBJ_EXT) smt/smt_setup$(OBJ_EXT) smt/smt_enode$(OBJ_EXT) smt/asserted_formulas$(OBJ_EXT) smt/arith_eq_solver$(OBJ_EXT) smt/theory_lra$(OBJ_EXT) smt/theory_str_regex$(OBJ_EXT) smt/smt_context_inv$(OBJ_EXT) smt/user_propagator$(OBJ_EXT) smt/theory_dummy$(OBJ_EXT) smt/smt_quantifier$(OBJ_EXT) smt/theory_diff_logic$(OBJ_EXT) smt/smt_context_pp$(OBJ_EXT) smt/smt_model_generator$(OBJ_EXT) smt/theory_seq$(OBJ_EXT) smt/smt_clause$(OBJ_EXT) smt/smt2_extra_cmds$(OBJ_EXT) smt/old_interval$(OBJ_EXT) smt/smt_literal$(OBJ_EXT) smt/smt_lookahead$(OBJ_EXT) smt/smt_case_split_queue$(OBJ_EXT) smt/theory_wmaxsat$(OBJ_EXT) smt/dyn_ack$(OBJ_EXT) smt/smt_theory$(OBJ_EXT) smt/cost_evaluator$(OBJ_EXT)
smt: smt/smt$(LIB_EXT)

includes_42 = -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/rewriter/bit_blaster -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src
tactic/bv/bv_bounds_tactic.h.node: ../src/tactic/bv/bv_bounds_tactic.h tactic/tactic.h.node
	@echo done > tactic/bv/bv_bounds_tactic.h.node
tactic/bv/bv_bounds_tactic$(OBJ_EXT): ../src/tactic/bv/bv_bounds_tactic.cpp tactic/bv/bv_bounds_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/core/dom_simplify_tactic.h.node ast/bv_decl_plugin.h.node ast/ast_pp.h.node
	@echo src/tactic/bv/bv_bounds_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bv_bounds_tactic$(OBJ_EXT) ../src/tactic/bv/bv_bounds_tactic.cpp
tactic/bv/bvarray2uf_tactic.h.node: ../src/tactic/bv/bvarray2uf_tactic.h util/params.h.node
	@echo done > tactic/bv/bvarray2uf_tactic.h.node
tactic/bv/bvarray2uf_rewriter.h.node: ../src/tactic/bv/bvarray2uf_rewriter.h ast/rewriter/rewriter.h.node tactic/generic_model_converter.h.node
	@echo done > tactic/bv/bvarray2uf_rewriter.h.node
tactic/bv/bvarray2uf_tactic$(OBJ_EXT): ../src/tactic/bv/bvarray2uf_tactic.cpp tactic/tactical.h.node ast/bv_decl_plugin.h.node ast/rewriter/expr_replacer.h.node tactic/generic_model_converter.h.node ast/ast_smt2_pp.h.node tactic/bv/bvarray2uf_tactic.h.node tactic/bv/bvarray2uf_rewriter.h.node
	@echo src/tactic/bv/bvarray2uf_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bvarray2uf_tactic$(OBJ_EXT) ../src/tactic/bv/bvarray2uf_tactic.cpp
tactic/bv/bit_blaster_model_converter.h.node: ../src/tactic/bv/bit_blaster_model_converter.h tactic/model_converter.h.node
	@echo done > tactic/bv/bit_blaster_model_converter.h.node
tactic/bv/bit_blaster_tactic$(OBJ_EXT): ../src/tactic/bv/bit_blaster_tactic.cpp tactic/tactical.h.node tactic/bv/bit_blaster_model_converter.h.node ast/rewriter/bit_blaster/bit_blaster_rewriter.h.node ast/ast_pp.h.node model/model_pp.h.node ast/rewriter/rewriter_types.h.node
	@echo src/tactic/bv/bit_blaster_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bit_blaster_tactic$(OBJ_EXT) ../src/tactic/bv/bit_blaster_tactic.cpp
tactic/bv/bit_blaster_model_converter$(OBJ_EXT): ../src/tactic/bv/bit_blaster_model_converter.cpp model/model.h.node model/model_pp.h.node tactic/model_converter.h.node ast/bv_decl_plugin.h.node ast/ast_smt2_pp.h.node ast/ast_util.h.node
	@echo src/tactic/bv/bit_blaster_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bit_blaster_model_converter$(OBJ_EXT) ../src/tactic/bv/bit_blaster_model_converter.cpp
tactic/bv/dt2bv_tactic.h.node: ../src/tactic/bv/dt2bv_tactic.h util/params.h.node util/obj_hashtable.h.node
	@echo done > tactic/bv/dt2bv_tactic.h.node
tactic/bv/dt2bv_tactic$(OBJ_EXT): ../src/tactic/bv/dt2bv_tactic.cpp tactic/bv/dt2bv_tactic.h.node tactic/tactical.h.node tactic/generic_model_converter.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node ast/ast_util.h.node ast/rewriter/enum2bv_rewriter.h.node ast/ast_pp.h.node
	@echo src/tactic/bv/dt2bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/dt2bv_tactic$(OBJ_EXT) ../src/tactic/bv/dt2bv_tactic.cpp
tactic/bv/elim_small_bv_tactic.h.node: ../src/tactic/bv/elim_small_bv_tactic.h util/params.h.node
	@echo done > tactic/bv/elim_small_bv_tactic.h.node
tactic/bv/elim_small_bv_tactic$(OBJ_EXT): ../src/tactic/bv/elim_small_bv_tactic.cpp tactic/tactical.h.node ast/rewriter/rewriter_def.h.node tactic/generic_model_converter.h.node ast/bv_decl_plugin.h.node ast/used_vars.h.node ast/well_sorted.h.node ast/rewriter/var_subst.h.node ast/rewriter/th_rewriter.h.node tactic/bv/elim_small_bv_tactic.h.node
	@echo src/tactic/bv/elim_small_bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/elim_small_bv_tactic$(OBJ_EXT) ../src/tactic/bv/elim_small_bv_tactic.cpp
tactic/bv/bv_bound_chk_tactic.h.node: ../src/tactic/bv/bv_bound_chk_tactic.h tactic/tactical.h.node util/params.h.node ast/ast.h.node
	@echo done > tactic/bv/bv_bound_chk_tactic.h.node
tactic/bv/bv_bound_chk_tactic$(OBJ_EXT): ../src/tactic/bv/bv_bound_chk_tactic.cpp tactic/bv/bv_bound_chk_tactic.h.node ast/ast.h.node params/rewriter_params.hpp.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/bv_bounds.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/tactic/bv/bv_bound_chk_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bv_bound_chk_tactic$(OBJ_EXT) ../src/tactic/bv/bv_bound_chk_tactic.cpp
tactic/bv/bv_size_reduction_tactic$(OBJ_EXT): ../src/tactic/bv/bv_size_reduction_tactic.cpp tactic/tactical.h.node ast/bv_decl_plugin.h.node ast/rewriter/expr_replacer.h.node tactic/generic_model_converter.h.node ast/ast_smt2_pp.h.node
	@echo src/tactic/bv/bv_size_reduction_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bv_size_reduction_tactic$(OBJ_EXT) ../src/tactic/bv/bv_size_reduction_tactic.cpp
tactic/bv/max_bv_sharing_tactic$(OBJ_EXT): ../src/tactic/bv/max_bv_sharing_tactic.cpp tactic/tactical.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter_def.h.node util/obj_pair_hashtable.h.node ast/ast_lt.h.node
	@echo src/tactic/bv/max_bv_sharing_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/max_bv_sharing_tactic$(OBJ_EXT) ../src/tactic/bv/max_bv_sharing_tactic.cpp
tactic/bv/bv1_blaster_tactic$(OBJ_EXT): ../src/tactic/bv/bv1_blaster_tactic.cpp tactic/tactical.h.node tactic/bv/bit_blaster_model_converter.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter_def.h.node ast/for_each_expr.h.node ast/ast_util.h.node ast/rewriter/bv_rewriter.h.node
	@echo src/tactic/bv/bv1_blaster_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bv1_blaster_tactic$(OBJ_EXT) ../src/tactic/bv/bv1_blaster_tactic.cpp
tactic/bv/bvarray2uf_rewriter$(OBJ_EXT): ../src/tactic/bv/bvarray2uf_rewriter.cpp ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node util/params.h.node ast/ast_pp.h.node tactic/bv/bvarray2uf_rewriter.h.node ast/rewriter/rewriter_def.h.node util/ref_util.h.node
	@echo src/tactic/bv/bvarray2uf_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_42) $(CXX_OUT_FLAG)tactic/bv/bvarray2uf_rewriter$(OBJ_EXT) ../src/tactic/bv/bvarray2uf_rewriter.cpp
tactic/bv/bv_tactics$(LIB_EXT): tactic/bv/bv_bounds_tactic$(OBJ_EXT) tactic/bv/bvarray2uf_tactic$(OBJ_EXT) tactic/bv/bit_blaster_tactic$(OBJ_EXT) tactic/bv/bit_blaster_model_converter$(OBJ_EXT) tactic/bv/dt2bv_tactic$(OBJ_EXT) tactic/bv/elim_small_bv_tactic$(OBJ_EXT) tactic/bv/bv_bound_chk_tactic$(OBJ_EXT) tactic/bv/bv_size_reduction_tactic$(OBJ_EXT) tactic/bv/max_bv_sharing_tactic$(OBJ_EXT) tactic/bv/bv1_blaster_tactic$(OBJ_EXT) tactic/bv/bvarray2uf_rewriter$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/bv/bv_tactics$(LIB_EXT) tactic/bv/bv_bounds_tactic$(OBJ_EXT) tactic/bv/bvarray2uf_tactic$(OBJ_EXT) tactic/bv/bit_blaster_tactic$(OBJ_EXT) tactic/bv/bit_blaster_model_converter$(OBJ_EXT) tactic/bv/dt2bv_tactic$(OBJ_EXT) tactic/bv/elim_small_bv_tactic$(OBJ_EXT) tactic/bv/bv_bound_chk_tactic$(OBJ_EXT) tactic/bv/bv_size_reduction_tactic$(OBJ_EXT) tactic/bv/max_bv_sharing_tactic$(OBJ_EXT) tactic/bv/bv1_blaster_tactic$(OBJ_EXT) tactic/bv/bvarray2uf_rewriter$(OBJ_EXT)
bv_tactics: tactic/bv/bv_tactics$(LIB_EXT)

includes_43 = -I../src/ast -I../src/util -I../src/math/polynomial -I../src
test/fuzzing/expr_delta.h.node: ../src/test/fuzzing/expr_delta.h ast/ast.h.node
	@echo done > test/fuzzing/expr_delta.h.node
test/fuzzing/expr_delta$(OBJ_EXT): ../src/test/fuzzing/expr_delta.cpp test/fuzzing/expr_delta.h.node ast/ast_pp.h.node
	@echo src/test/fuzzing/expr_delta.cpp
	@$(CXX) $(CXXFLAGS) $(includes_43) $(CXX_OUT_FLAG)test/fuzzing/expr_delta$(OBJ_EXT) ../src/test/fuzzing/expr_delta.cpp
test/fuzzing/expr_rand.h.node: ../src/test/fuzzing/expr_rand.h ast/ast.h.node util/obj_hashtable.h.node
	@echo done > test/fuzzing/expr_rand.h.node
test/fuzzing/expr_rand$(OBJ_EXT): ../src/test/fuzzing/expr_rand.cpp test/fuzzing/expr_rand.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node
	@echo src/test/fuzzing/expr_rand.cpp
	@$(CXX) $(CXXFLAGS) $(includes_43) $(CXX_OUT_FLAG)test/fuzzing/expr_rand$(OBJ_EXT) ../src/test/fuzzing/expr_rand.cpp
test/fuzzing/fuzzing$(LIB_EXT): test/fuzzing/expr_delta$(OBJ_EXT) test/fuzzing/expr_rand$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)test/fuzzing/fuzzing$(LIB_EXT) test/fuzzing/expr_delta$(OBJ_EXT) test/fuzzing/expr_rand$(OBJ_EXT)
fuzzing: test/fuzzing/fuzzing$(LIB_EXT)

includes_44 = -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src
smt/tactic/smt_tactic$(OBJ_EXT): ../src/smt/tactic/smt_tactic.cpp util/debug.h.node ast/rewriter/rewriter_types.h.node ast/ast_util.h.node ast/ast_ll_pp.h.node smt/smt_kernel.h.node smt/params/smt_params.h.node smt/params/smt_params_helper.hpp.node smt/smt_solver.h.node tactic/tactic.h.node tactic/tactical.h.node tactic/generic_model_converter.h.node solver/solver2tactic.h.node solver/solver.h.node solver/mus.h.node solver/parallel_tactic.h.node solver/parallel_params.hpp.node
	@echo src/smt/tactic/smt_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_44) $(CXX_OUT_FLAG)smt/tactic/smt_tactic$(OBJ_EXT) ../src/smt/tactic/smt_tactic.cpp
smt/tactic/ctx_solver_simplify_tactic.h.node: ../src/smt/tactic/ctx_solver_simplify_tactic.h tactic/tactical.h.node
	@echo done > smt/tactic/ctx_solver_simplify_tactic.h.node
smt/tactic/ctx_solver_simplify_tactic$(OBJ_EXT): ../src/smt/tactic/ctx_solver_simplify_tactic.cpp smt/tactic/ctx_solver_simplify_tactic.h.node ast/arith_decl_plugin.h.node smt/params/smt_params.h.node smt/smt_kernel.h.node ast/ast_pp.h.node ast/rewriter/mk_simplified_app.h.node ast/ast_util.h.node
	@echo src/smt/tactic/ctx_solver_simplify_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_44) $(CXX_OUT_FLAG)smt/tactic/ctx_solver_simplify_tactic$(OBJ_EXT) ../src/smt/tactic/ctx_solver_simplify_tactic.cpp
smt/tactic/unit_subsumption_tactic.h.node: ../src/smt/tactic/unit_subsumption_tactic.h tactic/tactic.h.node
	@echo done > smt/tactic/unit_subsumption_tactic.h.node
smt/tactic/unit_subsumption_tactic$(OBJ_EXT): ../src/smt/tactic/unit_subsumption_tactic.cpp ast/ast_util.h.node smt/tactic/unit_subsumption_tactic.h.node smt/smt_context.h.node
	@echo src/smt/tactic/unit_subsumption_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_44) $(CXX_OUT_FLAG)smt/tactic/unit_subsumption_tactic$(OBJ_EXT) ../src/smt/tactic/unit_subsumption_tactic.cpp
smt/tactic/smt_tactic$(LIB_EXT): smt/tactic/smt_tactic$(OBJ_EXT) smt/tactic/ctx_solver_simplify_tactic$(OBJ_EXT) smt/tactic/unit_subsumption_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)smt/tactic/smt_tactic$(LIB_EXT) smt/tactic/smt_tactic$(OBJ_EXT) smt/tactic/ctx_solver_simplify_tactic$(OBJ_EXT) smt/tactic/unit_subsumption_tactic$(OBJ_EXT)
smt_tactic: smt/tactic/smt_tactic$(LIB_EXT)

includes_45 = -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/tactic/core -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src
tactic/bv/bv_size_reduction_tactic.h.node: ../src/tactic/bv/bv_size_reduction_tactic.h util/params.h.node
	@echo done > tactic/bv/bv_size_reduction_tactic.h.node
tactic/bv/max_bv_sharing_tactic.h.node: ../src/tactic/bv/max_bv_sharing_tactic.h util/params.h.node
	@echo done > tactic/bv/max_bv_sharing_tactic.h.node
tactic/sls/sls_tactic.h.node: ../src/tactic/sls/sls_tactic.h util/params.h.node
	@echo done > tactic/sls/sls_tactic.h.node
tactic/sls/sls_params.hpp.node: ../src/tactic/sls/sls_params.hpp util/params.h.node util/gparams.h.node
	@echo done > tactic/sls/sls_params.hpp.node
tactic/sls/sls_powers.h.node: ../src/tactic/sls/sls_powers.h util/mpz.h.node
	@echo done > tactic/sls/sls_powers.h.node
tactic/sls/sls_tracker.h.node: ../src/tactic/sls/sls_tracker.h ast/for_each_expr.h.node ast/ast_smt2_pp.h.node ast/bv_decl_plugin.h.node model/model.h.node tactic/sls/sls_params.hpp.node tactic/sls/sls_powers.h.node
	@echo done > tactic/sls/sls_tracker.h.node
tactic/sls/sls_evaluator.h.node: ../src/tactic/sls/sls_evaluator.h model/model_evaluator.h.node tactic/sls/sls_powers.h.node tactic/sls/sls_tracker.h.node
	@echo done > tactic/sls/sls_evaluator.h.node
tactic/sls/sls_engine.h.node: ../src/tactic/sls/sls_engine.h util/stopwatch.h.node util/lbool.h.node tactic/model_converter.h.node tactic/goal.h.node tactic/sls/sls_tracker.h.node tactic/sls/sls_evaluator.h.node util/statistics.h.node
	@echo done > tactic/sls/sls_engine.h.node
tactic/sls/sls_tactic$(OBJ_EXT): ../src/tactic/sls/sls_tactic.cpp ast/normal_forms/nnf.h.node tactic/core/solve_eqs_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/core/nnf_tactic.h.node util/stopwatch.h.node tactic/sls/sls_tactic.h.node tactic/sls/sls_params.hpp.node tactic/sls/sls_engine.h.node
	@echo src/tactic/sls/sls_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_45) $(CXX_OUT_FLAG)tactic/sls/sls_tactic$(OBJ_EXT) ../src/tactic/sls/sls_tactic.cpp
tactic/sls/bvsls_opt_engine.h.node: ../src/tactic/sls/bvsls_opt_engine.h tactic/sls/sls_engine.h.node
	@echo done > tactic/sls/bvsls_opt_engine.h.node
tactic/sls/bvsls_opt_engine$(OBJ_EXT): ../src/tactic/sls/bvsls_opt_engine.cpp ast/normal_forms/nnf.h.node tactic/sls/bvsls_opt_engine.h.node
	@echo src/tactic/sls/bvsls_opt_engine.cpp
	@$(CXX) $(CXXFLAGS) $(includes_45) $(CXX_OUT_FLAG)tactic/sls/bvsls_opt_engine$(OBJ_EXT) ../src/tactic/sls/bvsls_opt_engine.cpp
tactic/sls/sls_engine$(OBJ_EXT): ../src/tactic/sls/sls_engine.cpp util/map.h.node ast/ast_smt2_pp.h.node ast/ast_pp.h.node ast/rewriter/var_subst.h.node model/model_pp.h.node tactic/tactic.h.node util/luby.h.node tactic/sls/sls_params.hpp.node tactic/sls/sls_engine.h.node
	@echo src/tactic/sls/sls_engine.cpp
	@$(CXX) $(CXXFLAGS) $(includes_45) $(CXX_OUT_FLAG)tactic/sls/sls_engine$(OBJ_EXT) ../src/tactic/sls/sls_engine.cpp
tactic/sls/sls_tactic$(LIB_EXT): tactic/sls/sls_tactic$(OBJ_EXT) tactic/sls/bvsls_opt_engine$(OBJ_EXT) tactic/sls/sls_engine$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/sls/sls_tactic$(LIB_EXT) tactic/sls/sls_tactic$(OBJ_EXT) tactic/sls/bvsls_opt_engine$(OBJ_EXT) tactic/sls/sls_engine$(OBJ_EXT)
sls_tactic: tactic/sls/sls_tactic$(LIB_EXT)

includes_46 = -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src
qe/qe.h.node: ../src/qe/qe.h ast/ast.h.node smt/params/smt_params.h.node util/statistics.h.node util/lbool.h.node ast/expr_functors.h.node ast/rewriter/rewriter.h.node model/model.h.node util/params.h.node
	@echo done > qe/qe.h.node
qe/qe_tactic$(OBJ_EXT): ../src/qe/qe_tactic.cpp tactic/tactical.h.node qe/qe.h.node
	@echo src/qe/qe_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_tactic$(OBJ_EXT) ../src/qe/qe_tactic.cpp
qe/qe_datatype_plugin$(OBJ_EXT): ../src/qe/qe_datatype_plugin.cpp qe/qe.h.node ast/datatype_decl_plugin.h.node ast/rewriter/expr_safe_replace.h.node util/obj_pair_hashtable.h.node ast/for_each_expr.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/qe/qe_datatype_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_datatype_plugin$(OBJ_EXT) ../src/qe/qe_datatype_plugin.cpp
qe/nlqsat.h.node: ../src/qe/nlqsat.h tactic/tactic.h.node
	@echo done > qe/nlqsat.h.node
qe/qe_mbp.h.node: ../src/qe/qe_mbp.h ast/ast.h.node util/params.h.node model/model.h.node math/simplex/model_based_opt.h.node
	@echo done > qe/qe_mbp.h.node
qe/qsat.h.node: ../src/qe/qsat.h tactic/tactic.h.node tactic/generic_model_converter.h.node qe/qe_mbp.h.node
	@echo done > qe/qsat.h.node
qe/nlqsat$(OBJ_EXT): ../src/qe/nlqsat.cpp util/uint_set.h.node util/scoped_ptr_vector.h.node ast/expr2var.h.node ast/ast_util.h.node ast/rewriter/expr_safe_replace.h.node ast/ast_pp.h.node ast/for_each_expr.h.node ast/rewriter/rewriter.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/quant_hoist.h.node qe/nlqsat.h.node qe/qsat.h.node nlsat/nlsat_solver.h.node nlsat/nlsat_explain.h.node nlsat/nlsat_assignment.h.node nlsat/tactic/goal2nlsat.h.node tactic/core/tseitin_cnf_tactic.h.node
	@echo src/qe/nlqsat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/nlqsat$(OBJ_EXT) ../src/qe/nlqsat.cpp
qe/qe_mbi.h.node: ../src/qe/qe_mbi.h qe/mbp/mbp_arith.h.node qe/mbp/mbp_plugin.h.node util/lbool.h.node
	@echo done > qe/qe_mbi.h.node
qe/qe_mbi$(OBJ_EXT): ../src/qe/qe_mbi.cpp ast/ast_util.h.node ast/ast_pp.h.node ast/for_each_expr.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node ast/arith_decl_plugin.h.node model/model_evaluator.h.node solver/solver.h.node qe/qe_mbi.h.node qe/mbp/mbp_term_graph.h.node qe/mbp/mbp_arith.h.node qe/mbp/mbp_arrays.h.node
	@echo src/qe/qe_mbi.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_mbi$(OBJ_EXT) ../src/qe/qe_mbi.cpp
qe/qe_lite.h.node: ../src/qe/qe_lite.h ast/ast.h.node util/uint_set.h.node util/params.h.node
	@echo done > qe/qe_lite.h.node
qe/qe_mbp$(OBJ_EXT): ../src/qe/qe_mbp.cpp ast/rewriter/expr_safe_replace.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/occurs.h.node ast/rewriter/th_rewriter.h.node ast/expr_functors.h.node ast/for_each_expr.h.node ast/scoped_proof.h.node qe/qe_mbp.h.node qe/mbp/mbp_arith.h.node qe/mbp/mbp_arrays.h.node qe/mbp/mbp_datatypes.h.node qe/qe_lite.h.node model/model_pp.h.node model/model_evaluator.h.node
	@echo src/qe/qe_mbp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_mbp$(OBJ_EXT) ../src/qe/qe_mbp.cpp
qe/nlarith_util.h.node: ../src/qe/nlarith_util.h ast/ast.h.node util/lbool.h.node
	@echo done > qe/nlarith_util.h.node
qe/qe_arith_plugin$(OBJ_EXT): ../src/qe/qe_arith_plugin.cpp qe/qe.h.node ast/ast_pp.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/bool_rewriter.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node smt/arith_eq_solver.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/factor_rewriter.h.node util/obj_pair_hashtable.h.node qe/nlarith_util.h.node model/model_evaluator.h.node smt/smt_kernel.h.node qe/mbp/mbp_arith.h.node
	@echo src/qe/qe_arith_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_arith_plugin$(OBJ_EXT) ../src/qe/qe_arith_plugin.cpp
qe/qe_bool_plugin$(OBJ_EXT): ../src/qe/qe_bool_plugin.cpp qe/qe.h.node ast/rewriter/expr_safe_replace.h.node ast/ast_pp.h.node model/model_evaluator.h.node
	@echo src/qe/qe_bool_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_bool_plugin$(OBJ_EXT) ../src/qe/qe_bool_plugin.cpp
qe/nlarith_util$(OBJ_EXT): ../src/qe/nlarith_util.cpp ast/ast.h.node qe/nlarith_util.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node qe/qe.h.node ast/rewriter/expr_replacer.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node ast/expr_functors.h.node
	@echo src/qe/nlarith_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/nlarith_util$(OBJ_EXT) ../src/qe/nlarith_util.cpp
qe/qsat$(OBJ_EXT): ../src/qe/qsat.cpp ast/expr_abstract.h.node ast/ast_util.h.node ast/occurs.h.node ast/rewriter/quant_hoist.h.node ast/ast_pp.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/expr_replacer.h.node model/model_v2_pp.h.node model/model_evaluator.h.node model/model_evaluator_params.hpp.node smt/smt_kernel.h.node smt/params/smt_params.h.node smt/smt_solver.h.node solver/solver.h.node solver/mus.h.node qe/qsat.h.node qe/qe_mbp.h.node qe/qe.h.node ast/rewriter/label_rewriter.h.node
	@echo src/qe/qsat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qsat$(OBJ_EXT) ../src/qe/qsat.cpp
qe/qe_dl_plugin$(OBJ_EXT): ../src/qe/qe_dl_plugin.cpp qe/qe.h.node ast/rewriter/expr_safe_replace.h.node ast/dl_decl_plugin.h.node util/obj_pair_hashtable.h.node ast/ast_pp.h.node
	@echo src/qe/qe_dl_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_dl_plugin$(OBJ_EXT) ../src/qe/qe_dl_plugin.cpp
qe/qe_array_plugin$(OBJ_EXT): ../src/qe/qe_array_plugin.cpp qe/qe.h.node ast/array_decl_plugin.h.node ast/rewriter/expr_safe_replace.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node
	@echo src/qe/qe_array_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_array_plugin$(OBJ_EXT) ../src/qe/qe_array_plugin.cpp
qe/qe_cmd.h.node: ../src/qe/qe_cmd.h
	@echo done > qe/qe_cmd.h.node
qe/qe_cmd$(OBJ_EXT): ../src/qe/qe_cmd.cpp qe/qe_cmd.h.node qe/qe.h.node cmd_context/cmd_context.h.node cmd_context/parametric_cmd.h.node
	@echo src/qe/qe_cmd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_cmd$(OBJ_EXT) ../src/qe/qe_cmd.cpp
qe/qe$(OBJ_EXT): ../src/qe/qe.cpp qe/qe.h.node smt/smt_theory.h.node ast/bv_decl_plugin.h.node smt/smt_context.h.node smt/theory_bv.h.node ast/ast_ll_pp.h.node ast/ast_pp.h.node ast/ast_smt_pp.h.node ast/expr_abstract.h.node ast/rewriter/var_subst.h.node ast/for_each_expr.h.node ast/dl_decl_plugin.h.node qe/nlarith_util.h.node ast/rewriter/expr_replacer.h.node ast/rewriter/factor_rewriter.h.node ast/expr_functors.h.node ast/rewriter/quant_hoist.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node smt/smt_kernel.h.node model/model_evaluator.h.node ast/has_free_vars.h.node ast/rewriter/rewriter_def.h.node tactic/tactical.h.node model/model_v2_pp.h.node util/obj_hashtable.h.node
	@echo src/qe/qe.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe$(OBJ_EXT) ../src/qe/qe.cpp
qe/qe_lite$(OBJ_EXT): ../src/qe/qe_lite.cpp util/uint_set.h.node ast/expr_abstract.h.node ast/used_vars.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt2_pp.h.node ast/is_variable_test.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/var_subst.h.node ast/ast_util.h.node ast/rewriter/th_rewriter.h.node ast/for_each_expr.h.node ast/rewriter/expr_safe_replace.h.node ast/datatype_decl_plugin.h.node tactic/tactical.h.node qe/mbp/mbp_solve_plugin.h.node qe/qe_lite.h.node
	@echo src/qe/qe_lite.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_lite$(OBJ_EXT) ../src/qe/qe_lite.cpp
qe/qe_bv_plugin$(OBJ_EXT): ../src/qe/qe_bv_plugin.cpp qe/qe.h.node ast/rewriter/expr_safe_replace.h.node ast/bv_decl_plugin.h.node model/model_evaluator.h.node
	@echo src/qe/qe_bv_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_46) $(CXX_OUT_FLAG)qe/qe_bv_plugin$(OBJ_EXT) ../src/qe/qe_bv_plugin.cpp
qe/qe$(LIB_EXT): qe/qe_tactic$(OBJ_EXT) qe/qe_datatype_plugin$(OBJ_EXT) qe/nlqsat$(OBJ_EXT) qe/qe_mbi$(OBJ_EXT) qe/qe_mbp$(OBJ_EXT) qe/qe_arith_plugin$(OBJ_EXT) qe/qe_bool_plugin$(OBJ_EXT) qe/nlarith_util$(OBJ_EXT) qe/qsat$(OBJ_EXT) qe/qe_dl_plugin$(OBJ_EXT) qe/qe_array_plugin$(OBJ_EXT) qe/qe_cmd$(OBJ_EXT) qe/qe$(OBJ_EXT) qe/qe_lite$(OBJ_EXT) qe/qe_bv_plugin$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)qe/qe$(LIB_EXT) qe/qe_tactic$(OBJ_EXT) qe/qe_datatype_plugin$(OBJ_EXT) qe/nlqsat$(OBJ_EXT) qe/qe_mbi$(OBJ_EXT) qe/qe_mbp$(OBJ_EXT) qe/qe_arith_plugin$(OBJ_EXT) qe/qe_bool_plugin$(OBJ_EXT) qe/nlarith_util$(OBJ_EXT) qe/qsat$(OBJ_EXT) qe/qe_dl_plugin$(OBJ_EXT) qe/qe_array_plugin$(OBJ_EXT) qe/qe_cmd$(OBJ_EXT) qe/qe$(OBJ_EXT) qe/qe_lite$(OBJ_EXT) qe/qe_bv_plugin$(OBJ_EXT)
qe: qe/qe$(LIB_EXT)

includes_47 = -I../src/solver -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/tactic -I../src/ast/proofs -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src
tactic/aig/aig_tactic.h.node: ../src/tactic/aig/aig_tactic.h util/params.h.node
	@echo done > tactic/aig/aig_tactic.h.node
tactic/bv/bit_blaster_tactic.h.node: ../src/tactic/bv/bit_blaster_tactic.h util/params.h.node ast/rewriter/bit_blaster/bit_blaster_rewriter.h.node
	@echo done > tactic/bv/bit_blaster_tactic.h.node
sat/tactic/sat_tactic.h.node: ../src/sat/tactic/sat_tactic.h util/params.h.node
	@echo done > sat/tactic/sat_tactic.h.node
sat/sat_solver/inc_sat_solver$(OBJ_EXT): ../src/sat/sat_solver/inc_sat_solver.cpp util/gparams.h.node util/stacked_value.h.node ast/ast_pp.h.node ast/ast_translation.h.node ast/ast_util.h.node solver/solver.h.node solver/tactic2solver.h.node solver/parallel_params.hpp.node solver/parallel_tactic.h.node tactic/tactical.h.node tactic/aig/aig_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/arith/card2bv_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/core/simplify_tactic.h.node tactic/bv/bit_blaster_model_converter.h.node model/model_smt2_pp.h.node model/model_v2_pp.h.node model/model_evaluator.h.node sat/sat_solver.h.node sat/sat_params.hpp.node sat/smt/euf_solver.h.node sat/tactic/goal2sat.h.node sat/tactic/sat_tactic.h.node sat/sat_simplifier_params.hpp.node
	@echo src/sat/sat_solver/inc_sat_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_47) $(CXX_OUT_FLAG)sat/sat_solver/inc_sat_solver$(OBJ_EXT) ../src/sat/sat_solver/inc_sat_solver.cpp
sat/sat_solver/sat_solver$(LIB_EXT): sat/sat_solver/inc_sat_solver$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)sat/sat_solver/sat_solver$(LIB_EXT) sat/sat_solver/inc_sat_solver$(OBJ_EXT)
sat_solver: sat/sat_solver/sat_solver$(LIB_EXT)

includes_48 = -I../src/tactic/core -I../src/tactic -I../src/ast -I../src/util -I../src/math/polynomial -I../src/model -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/sat_solver -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src
tactic/fd_solver/enum2bv_solver.h.node: ../src/tactic/fd_solver/enum2bv_solver.h ast/ast.h.node util/params.h.node
	@echo done > tactic/fd_solver/enum2bv_solver.h.node
tactic/fd_solver/enum2bv_solver$(OBJ_EXT): ../src/tactic/fd_solver/enum2bv_solver.cpp ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/enum2bv_rewriter.h.node model/model_smt2_pp.h.node tactic/tactic.h.node tactic/generic_model_converter.h.node tactic/fd_solver/enum2bv_solver.h.node solver/solver_na2as.h.node
	@echo src/tactic/fd_solver/enum2bv_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_48) $(CXX_OUT_FLAG)tactic/fd_solver/enum2bv_solver$(OBJ_EXT) ../src/tactic/fd_solver/enum2bv_solver.cpp
tactic/fd_solver/pb2bv_solver.h.node: ../src/tactic/fd_solver/pb2bv_solver.h ast/ast.h.node util/params.h.node
	@echo done > tactic/fd_solver/pb2bv_solver.h.node
tactic/fd_solver/pb2bv_solver$(OBJ_EXT): ../src/tactic/fd_solver/pb2bv_solver.cpp util/statistics.h.node ast/ast_pp.h.node ast/rewriter/pb2bv_rewriter.h.node ast/rewriter/th_rewriter.h.node model/model_smt2_pp.h.node tactic/tactic.h.node tactic/generic_model_converter.h.node solver/solver_na2as.h.node tactic/fd_solver/pb2bv_solver.h.node
	@echo src/tactic/fd_solver/pb2bv_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_48) $(CXX_OUT_FLAG)tactic/fd_solver/pb2bv_solver$(OBJ_EXT) ../src/tactic/fd_solver/pb2bv_solver.cpp
tactic/fd_solver/fd_solver.h.node: ../src/tactic/fd_solver/fd_solver.h ast/ast.h.node util/params.h.node
	@echo done > tactic/fd_solver/fd_solver.h.node
sat/sat_solver/inc_sat_solver.h.node: ../src/sat/sat_solver/inc_sat_solver.h solver/solver.h.node
	@echo done > sat/sat_solver/inc_sat_solver.h.node
tactic/fd_solver/bounded_int2bv_solver.h.node: ../src/tactic/fd_solver/bounded_int2bv_solver.h ast/ast.h.node util/params.h.node
	@echo done > tactic/fd_solver/bounded_int2bv_solver.h.node
tactic/fd_solver/fd_solver$(OBJ_EXT): ../src/tactic/fd_solver/fd_solver.cpp tactic/fd_solver/fd_solver.h.node tactic/tactic.h.node sat/sat_solver/inc_sat_solver.h.node tactic/fd_solver/enum2bv_solver.h.node tactic/fd_solver/pb2bv_solver.h.node tactic/fd_solver/bounded_int2bv_solver.h.node solver/solver2tactic.h.node solver/parallel_tactic.h.node solver/parallel_params.hpp.node
	@echo src/tactic/fd_solver/fd_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_48) $(CXX_OUT_FLAG)tactic/fd_solver/fd_solver$(OBJ_EXT) ../src/tactic/fd_solver/fd_solver.cpp
tactic/fd_solver/smtfd_solver$(OBJ_EXT): ../src/tactic/fd_solver/smtfd_solver.cpp util/scoped_ptr_vector.h.node util/obj_hashtable.h.node util/obj_pair_hashtable.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/for_each_expr.h.node ast/pb_decl_plugin.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/var_subst.h.node tactic/tactic_exception.h.node tactic/fd_solver/fd_solver.h.node solver/solver.h.node solver/solver_na2as.h.node solver/solver2tactic.h.node
	@echo src/tactic/fd_solver/smtfd_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_48) $(CXX_OUT_FLAG)tactic/fd_solver/smtfd_solver$(OBJ_EXT) ../src/tactic/fd_solver/smtfd_solver.cpp
tactic/fd_solver/bounded_int2bv_solver$(OBJ_EXT): ../src/tactic/fd_solver/bounded_int2bv_solver.cpp tactic/fd_solver/bounded_int2bv_solver.h.node solver/solver_na2as.h.node tactic/tactic.h.node ast/rewriter/pb2bv_rewriter.h.node tactic/generic_model_converter.h.node ast/ast_pp.h.node model/model_smt2_pp.h.node tactic/arith/bound_manager.h.node tactic/arith/bv2int_rewriter.h.node ast/rewriter/expr_safe_replace.h.node ast/bv_decl_plugin.h.node ast/arith_decl_plugin.h.node
	@echo src/tactic/fd_solver/bounded_int2bv_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_48) $(CXX_OUT_FLAG)tactic/fd_solver/bounded_int2bv_solver$(OBJ_EXT) ../src/tactic/fd_solver/bounded_int2bv_solver.cpp
tactic/fd_solver/fd_solver$(LIB_EXT): tactic/fd_solver/enum2bv_solver$(OBJ_EXT) tactic/fd_solver/pb2bv_solver$(OBJ_EXT) tactic/fd_solver/fd_solver$(OBJ_EXT) tactic/fd_solver/smtfd_solver$(OBJ_EXT) tactic/fd_solver/bounded_int2bv_solver$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/fd_solver/fd_solver$(LIB_EXT) tactic/fd_solver/enum2bv_solver$(OBJ_EXT) tactic/fd_solver/pb2bv_solver$(OBJ_EXT) tactic/fd_solver/fd_solver$(OBJ_EXT) tactic/fd_solver/smtfd_solver$(OBJ_EXT) tactic/fd_solver/bounded_int2bv_solver$(OBJ_EXT)
fd_solver: tactic/fd_solver/fd_solver$(LIB_EXT)

includes_49 = -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src
muz/base/dl_boogie_proof.h.node: ../src/muz/base/dl_boogie_proof.h ast/ast.h.node model/model.h.node
	@echo done > muz/base/dl_boogie_proof.h.node
muz/base/dl_boogie_proof$(OBJ_EXT): ../src/muz/base/dl_boogie_proof.cpp muz/base/dl_boogie_proof.h.node model/model_pp.h.node ast/proofs/proof_utils.h.node ast/ast_pp.h.node ast/ast_util.h.node
	@echo src/muz/base/dl_boogie_proof.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_boogie_proof$(OBJ_EXT) ../src/muz/base/dl_boogie_proof.cpp
muz/base/dl_costs.h.node: ../src/muz/base/dl_costs.h ast/ast.h.node
	@echo done > muz/base/dl_costs.h.node
muz/base/dl_util.h.node: ../src/muz/base/dl_util.h ast/ast.h.node util/hashtable.h.node util/obj_hashtable.h.node util/uint_set.h.node tactic/horn_subsume_model_converter.h.node tactic/replace_proof_converter.h.node ast/substitution/substitution.h.node ast/rewriter/ast_counter.h.node util/statistics.h.node util/stopwatch.h.node util/lbool.h.node util/container_util.h.node
	@echo done > muz/base/dl_util.h.node
muz/base/hnf.h.node: ../src/muz/base/hnf.h ast/ast.h.node util/params.h.node ast/normal_forms/defined_names.h.node tactic/proof_converter.h.node
	@echo done > muz/base/hnf.h.node
muz/base/dl_rule.h.node: ../src/muz/base/dl_rule.h ast/ast.h.node muz/base/dl_costs.h.node muz/base/dl_util.h.node ast/used_vars.h.node tactic/proof_converter.h.node tactic/model_converter.h.node ast/rewriter/ast_counter.h.node ast/rewriter/rewriter.h.node muz/base/hnf.h.node qe/qe_lite.h.node ast/rewriter/var_subst.h.node ast/datatype_decl_plugin.h.node ast/rewriter/label_rewriter.h.node
	@echo done > muz/base/dl_rule.h.node
muz/base/dl_rule_set.h.node: ../src/muz/base/dl_rule_set.h util/obj_hashtable.h.node muz/base/dl_rule.h.node
	@echo done > muz/base/dl_rule_set.h.node
muz/base/dl_rule_transformer.h.node: ../src/muz/base/dl_rule_transformer.h util/map.h.node util/vector.h.node muz/base/dl_rule.h.node muz/base/dl_rule_set.h.node
	@echo done > muz/base/dl_rule_transformer.h.node
muz/base/dl_engine_base.h.node: ../src/muz/base/dl_engine_base.h model/model.h.node muz/base/dl_util.h.node
	@echo done > muz/base/dl_engine_base.h.node
muz/base/bind_variables.h.node: ../src/muz/base/bind_variables.h ast/ast.h.node
	@echo done > muz/base/bind_variables.h.node
muz/base/rule_properties.h.node: ../src/muz/base/rule_properties.h ast/ast.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/recfun_decl_plugin.h.node muz/base/dl_rule.h.node ast/expr_functors.h.node
	@echo done > muz/base/rule_properties.h.node
muz/base/dl_context.h.node: ../src/muz/base/dl_context.h ast/arith_decl_plugin.h.node util/map.h.node ast/rewriter/th_rewriter.h.node util/str_hashtable.h.node ast/rewriter/var_subst.h.node muz/base/dl_costs.h.node ast/dl_decl_plugin.h.node muz/base/dl_rule_set.h.node util/lbool.h.node util/statistics.h.node util/params.h.node util/trail.h.node tactic/model_converter.h.node model/model2expr.h.node smt/params/smt_params.h.node muz/base/dl_rule_transformer.h.node ast/expr_functors.h.node muz/base/dl_engine_base.h.node muz/base/bind_variables.h.node muz/base/rule_properties.h.node
	@echo done > muz/base/dl_context.h.node
muz/base/dl_util$(OBJ_EXT): ../src/muz/base/dl_util.cpp ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node ast/for_each_expr.h.node ast/scoped_proof.h.node muz/base/dl_context.h.node muz/base/dl_rule.h.node muz/base/dl_util.h.node util/stopwatch.h.node
	@echo src/muz/base/dl_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_util$(OBJ_EXT) ../src/muz/base/dl_util.cpp
muz/base/dl_rule_subsumption_index.h.node: ../src/muz/base/dl_rule_subsumption_index.h muz/base/dl_context.h.node
	@echo done > muz/base/dl_rule_subsumption_index.h.node
muz/base/dl_rule_subsumption_index$(OBJ_EXT): ../src/muz/base/dl_rule_subsumption_index.cpp ast/ast_pp.h.node muz/base/dl_rule_subsumption_index.h.node
	@echo src/muz/base/dl_rule_subsumption_index.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_rule_subsumption_index$(OBJ_EXT) ../src/muz/base/dl_rule_subsumption_index.cpp
muz/base/fp_params.hpp.node: ../src/muz/base/fp_params.hpp util/params.h.node util/gparams.h.node
	@echo done > muz/base/fp_params.hpp.node
muz/base/dl_context$(OBJ_EXT): ../src/muz/base/dl_context.cpp ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/for_each_expr.h.node ast/ast_smt_pp.h.node ast/ast_smt2_pp.h.node ast/datatype_decl_plugin.h.node ast/scoped_proof.h.node ast/ast_pp_util.h.node ast/ast_util.h.node muz/base/dl_context.h.node muz/base/fp_params.hpp.node
	@echo src/muz/base/dl_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_context$(OBJ_EXT) ../src/muz/base/dl_context.cpp
muz/base/hnf$(OBJ_EXT): ../src/muz/base/hnf.cpp muz/base/hnf.h.node util/warning.h.node ast/used_vars.h.node ast/well_sorted.h.node ast/rewriter/var_subst.h.node ast/normal_forms/name_exprs.h.node ast/act_cache.h.node ast/ast_pp.h.node ast/rewriter/quant_hoist.h.node ast/ast_util.h.node muz/base/dl_util.h.node ast/for_each_ast.h.node ast/for_each_expr.h.node
	@echo src/muz/base/hnf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/hnf$(OBJ_EXT) ../src/muz/base/hnf.cpp
muz/base/rule_properties$(OBJ_EXT): ../src/muz/base/rule_properties.cpp ast/expr_functors.h.node ast/for_each_expr.h.node ast/rewriter/th_rewriter.h.node muz/base/rule_properties.h.node muz/base/dl_rule_set.h.node muz/base/dl_context.h.node
	@echo src/muz/base/rule_properties.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/rule_properties$(OBJ_EXT) ../src/muz/base/rule_properties.cpp
muz/base/dl_costs$(OBJ_EXT): ../src/muz/base/dl_costs.cpp util/debug.h.node util/stopwatch.h.node muz/base/dl_context.h.node muz/base/dl_rule.h.node muz/base/dl_costs.h.node
	@echo src/muz/base/dl_costs.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_costs$(OBJ_EXT) ../src/muz/base/dl_costs.cpp
muz/base/bind_variables$(OBJ_EXT): ../src/muz/base/bind_variables.cpp muz/base/bind_variables.h.node
	@echo src/muz/base/bind_variables.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/bind_variables$(OBJ_EXT) ../src/muz/base/bind_variables.cpp
muz/base/dl_rule_transformer$(OBJ_EXT): ../src/muz/base/dl_rule_transformer.cpp muz/base/dl_context.h.node muz/base/dl_rule_transformer.h.node util/stopwatch.h.node
	@echo src/muz/base/dl_rule_transformer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_rule_transformer$(OBJ_EXT) ../src/muz/base/dl_rule_transformer.cpp
muz/base/dl_rule_set$(OBJ_EXT): ../src/muz/base/dl_rule_set.cpp muz/base/dl_context.h.node muz/base/dl_rule_set.h.node ast/ast_pp.h.node
	@echo src/muz/base/dl_rule_set.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_rule_set$(OBJ_EXT) ../src/muz/base/dl_rule_set.cpp
muz/base/dl_rule$(OBJ_EXT): ../src/muz/base/dl_rule.cpp ast/ast_pp.h.node muz/base/dl_context.h.node util/map.h.node ast/recurse_expr_def.h.node muz/base/dl_rule.h.node qe/qe.h.node ast/for_each_expr.h.node ast/used_vars.h.node ast/rewriter/var_subst.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/th_rewriter.h.node ast/ast_smt2_pp.h.node ast/used_symbols.h.node ast/rewriter/quant_hoist.h.node ast/rewriter/expr_replacer.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/expr_safe_replace.h.node tactic/generic_model_converter.h.node ast/scoped_proof.h.node ast/datatype_decl_plugin.h.node ast/ast_util.h.node
	@echo src/muz/base/dl_rule.cpp
	@$(CXX) $(CXXFLAGS) $(includes_49) $(CXX_OUT_FLAG)muz/base/dl_rule$(OBJ_EXT) ../src/muz/base/dl_rule.cpp
muz/base/muz$(LIB_EXT): muz/base/dl_boogie_proof$(OBJ_EXT) muz/base/dl_util$(OBJ_EXT) muz/base/dl_rule_subsumption_index$(OBJ_EXT) muz/base/dl_context$(OBJ_EXT) muz/base/hnf$(OBJ_EXT) muz/base/rule_properties$(OBJ_EXT) muz/base/dl_costs$(OBJ_EXT) muz/base/bind_variables$(OBJ_EXT) muz/base/dl_rule_transformer$(OBJ_EXT) muz/base/dl_rule_set$(OBJ_EXT) muz/base/dl_rule$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/base/muz$(LIB_EXT) muz/base/dl_boogie_proof$(OBJ_EXT) muz/base/dl_util$(OBJ_EXT) muz/base/dl_rule_subsumption_index$(OBJ_EXT) muz/base/dl_context$(OBJ_EXT) muz/base/hnf$(OBJ_EXT) muz/base/rule_properties$(OBJ_EXT) muz/base/dl_costs$(OBJ_EXT) muz/base/bind_variables$(OBJ_EXT) muz/base/dl_rule_transformer$(OBJ_EXT) muz/base/dl_rule_set$(OBJ_EXT) muz/base/dl_rule$(OBJ_EXT)
muz: muz/base/muz$(LIB_EXT)

includes_50 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src
muz/dataflow/dataflow.h.node: ../src/muz/dataflow/dataflow.h muz/base/dl_rule.h.node muz/base/dl_rule_set.h.node util/hashtable.h.node util/vector.h.node
	@echo done > muz/dataflow/dataflow.h.node
muz/dataflow/reachability.h.node: ../src/muz/dataflow/reachability.h muz/dataflow/dataflow.h.node
	@echo done > muz/dataflow/reachability.h.node
muz/dataflow/dataflow$(OBJ_EXT): ../src/muz/dataflow/dataflow.cpp muz/dataflow/dataflow.h.node muz/dataflow/reachability.h.node
	@echo src/muz/dataflow/dataflow.cpp
	@$(CXX) $(CXXFLAGS) $(includes_50) $(CXX_OUT_FLAG)muz/dataflow/dataflow$(OBJ_EXT) ../src/muz/dataflow/dataflow.cpp
muz/dataflow/dataflow$(LIB_EXT): muz/dataflow/dataflow$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/dataflow/dataflow$(LIB_EXT) muz/dataflow/dataflow$(OBJ_EXT)
dataflow: muz/dataflow/dataflow$(LIB_EXT)

includes_51 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/math/hilbert -I../src/muz/dataflow -I../src
muz/transforms/dl_mk_bit_blast.h.node: ../src/muz/transforms/dl_mk_bit_blast.h muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_bit_blast.h.node
muz/transforms/dl_mk_interp_tail_simplifier.h.node: ../src/muz/transforms/dl_mk_interp_tail_simplifier.h muz/base/dl_context.h.node muz/base/dl_rule_transformer.h.node ast/substitution/unifier.h.node ast/substitution/substitution.h.node ast/arith_decl_plugin.h.node
	@echo done > muz/transforms/dl_mk_interp_tail_simplifier.h.node
muz/transforms/dl_mk_bit_blast$(OBJ_EXT): ../src/muz/transforms/dl_mk_bit_blast.cpp muz/transforms/dl_mk_bit_blast.h.node ast/rewriter/bit_blaster/bit_blaster_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node ast/rewriter/expr_safe_replace.h.node tactic/generic_model_converter.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node muz/base/fp_params.hpp.node ast/scoped_proof.h.node model/model_v2_pp.h.node
	@echo src/muz/transforms/dl_mk_bit_blast.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_bit_blast$(OBJ_EXT) ../src/muz/transforms/dl_mk_bit_blast.cpp
muz/transforms/dl_mk_rule_inliner.h.node: ../src/muz/transforms/dl_mk_rule_inliner.h muz/base/dl_context.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node ast/substitution/unifier.h.node ast/substitution/substitution.h.node ast/substitution/substitution_tree.h.node
	@echo done > muz/transforms/dl_mk_rule_inliner.h.node
muz/transforms/dl_mk_rule_inliner$(OBJ_EXT): ../src/muz/transforms/dl_mk_rule_inliner.cpp ast/ast_pp.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/base/fp_params.hpp.node
	@echo src/muz/transforms/dl_mk_rule_inliner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_rule_inliner$(OBJ_EXT) ../src/muz/transforms/dl_mk_rule_inliner.cpp
muz/transforms/dl_mk_coalesce.h.node: ../src/muz/transforms/dl_mk_coalesce.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node util/uint_set.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_rule_inliner.h.node
	@echo done > muz/transforms/dl_mk_coalesce.h.node
muz/transforms/dl_mk_coalesce$(OBJ_EXT): ../src/muz/transforms/dl_mk_coalesce.cpp muz/transforms/dl_mk_coalesce.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/muz/transforms/dl_mk_coalesce.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_coalesce$(OBJ_EXT) ../src/muz/transforms/dl_mk_coalesce.cpp
muz/transforms/dl_mk_scale.h.node: ../src/muz/transforms/dl_mk_scale.h muz/base/dl_rule_transformer.h.node ast/arith_decl_plugin.h.node
	@echo done > muz/transforms/dl_mk_scale.h.node
muz/transforms/dl_mk_scale$(OBJ_EXT): ../src/muz/transforms/dl_mk_scale.cpp muz/transforms/dl_mk_scale.h.node muz/base/dl_context.h.node muz/base/fp_params.hpp.node
	@echo src/muz/transforms/dl_mk_scale.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_scale$(OBJ_EXT) ../src/muz/transforms/dl_mk_scale.cpp
muz/transforms/dl_mk_filter_rules.h.node: ../src/muz/transforms/dl_mk_filter_rules.h util/map.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_filter_rules.h.node
muz/transforms/dl_mk_filter_rules$(OBJ_EXT): ../src/muz/transforms/dl_mk_filter_rules.cpp muz/transforms/dl_mk_filter_rules.h.node muz/base/dl_context.h.node ast/for_each_expr.h.node ast/ast_pp.h.node
	@echo src/muz/transforms/dl_mk_filter_rules.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_filter_rules$(OBJ_EXT) ../src/muz/transforms/dl_mk_filter_rules.cpp
muz/transforms/dl_mk_interp_tail_simplifier$(OBJ_EXT): ../src/muz/transforms/dl_mk_interp_tail_simplifier.cpp ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node ast/ast_util.h.node muz/base/fp_params.hpp.node
	@echo src/muz/transforms/dl_mk_interp_tail_simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_interp_tail_simplifier$(OBJ_EXT) ../src/muz/transforms/dl_mk_interp_tail_simplifier.cpp
muz/transforms/dl_mk_quantifier_abstraction.h.node: ../src/muz/transforms/dl_mk_quantifier_abstraction.h muz/base/dl_rule_transformer.h.node ast/array_decl_plugin.h.node
	@echo done > muz/transforms/dl_mk_quantifier_abstraction.h.node
muz/transforms/dl_mk_quantifier_abstraction$(OBJ_EXT): ../src/muz/transforms/dl_mk_quantifier_abstraction.cpp muz/transforms/dl_mk_quantifier_abstraction.h.node muz/base/dl_context.h.node ast/rewriter/expr_safe_replace.h.node ast/expr_abstract.h.node
	@echo src/muz/transforms/dl_mk_quantifier_abstraction.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_quantifier_abstraction$(OBJ_EXT) ../src/muz/transforms/dl_mk_quantifier_abstraction.cpp
muz/transforms/dl_transforms.h.node: ../src/muz/transforms/dl_transforms.h muz/base/dl_context.h.node
	@echo done > muz/transforms/dl_transforms.h.node
muz/transforms/dl_mk_coi_filter.h.node: ../src/muz/transforms/dl_mk_coi_filter.h muz/base/dl_rule_transformer.h.node muz/base/dl_context.h.node
	@echo done > muz/transforms/dl_mk_coi_filter.h.node
muz/transforms/dl_mk_array_blast.h.node: ../src/muz/transforms/dl_mk_array_blast.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node tactic/equiv_proof_converter.h.node ast/array_decl_plugin.h.node ast/rewriter/expr_safe_replace.h.node
	@echo done > muz/transforms/dl_mk_array_blast.h.node
muz/transforms/dl_mk_karr_invariants.h.node: ../src/muz/transforms/dl_mk_karr_invariants.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node ast/arith_decl_plugin.h.node math/hilbert/hilbert_basis.h.node
	@echo done > muz/transforms/dl_mk_karr_invariants.h.node
muz/transforms/dl_mk_magic_symbolic.h.node: ../src/muz/transforms/dl_mk_magic_symbolic.h muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_magic_symbolic.h.node
muz/transforms/dl_mk_quantifier_instantiation.h.node: ../src/muz/transforms/dl_mk_quantifier_instantiation.h muz/base/dl_rule_transformer.h.node ast/rewriter/expr_safe_replace.h.node util/union_find.h.node
	@echo done > muz/transforms/dl_mk_quantifier_instantiation.h.node
muz/transforms/dl_mk_subsumption_checker.h.node: ../src/muz/transforms/dl_mk_subsumption_checker.h muz/base/dl_context.h.node muz/base/dl_rule_transformer.h.node muz/base/dl_rule_subsumption_index.h.node
	@echo done > muz/transforms/dl_mk_subsumption_checker.h.node
muz/transforms/dl_mk_array_eq_rewrite.h.node: ../src/muz/transforms/dl_mk_array_eq_rewrite.h muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_array_eq_rewrite.h.node
muz/transforms/dl_mk_array_instantiation.h.node: ../src/muz/transforms/dl_mk_array_instantiation.h ast/rewriter/factor_equivs.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_array_instantiation.h.node
muz/transforms/dl_mk_elim_term_ite.h.node: ../src/muz/transforms/dl_mk_elim_term_ite.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node tactic/equiv_proof_converter.h.node
	@echo done > muz/transforms/dl_mk_elim_term_ite.h.node
muz/transforms/dl_transforms$(OBJ_EXT): ../src/muz/transforms/dl_transforms.cpp muz/transforms/dl_transforms.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_coi_filter.h.node muz/transforms/dl_mk_filter_rules.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/transforms/dl_mk_bit_blast.h.node muz/transforms/dl_mk_array_blast.h.node muz/transforms/dl_mk_karr_invariants.h.node muz/transforms/dl_mk_magic_symbolic.h.node muz/transforms/dl_mk_quantifier_abstraction.h.node muz/transforms/dl_mk_quantifier_instantiation.h.node muz/transforms/dl_mk_subsumption_checker.h.node muz/transforms/dl_mk_scale.h.node muz/transforms/dl_mk_array_eq_rewrite.h.node muz/transforms/dl_mk_array_instantiation.h.node muz/transforms/dl_mk_elim_term_ite.h.node muz/base/fp_params.hpp.node
	@echo src/muz/transforms/dl_transforms.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_transforms$(OBJ_EXT) ../src/muz/transforms/dl_transforms.cpp
muz/transforms/dl_mk_backwards.h.node: ../src/muz/transforms/dl_mk_backwards.h muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_backwards.h.node
muz/transforms/dl_mk_loop_counter.h.node: ../src/muz/transforms/dl_mk_loop_counter.h muz/base/dl_rule_transformer.h.node ast/arith_decl_plugin.h.node
	@echo done > muz/transforms/dl_mk_loop_counter.h.node
muz/transforms/dl_mk_karr_invariants$(OBJ_EXT): ../src/muz/transforms/dl_mk_karr_invariants.cpp ast/rewriter/expr_safe_replace.h.node ast/rewriter/bool_rewriter.h.node ast/for_each_expr.h.node muz/transforms/dl_mk_karr_invariants.h.node muz/transforms/dl_mk_backwards.h.node muz/transforms/dl_mk_loop_counter.h.node
	@echo src/muz/transforms/dl_mk_karr_invariants.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_karr_invariants$(OBJ_EXT) ../src/muz/transforms/dl_mk_karr_invariants.cpp
muz/transforms/dl_mk_backwards$(OBJ_EXT): ../src/muz/transforms/dl_mk_backwards.cpp muz/transforms/dl_mk_backwards.h.node muz/base/dl_context.h.node
	@echo src/muz/transforms/dl_mk_backwards.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_backwards$(OBJ_EXT) ../src/muz/transforms/dl_mk_backwards.cpp
muz/transforms/dl_mk_array_instantiation$(OBJ_EXT): ../src/muz/transforms/dl_mk_array_instantiation.cpp muz/transforms/dl_mk_array_instantiation.h.node muz/base/dl_context.h.node ast/pattern/pattern_inference.h.node muz/base/dl_context.h.node ast/rewriter/expr_safe_replace.h.node ast/expr_abstract.h.node muz/base/fp_params.hpp.node
	@echo src/muz/transforms/dl_mk_array_instantiation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_array_instantiation$(OBJ_EXT) ../src/muz/transforms/dl_mk_array_instantiation.cpp
muz/transforms/dl_mk_array_eq_rewrite$(OBJ_EXT): ../src/muz/transforms/dl_mk_array_eq_rewrite.cpp ast/pattern/pattern_inference.h.node ast/rewriter/expr_safe_replace.h.node ast/expr_abstract.h.node muz/base/dl_context.h.node muz/base/dl_context.h.node muz/base/fp_params.hpp.node muz/transforms/dl_mk_array_eq_rewrite.h.node ast/rewriter/factor_equivs.h.node
	@echo src/muz/transforms/dl_mk_array_eq_rewrite.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_array_eq_rewrite$(OBJ_EXT) ../src/muz/transforms/dl_mk_array_eq_rewrite.cpp
muz/transforms/dl_mk_subsumption_checker$(OBJ_EXT): ../src/muz/transforms/dl_mk_subsumption_checker.cpp ast/ast_pp.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node muz/transforms/dl_mk_subsumption_checker.h.node muz/base/fp_params.hpp.node tactic/generic_model_converter.h.node
	@echo src/muz/transforms/dl_mk_subsumption_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_subsumption_checker$(OBJ_EXT) ../src/muz/transforms/dl_mk_subsumption_checker.cpp
tactic/core/blast_term_ite_tactic.h.node: ../src/tactic/core/blast_term_ite_tactic.h util/params.h.node
	@echo done > tactic/core/blast_term_ite_tactic.h.node
muz/transforms/dl_mk_elim_term_ite$(OBJ_EXT): ../src/muz/transforms/dl_mk_elim_term_ite.cpp ast/ast_util.h.node ast/expr_abstract.h.node tactic/core/blast_term_ite_tactic.h.node tactic/tactic.h.node tactic/core/elim_term_ite_tactic.h.node muz/transforms/dl_mk_elim_term_ite.h.node
	@echo src/muz/transforms/dl_mk_elim_term_ite.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_elim_term_ite$(OBJ_EXT) ../src/muz/transforms/dl_mk_elim_term_ite.cpp
muz/transforms/dl_mk_unbound_compressor.h.node: ../src/muz/transforms/dl_mk_unbound_compressor.h util/map.h.node util/obj_pair_hashtable.h.node muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_unbound_compressor.h.node
muz/transforms/dl_mk_unbound_compressor$(OBJ_EXT): ../src/muz/transforms/dl_mk_unbound_compressor.cpp muz/transforms/dl_mk_unbound_compressor.h.node
	@echo src/muz/transforms/dl_mk_unbound_compressor.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_unbound_compressor$(OBJ_EXT) ../src/muz/transforms/dl_mk_unbound_compressor.cpp
muz/transforms/dl_mk_quantifier_instantiation$(OBJ_EXT): ../src/muz/transforms/dl_mk_quantifier_instantiation.cpp muz/transforms/dl_mk_quantifier_instantiation.h.node muz/base/dl_context.h.node ast/pattern/pattern_inference.h.node ast/rewriter/rewriter_def.h.node ast/ast_util.h.node
	@echo src/muz/transforms/dl_mk_quantifier_instantiation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_quantifier_instantiation$(OBJ_EXT) ../src/muz/transforms/dl_mk_quantifier_instantiation.cpp
muz/transforms/dl_mk_slice.h.node: ../src/muz/transforms/dl_mk_slice.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node util/uint_set.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_slice.h.node
muz/transforms/dl_mk_slice$(OBJ_EXT): ../src/muz/transforms/dl_mk_slice.cpp muz/transforms/dl_mk_slice.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/expr_functors.h.node muz/transforms/dl_mk_rule_inliner.h.node model/model_smt2_pp.h.node
	@echo src/muz/transforms/dl_mk_slice.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_slice$(OBJ_EXT) ../src/muz/transforms/dl_mk_slice.cpp
muz/transforms/dl_mk_coi_filter$(OBJ_EXT): ../src/muz/transforms/dl_mk_coi_filter.cpp muz/transforms/dl_mk_coi_filter.h.node muz/dataflow/dataflow.h.node muz/dataflow/reachability.h.node ast/ast_pp.h.node tactic/generic_model_converter.h.node ast/ast_util.h.node
	@echo src/muz/transforms/dl_mk_coi_filter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_coi_filter$(OBJ_EXT) ../src/muz/transforms/dl_mk_coi_filter.cpp
muz/transforms/dl_mk_magic_sets.h.node: ../src/muz/transforms/dl_mk_magic_sets.h util/map.h.node util/obj_pair_hashtable.h.node muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/transforms/dl_mk_magic_sets.h.node
muz/transforms/dl_mk_magic_sets$(OBJ_EXT): ../src/muz/transforms/dl_mk_magic_sets.cpp ast/ast_pp.h.node muz/transforms/dl_mk_magic_sets.h.node
	@echo src/muz/transforms/dl_mk_magic_sets.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_magic_sets$(OBJ_EXT) ../src/muz/transforms/dl_mk_magic_sets.cpp
muz/transforms/dl_mk_magic_symbolic$(OBJ_EXT): ../src/muz/transforms/dl_mk_magic_symbolic.cpp muz/transforms/dl_mk_magic_symbolic.h.node muz/base/dl_context.h.node
	@echo src/muz/transforms/dl_mk_magic_symbolic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_magic_symbolic$(OBJ_EXT) ../src/muz/transforms/dl_mk_magic_symbolic.cpp
muz/transforms/dl_mk_separate_negated_tails.h.node: ../src/muz/transforms/dl_mk_separate_negated_tails.h muz/base/dl_rule_transformer.h.node muz/base/dl_context.h.node
	@echo done > muz/transforms/dl_mk_separate_negated_tails.h.node
muz/transforms/dl_mk_separate_negated_tails$(OBJ_EXT): ../src/muz/transforms/dl_mk_separate_negated_tails.cpp muz/transforms/dl_mk_separate_negated_tails.h.node muz/base/dl_context.h.node
	@echo src/muz/transforms/dl_mk_separate_negated_tails.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_separate_negated_tails$(OBJ_EXT) ../src/muz/transforms/dl_mk_separate_negated_tails.cpp
muz/transforms/dl_mk_unfold.h.node: ../src/muz/transforms/dl_mk_unfold.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node util/uint_set.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_rule_inliner.h.node
	@echo done > muz/transforms/dl_mk_unfold.h.node
muz/transforms/dl_mk_unfold$(OBJ_EXT): ../src/muz/transforms/dl_mk_unfold.cpp muz/transforms/dl_mk_unfold.h.node
	@echo src/muz/transforms/dl_mk_unfold.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_unfold$(OBJ_EXT) ../src/muz/transforms/dl_mk_unfold.cpp
muz/transforms/dl_mk_loop_counter$(OBJ_EXT): ../src/muz/transforms/dl_mk_loop_counter.cpp muz/transforms/dl_mk_loop_counter.h.node muz/base/dl_context.h.node
	@echo src/muz/transforms/dl_mk_loop_counter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_loop_counter$(OBJ_EXT) ../src/muz/transforms/dl_mk_loop_counter.cpp
muz/transforms/dl_mk_array_blast$(OBJ_EXT): ../src/muz/transforms/dl_mk_array_blast.cpp muz/transforms/dl_mk_array_blast.h.node ast/ast_util.h.node ast/scoped_proof.h.node
	@echo src/muz/transforms/dl_mk_array_blast.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_array_blast$(OBJ_EXT) ../src/muz/transforms/dl_mk_array_blast.cpp
muz/transforms/dl_mk_synchronize.h.node: ../src/muz/transforms/dl_mk_synchronize.h muz/base/dl_context.h.node muz/base/dl_rule_set.h.node util/uint_set.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_rule_inliner.h.node
	@echo done > muz/transforms/dl_mk_synchronize.h.node
muz/transforms/dl_mk_synchronize$(OBJ_EXT): ../src/muz/transforms/dl_mk_synchronize.cpp muz/transforms/dl_mk_synchronize.h.node
	@echo src/muz/transforms/dl_mk_synchronize.cpp
	@$(CXX) $(CXXFLAGS) $(includes_51) $(CXX_OUT_FLAG)muz/transforms/dl_mk_synchronize$(OBJ_EXT) ../src/muz/transforms/dl_mk_synchronize.cpp
muz/transforms/transforms$(LIB_EXT): muz/transforms/dl_mk_bit_blast$(OBJ_EXT) muz/transforms/dl_mk_rule_inliner$(OBJ_EXT) muz/transforms/dl_mk_coalesce$(OBJ_EXT) muz/transforms/dl_mk_scale$(OBJ_EXT) muz/transforms/dl_mk_filter_rules$(OBJ_EXT) muz/transforms/dl_mk_interp_tail_simplifier$(OBJ_EXT) muz/transforms/dl_mk_quantifier_abstraction$(OBJ_EXT) muz/transforms/dl_transforms$(OBJ_EXT) muz/transforms/dl_mk_karr_invariants$(OBJ_EXT) muz/transforms/dl_mk_backwards$(OBJ_EXT) muz/transforms/dl_mk_array_instantiation$(OBJ_EXT) muz/transforms/dl_mk_array_eq_rewrite$(OBJ_EXT) muz/transforms/dl_mk_subsumption_checker$(OBJ_EXT) muz/transforms/dl_mk_elim_term_ite$(OBJ_EXT) muz/transforms/dl_mk_unbound_compressor$(OBJ_EXT) muz/transforms/dl_mk_quantifier_instantiation$(OBJ_EXT) muz/transforms/dl_mk_slice$(OBJ_EXT) muz/transforms/dl_mk_coi_filter$(OBJ_EXT) muz/transforms/dl_mk_magic_sets$(OBJ_EXT) muz/transforms/dl_mk_magic_symbolic$(OBJ_EXT) muz/transforms/dl_mk_separate_negated_tails$(OBJ_EXT) muz/transforms/dl_mk_unfold$(OBJ_EXT) muz/transforms/dl_mk_loop_counter$(OBJ_EXT) muz/transforms/dl_mk_array_blast$(OBJ_EXT) muz/transforms/dl_mk_synchronize$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/transforms/transforms$(LIB_EXT) muz/transforms/dl_mk_bit_blast$(OBJ_EXT) muz/transforms/dl_mk_rule_inliner$(OBJ_EXT) muz/transforms/dl_mk_coalesce$(OBJ_EXT) muz/transforms/dl_mk_scale$(OBJ_EXT) muz/transforms/dl_mk_filter_rules$(OBJ_EXT) muz/transforms/dl_mk_interp_tail_simplifier$(OBJ_EXT) muz/transforms/dl_mk_quantifier_abstraction$(OBJ_EXT) muz/transforms/dl_transforms$(OBJ_EXT) muz/transforms/dl_mk_karr_invariants$(OBJ_EXT) muz/transforms/dl_mk_backwards$(OBJ_EXT) muz/transforms/dl_mk_array_instantiation$(OBJ_EXT) muz/transforms/dl_mk_array_eq_rewrite$(OBJ_EXT) muz/transforms/dl_mk_subsumption_checker$(OBJ_EXT) muz/transforms/dl_mk_elim_term_ite$(OBJ_EXT) muz/transforms/dl_mk_unbound_compressor$(OBJ_EXT) muz/transforms/dl_mk_quantifier_instantiation$(OBJ_EXT) muz/transforms/dl_mk_slice$(OBJ_EXT) muz/transforms/dl_mk_coi_filter$(OBJ_EXT) muz/transforms/dl_mk_magic_sets$(OBJ_EXT) muz/transforms/dl_mk_magic_symbolic$(OBJ_EXT) muz/transforms/dl_mk_separate_negated_tails$(OBJ_EXT) muz/transforms/dl_mk_unfold$(OBJ_EXT) muz/transforms/dl_mk_loop_counter$(OBJ_EXT) muz/transforms/dl_mk_array_blast$(OBJ_EXT) muz/transforms/dl_mk_synchronize$(OBJ_EXT)
transforms: muz/transforms/transforms$(LIB_EXT)

includes_52 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src
muz/rel/tbv.h.node: ../src/muz/rel/tbv.h util/fixed_bit_vector.h.node util/rational.h.node ast/ast.h.node
	@echo done > muz/rel/tbv.h.node
muz/rel/doc.h.node: ../src/muz/rel/doc.h muz/rel/tbv.h.node util/union_find.h.node util/buffer.h.node
	@echo done > muz/rel/doc.h.node
muz/rel/dl_base.h.node: ../src/muz/rel/dl_base.h ast/ast.h.node util/map.h.node util/vector.h.node util/ref.h.node muz/base/dl_util.h.node muz/base/dl_context.h.node
	@echo done > muz/rel/dl_base.h.node
muz/rel/check_relation.h.node: ../src/muz/rel/check_relation.h muz/rel/doc.h.node muz/rel/dl_base.h.node
	@echo done > muz/rel/check_relation.h.node
muz/rel/dl_relation_manager.h.node: ../src/muz/rel/dl_relation_manager.h util/map.h.node util/vector.h.node muz/rel/dl_base.h.node
	@echo done > muz/rel/dl_relation_manager.h.node
muz/rel/check_relation$(OBJ_EXT): ../src/muz/rel/check_relation.cpp muz/rel/check_relation.h.node muz/rel/dl_relation_manager.h.node ast/ast_util.h.node smt/smt_kernel.h.node
	@echo src/muz/rel/check_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/check_relation$(OBJ_EXT) ../src/muz/rel/check_relation.cpp
muz/rel/dl_vector_relation.h.node: ../src/muz/rel/dl_vector_relation.h ast/ast_pp.h.node muz/base/dl_context.h.node util/union_find.h.node
	@echo done > muz/rel/dl_vector_relation.h.node
muz/rel/dl_interval_relation.h.node: ../src/muz/rel/dl_interval_relation.h ast/arith_decl_plugin.h.node smt/old_interval.h.node muz/base/dl_context.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_base.h.node muz/rel/dl_vector_relation.h.node
	@echo done > muz/rel/dl_interval_relation.h.node
muz/rel/dl_bound_relation.h.node: ../src/muz/rel/dl_bound_relation.h muz/base/dl_context.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_base.h.node util/uint_set.h.node muz/rel/dl_vector_relation.h.node muz/rel/dl_interval_relation.h.node ast/arith_decl_plugin.h.node ast/rewriter/bool_rewriter.h.node
	@echo done > muz/rel/dl_bound_relation.h.node
muz/rel/dl_bound_relation$(OBJ_EXT): ../src/muz/rel/dl_bound_relation.cpp muz/rel/dl_bound_relation.h.node util/debug.h.node ast/ast_pp.h.node
	@echo src/muz/rel/dl_bound_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_bound_relation$(OBJ_EXT) ../src/muz/rel/dl_bound_relation.cpp
muz/rel/dl_sieve_relation.h.node: ../src/muz/rel/dl_sieve_relation.h muz/base/dl_context.h.node muz/rel/dl_relation_manager.h.node
	@echo done > muz/rel/dl_sieve_relation.h.node
muz/rel/dl_table_relation.h.node: ../src/muz/rel/dl_table_relation.h muz/rel/dl_base.h.node muz/base/dl_util.h.node
	@echo done > muz/rel/dl_table_relation.h.node
muz/rel/dl_product_relation.h.node: ../src/muz/rel/dl_product_relation.h muz/base/dl_context.h.node muz/rel/dl_relation_manager.h.node
	@echo done > muz/rel/dl_product_relation.h.node
muz/rel/dl_product_relation$(OBJ_EXT): ../src/muz/rel/dl_product_relation.cpp muz/rel/dl_sieve_relation.h.node muz/rel/dl_table_relation.h.node muz/rel/dl_product_relation.h.node ast/rewriter/bool_rewriter.h.node ast/ast_pp.h.node
	@echo src/muz/rel/dl_product_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_product_relation$(OBJ_EXT) ../src/muz/rel/dl_product_relation.cpp
muz/rel/dl_mk_similarity_compressor.h.node: ../src/muz/rel/dl_mk_similarity_compressor.h util/map.h.node util/obj_pair_hashtable.h.node muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/rel/dl_mk_similarity_compressor.h.node
muz/rel/dl_mk_similarity_compressor$(OBJ_EXT): ../src/muz/rel/dl_mk_similarity_compressor.cpp muz/rel/dl_mk_similarity_compressor.h.node muz/rel/dl_relation_manager.h.node
	@echo src/muz/rel/dl_mk_similarity_compressor.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_mk_similarity_compressor$(OBJ_EXT) ../src/muz/rel/dl_mk_similarity_compressor.cpp
muz/rel/karr_relation.h.node: ../src/muz/rel/karr_relation.h muz/transforms/dl_mk_karr_invariants.h.node muz/rel/dl_relation_manager.h.node
	@echo done > muz/rel/karr_relation.h.node
muz/rel/karr_relation$(OBJ_EXT): ../src/muz/rel/karr_relation.cpp muz/rel/karr_relation.h.node ast/rewriter/bool_rewriter.h.node ast/ast_util.h.node
	@echo src/muz/rel/karr_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/karr_relation$(OBJ_EXT) ../src/muz/rel/karr_relation.cpp
muz/rel/dl_lazy_table.h.node: ../src/muz/rel/dl_lazy_table.h muz/rel/dl_base.h.node util/ref.h.node
	@echo done > muz/rel/dl_lazy_table.h.node
muz/rel/dl_lazy_table$(OBJ_EXT): ../src/muz/rel/dl_lazy_table.cpp muz/rel/dl_lazy_table.h.node muz/rel/dl_relation_manager.h.node
	@echo src/muz/rel/dl_lazy_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_lazy_table$(OBJ_EXT) ../src/muz/rel/dl_lazy_table.cpp
muz/rel/dl_base$(OBJ_EXT): ../src/muz/rel/dl_base.cpp ast/ast_pp.h.node util/union_find.h.node util/vector.h.node muz/base/dl_context.h.node muz/rel/dl_base.h.node ast/rewriter/bool_rewriter.h.node muz/rel/dl_relation_manager.h.node
	@echo src/muz/rel/dl_base.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_base$(OBJ_EXT) ../src/muz/rel/dl_base.cpp
muz/rel/dl_external_relation.h.node: ../src/muz/rel/dl_external_relation.h muz/rel/dl_base.h.node
	@echo done > muz/rel/dl_external_relation.h.node
muz/rel/dl_external_relation$(OBJ_EXT): ../src/muz/rel/dl_external_relation.cpp util/debug.h.node ast/ast_pp.h.node muz/base/dl_context.h.node muz/rel/dl_external_relation.h.node ast/dl_decl_plugin.h.node
	@echo src/muz/rel/dl_external_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_external_relation$(OBJ_EXT) ../src/muz/rel/dl_external_relation.cpp
muz/rel/dl_sparse_table.h.node: ../src/muz/rel/dl_sparse_table.h ast/ast.h.node util/bit_vector.h.node util/buffer.h.node util/hashtable.h.node util/map.h.node util/ref_vector.h.node util/vector.h.node muz/rel/dl_base.h.node
	@echo done > muz/rel/dl_sparse_table.h.node
muz/rel/dl_sparse_table$(OBJ_EXT): ../src/muz/rel/dl_sparse_table.cpp muz/base/dl_context.h.node muz/base/dl_util.h.node muz/rel/dl_sparse_table.h.node
	@echo src/muz/rel/dl_sparse_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_sparse_table$(OBJ_EXT) ../src/muz/rel/dl_sparse_table.cpp
muz/rel/dl_mk_simple_joins.h.node: ../src/muz/rel/dl_mk_simple_joins.h util/map.h.node util/obj_pair_hashtable.h.node muz/base/dl_context.h.node muz/base/dl_rule_set.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/rel/dl_mk_simple_joins.h.node
muz/rel/dl_mk_simple_joins$(OBJ_EXT): ../src/muz/rel/dl_mk_simple_joins.cpp ast/ast_pp.h.node util/trace.h.node muz/rel/dl_mk_simple_joins.h.node muz/rel/dl_relation_manager.h.node
	@echo src/muz/rel/dl_mk_simple_joins.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_mk_simple_joins$(OBJ_EXT) ../src/muz/rel/dl_mk_simple_joins.cpp
muz/rel/dl_instruction.h.node: ../src/muz/rel/dl_instruction.h ast/ast.h.node util/vector.h.node muz/rel/dl_base.h.node muz/base/dl_costs.h.node muz/base/dl_context.h.node
	@echo done > muz/rel/dl_instruction.h.node
muz/rel/rel_context.h.node: ../src/muz/rel/rel_context.h ast/ast.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_instruction.h.node muz/base/dl_engine_base.h.node muz/base/dl_context.h.node util/lbool.h.node
	@echo done > muz/rel/rel_context.h.node
muz/rel/dl_compiler.h.node: ../src/muz/rel/dl_compiler.h ast/ast.h.node util/hashtable.h.node util/map.h.node util/obj_pair_hashtable.h.node util/ref_vector.h.node util/vector.h.node muz/rel/dl_base.h.node muz/base/dl_context.h.node muz/rel/dl_instruction.h.node
	@echo done > muz/rel/dl_compiler.h.node
muz/rel/dl_compiler$(OBJ_EXT): ../src/muz/rel/dl_compiler.cpp util/ref_vector.h.node muz/base/dl_context.h.node muz/rel/rel_context.h.node muz/base/dl_rule.h.node muz/base/dl_util.h.node muz/rel/dl_compiler.h.node ast/ast_pp.h.node ast/ast_util.h.node
	@echo src/muz/rel/dl_compiler.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_compiler$(OBJ_EXT) ../src/muz/rel/dl_compiler.cpp
muz/rel/dl_sieve_relation$(OBJ_EXT): ../src/muz/rel/dl_sieve_relation.cpp ast/ast_pp.h.node muz/rel/dl_sieve_relation.h.node
	@echo src/muz/rel/dl_sieve_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_sieve_relation$(OBJ_EXT) ../src/muz/rel/dl_sieve_relation.cpp
muz/rel/dl_mk_explanations.h.node: ../src/muz/rel/dl_mk_explanations.h muz/base/dl_context.h.node muz/base/dl_rule_transformer.h.node
	@echo done > muz/rel/dl_mk_explanations.h.node
muz/rel/dl_finite_product_relation.h.node: ../src/muz/rel/dl_finite_product_relation.h muz/rel/dl_base.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_table_relation.h.node
	@echo done > muz/rel/dl_finite_product_relation.h.node
muz/rel/udoc_relation.h.node: ../src/muz/rel/udoc_relation.h muz/rel/doc.h.node muz/rel/dl_base.h.node
	@echo done > muz/rel/udoc_relation.h.node
muz/rel/dl_table.h.node: ../src/muz/rel/dl_table.h ast/ast.h.node util/bit_vector.h.node util/buffer.h.node util/hashtable.h.node util/map.h.node util/ref_vector.h.node util/vector.h.node util/union_find.h.node muz/rel/dl_base.h.node muz/base/dl_util.h.node util/bit_vector.h.node
	@echo done > muz/rel/dl_table.h.node
muz/rel/aig_exporter.h.node: ../src/muz/rel/aig_exporter.h tactic/aig/aig.h.node muz/base/dl_rule_set.h.node muz/rel/rel_context.h.node
	@echo done > muz/rel/aig_exporter.h.node
muz/rel/rel_context$(OBJ_EXT): ../src/muz/rel/rel_context.cpp muz/rel/rel_context.h.node util/stopwatch.h.node muz/base/dl_context.h.node muz/rel/dl_compiler.h.node muz/rel/dl_instruction.h.node muz/rel/dl_mk_explanations.h.node muz/transforms/dl_mk_magic_sets.h.node muz/rel/dl_product_relation.h.node muz/rel/dl_bound_relation.h.node muz/rel/dl_interval_relation.h.node muz/rel/karr_relation.h.node muz/rel/dl_finite_product_relation.h.node muz/rel/udoc_relation.h.node muz/rel/check_relation.h.node muz/rel/dl_lazy_table.h.node muz/rel/dl_sparse_table.h.node muz/rel/dl_table.h.node muz/rel/dl_table_relation.h.node muz/rel/aig_exporter.h.node muz/rel/dl_mk_simple_joins.h.node muz/rel/dl_mk_similarity_compressor.h.node muz/transforms/dl_mk_unbound_compressor.h.node muz/transforms/dl_mk_subsumption_checker.h.node muz/transforms/dl_mk_coi_filter.h.node muz/transforms/dl_mk_filter_rules.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node muz/transforms/dl_mk_bit_blast.h.node muz/transforms/dl_mk_separate_negated_tails.h.node ast/ast_util.h.node
	@echo src/muz/rel/rel_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/rel_context$(OBJ_EXT) ../src/muz/rel/rel_context.cpp
muz/rel/dl_interval_relation$(OBJ_EXT): ../src/muz/rel/dl_interval_relation.cpp util/debug.h.node ast/ast_pp.h.node muz/rel/dl_interval_relation.h.node muz/rel/dl_relation_manager.h.node ast/rewriter/bool_rewriter.h.node
	@echo src/muz/rel/dl_interval_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_interval_relation$(OBJ_EXT) ../src/muz/rel/dl_interval_relation.cpp
muz/rel/dl_check_table.h.node: ../src/muz/rel/dl_check_table.h muz/rel/dl_base.h.node ast/dl_decl_plugin.h.node muz/rel/dl_relation_manager.h.node
	@echo done > muz/rel/dl_check_table.h.node
muz/rel/dl_relation_manager$(OBJ_EXT): ../src/muz/rel/dl_relation_manager.cpp ast/ast_pp.h.node muz/rel/dl_check_table.h.node muz/base/dl_context.h.node muz/rel/dl_finite_product_relation.h.node muz/rel/dl_product_relation.h.node muz/rel/dl_sieve_relation.h.node muz/rel/dl_table_relation.h.node muz/rel/dl_relation_manager.h.node
	@echo src/muz/rel/dl_relation_manager.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_relation_manager$(OBJ_EXT) ../src/muz/rel/dl_relation_manager.cpp
muz/rel/tbv$(OBJ_EXT): ../src/muz/rel/tbv.cpp muz/rel/tbv.h.node util/hashtable.h.node ast/ast_util.h.node
	@echo src/muz/rel/tbv.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/tbv$(OBJ_EXT) ../src/muz/rel/tbv.cpp
muz/rel/doc$(OBJ_EXT): ../src/muz/rel/doc.cpp muz/rel/doc.h.node smt/smt_kernel.h.node ast/rewriter/expr_safe_replace.h.node smt/params/smt_params.h.node ast/ast_util.h.node ast/ast_pp.h.node
	@echo src/muz/rel/doc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/doc$(OBJ_EXT) ../src/muz/rel/doc.cpp
muz/rel/dl_finite_product_relation$(OBJ_EXT): ../src/muz/rel/dl_finite_product_relation.cpp muz/base/dl_context.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_table_relation.h.node muz/rel/dl_finite_product_relation.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node
	@echo src/muz/rel/dl_finite_product_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_finite_product_relation$(OBJ_EXT) ../src/muz/rel/dl_finite_product_relation.cpp
muz/rel/dl_table$(OBJ_EXT): ../src/muz/rel/dl_table.cpp muz/base/dl_context.h.node muz/base/dl_util.h.node muz/rel/dl_table.h.node muz/rel/dl_relation_manager.h.node
	@echo src/muz/rel/dl_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_table$(OBJ_EXT) ../src/muz/rel/dl_table.cpp
muz/rel/dl_mk_explanations$(OBJ_EXT): ../src/muz/rel/dl_mk_explanations.cpp ast/ast_pp.h.node ast/ast_smt_pp.h.node muz/rel/dl_finite_product_relation.h.node muz/rel/dl_product_relation.h.node muz/rel/dl_sieve_relation.h.node muz/rel/dl_mk_explanations.h.node
	@echo src/muz/rel/dl_mk_explanations.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_mk_explanations$(OBJ_EXT) ../src/muz/rel/dl_mk_explanations.cpp
muz/rel/udoc_relation$(OBJ_EXT): ../src/muz/rel/udoc_relation.cpp muz/rel/udoc_relation.h.node muz/rel/dl_relation_manager.h.node ast/ast_util.h.node smt/smt_kernel.h.node
	@echo src/muz/rel/udoc_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/udoc_relation$(OBJ_EXT) ../src/muz/rel/udoc_relation.cpp
muz/rel/dl_check_table$(OBJ_EXT): ../src/muz/rel/dl_check_table.cpp muz/rel/dl_check_table.h.node muz/rel/dl_table.h.node
	@echo src/muz/rel/dl_check_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_check_table$(OBJ_EXT) ../src/muz/rel/dl_check_table.cpp
muz/rel/dl_table_relation$(OBJ_EXT): ../src/muz/rel/dl_table_relation.cpp muz/base/dl_context.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_table_relation.h.node
	@echo src/muz/rel/dl_table_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_table_relation$(OBJ_EXT) ../src/muz/rel/dl_table_relation.cpp
muz/rel/aig_exporter$(OBJ_EXT): ../src/muz/rel/aig_exporter.cpp muz/rel/aig_exporter.h.node muz/base/dl_context.h.node
	@echo src/muz/rel/aig_exporter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/aig_exporter$(OBJ_EXT) ../src/muz/rel/aig_exporter.cpp
muz/rel/dl_instruction$(OBJ_EXT): ../src/muz/rel/dl_instruction.cpp ast/ast_pp.h.node util/stopwatch.h.node muz/base/dl_context.h.node muz/base/dl_util.h.node muz/rel/dl_instruction.h.node muz/rel/rel_context.h.node util/debug.h.node util/warning.h.node
	@echo src/muz/rel/dl_instruction.cpp
	@$(CXX) $(CXXFLAGS) $(includes_52) $(CXX_OUT_FLAG)muz/rel/dl_instruction$(OBJ_EXT) ../src/muz/rel/dl_instruction.cpp
muz/rel/rel$(LIB_EXT): muz/rel/check_relation$(OBJ_EXT) muz/rel/dl_bound_relation$(OBJ_EXT) muz/rel/dl_product_relation$(OBJ_EXT) muz/rel/dl_mk_similarity_compressor$(OBJ_EXT) muz/rel/karr_relation$(OBJ_EXT) muz/rel/dl_lazy_table$(OBJ_EXT) muz/rel/dl_base$(OBJ_EXT) muz/rel/dl_external_relation$(OBJ_EXT) muz/rel/dl_sparse_table$(OBJ_EXT) muz/rel/dl_mk_simple_joins$(OBJ_EXT) muz/rel/dl_compiler$(OBJ_EXT) muz/rel/dl_sieve_relation$(OBJ_EXT) muz/rel/rel_context$(OBJ_EXT) muz/rel/dl_interval_relation$(OBJ_EXT) muz/rel/dl_relation_manager$(OBJ_EXT) muz/rel/tbv$(OBJ_EXT) muz/rel/doc$(OBJ_EXT) muz/rel/dl_finite_product_relation$(OBJ_EXT) muz/rel/dl_table$(OBJ_EXT) muz/rel/dl_mk_explanations$(OBJ_EXT) muz/rel/udoc_relation$(OBJ_EXT) muz/rel/dl_check_table$(OBJ_EXT) muz/rel/dl_table_relation$(OBJ_EXT) muz/rel/aig_exporter$(OBJ_EXT) muz/rel/dl_instruction$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/rel/rel$(LIB_EXT) muz/rel/check_relation$(OBJ_EXT) muz/rel/dl_bound_relation$(OBJ_EXT) muz/rel/dl_product_relation$(OBJ_EXT) muz/rel/dl_mk_similarity_compressor$(OBJ_EXT) muz/rel/karr_relation$(OBJ_EXT) muz/rel/dl_lazy_table$(OBJ_EXT) muz/rel/dl_base$(OBJ_EXT) muz/rel/dl_external_relation$(OBJ_EXT) muz/rel/dl_sparse_table$(OBJ_EXT) muz/rel/dl_mk_simple_joins$(OBJ_EXT) muz/rel/dl_compiler$(OBJ_EXT) muz/rel/dl_sieve_relation$(OBJ_EXT) muz/rel/rel_context$(OBJ_EXT) muz/rel/dl_interval_relation$(OBJ_EXT) muz/rel/dl_relation_manager$(OBJ_EXT) muz/rel/tbv$(OBJ_EXT) muz/rel/doc$(OBJ_EXT) muz/rel/dl_finite_product_relation$(OBJ_EXT) muz/rel/dl_table$(OBJ_EXT) muz/rel/dl_mk_explanations$(OBJ_EXT) muz/rel/udoc_relation$(OBJ_EXT) muz/rel/dl_check_table$(OBJ_EXT) muz/rel/dl_table_relation$(OBJ_EXT) muz/rel/aig_exporter$(OBJ_EXT) muz/rel/dl_instruction$(OBJ_EXT)
rel: muz/rel/rel$(LIB_EXT)

includes_53 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/smt/tactic -I../src
muz/spacer/spacer_antiunify.h.node: ../src/muz/spacer/spacer_antiunify.h ast/ast.h.node ast/substitution/substitution.h.node util/obj_pair_hashtable.h.node
	@echo done > muz/spacer/spacer_antiunify.h.node
muz/spacer/spacer_util.h.node: ../src/muz/spacer/spacer_util.h ast/ast.h.node ast/ast_pp.h.node util/obj_hashtable.h.node util/ref_vector.h.node util/trace.h.node util/vector.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/ast_util.h.node ast/expr_map.h.node model/model.h.node util/stopwatch.h.node muz/spacer/spacer_antiunify.h.node
	@echo done > muz/spacer/spacer_util.h.node
muz/spacer/spacer_sym_mux.h.node: ../src/muz/spacer/spacer_sym_mux.h ast/ast.h.node util/map.h.node util/vector.h.node
	@echo done > muz/spacer/spacer_sym_mux.h.node
muz/spacer/spacer_farkas_learner.h.node: ../src/muz/spacer/spacer_farkas_learner.h ast/ast.h.node
	@echo done > muz/spacer/spacer_farkas_learner.h.node
muz/spacer/spacer_manager.h.node: ../src/muz/spacer/spacer_manager.h ast/rewriter/bool_rewriter.h.node ast/rewriter/expr_replacer.h.node ast/expr_substitution.h.node util/map.h.node util/ref_vector.h.node smt/smt_kernel.h.node muz/spacer/spacer_util.h.node muz/spacer/spacer_sym_mux.h.node muz/spacer/spacer_farkas_learner.h.node muz/base/dl_rule.h.node solver/solver.h.node solver/solver_pool.h.node
	@echo done > muz/spacer/spacer_manager.h.node
muz/spacer/spacer_iuc_solver.h.node: ../src/muz/spacer/spacer_iuc_solver.h solver/solver.h.node ast/expr_substitution.h.node util/stopwatch.h.node
	@echo done > muz/spacer/spacer_iuc_solver.h.node
muz/spacer/spacer_prop_solver.h.node: ../src/muz/spacer/spacer_prop_solver.h ast/ast.h.node util/obj_hashtable.h.node smt/smt_kernel.h.node util/util.h.node util/vector.h.node solver/solver.h.node muz/spacer/spacer_iuc_solver.h.node muz/spacer/spacer_util.h.node
	@echo done > muz/spacer/spacer_prop_solver.h.node
muz/spacer/spacer_json.h.node: ../src/muz/spacer/spacer_json.h util/ref.h.node util/ref_vector.h.node
	@echo done > muz/spacer/spacer_json.h.node
muz/spacer/spacer_legacy_frames.h.node: ../src/muz/spacer/spacer_legacy_frames.h
	@echo done > muz/spacer/spacer_legacy_frames.h.node
muz/spacer/spacer_context.h.node: ../src/muz/spacer/spacer_context.h util/scoped_ptr_vector.h.node muz/spacer/spacer_manager.h.node muz/spacer/spacer_prop_solver.h.node muz/spacer/spacer_json.h.node muz/base/fp_params.hpp.node muz/spacer/spacer_legacy_frames.h.node
	@echo done > muz/spacer/spacer_context.h.node
muz/spacer/spacer_generalizers.h.node: ../src/muz/spacer/spacer_generalizers.h ast/arith_decl_plugin.h.node muz/spacer/spacer_context.h.node
	@echo done > muz/spacer/spacer_generalizers.h.node
muz/spacer/spacer_qe_project.h.node: ../src/muz/spacer/spacer_qe_project.h model/model.h.node ast/expr_map.h.node
	@echo done > muz/spacer/spacer_qe_project.h.node
muz/spacer/spacer_legacy_frames$(OBJ_EXT): ../src/muz/spacer/spacer_legacy_frames.cpp muz/spacer/spacer_context.h.node muz/base/dl_util.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node util/util.h.node muz/spacer/spacer_prop_solver.h.node muz/spacer/spacer_context.h.node muz/spacer/spacer_generalizers.h.node ast/for_each_expr.h.node muz/base/dl_rule_set.h.node smt/tactic/unit_subsumption_tactic.h.node model/model_smt2_pp.h.node muz/transforms/dl_mk_rule_inliner.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node ast/ast_util.h.node ast/proofs/proof_checker.h.node smt/smt_value_sort.h.node ast/proofs/proof_utils.h.node ast/scoped_proof.h.node muz/spacer/spacer_qe_project.h.node tactic/core/blast_term_ite_tactic.h.node util/timeit.h.node util/luby.h.node ast/rewriter/expr_safe_replace.h.node ast/expr_abstract.h.node ast/rewriter/factor_equivs.h.node
	@echo src/muz/spacer/spacer_legacy_frames.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_legacy_frames$(OBJ_EXT) ../src/muz/spacer/spacer_legacy_frames.cpp
muz/spacer/spacer_matrix.h.node: ../src/muz/spacer/spacer_matrix.h util/rational.h.node util/vector.h.node
	@echo done > muz/spacer/spacer_matrix.h.node
muz/spacer/spacer_matrix$(OBJ_EXT): ../src/muz/spacer/spacer_matrix.cpp muz/spacer/spacer_matrix.h.node
	@echo src/muz/spacer/spacer_matrix.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_matrix$(OBJ_EXT) ../src/muz/spacer/spacer_matrix.cpp
muz/spacer/spacer_iuc_proof.h.node: ../src/muz/spacer/spacer_iuc_proof.h ast/ast.h.node
	@echo done > muz/spacer/spacer_iuc_proof.h.node
muz/spacer/spacer_proof_utils.h.node: ../src/muz/spacer/spacer_proof_utils.h ast/ast.h.node
	@echo done > muz/spacer/spacer_proof_utils.h.node
muz/spacer/spacer_proof_utils$(OBJ_EXT): ../src/muz/spacer/spacer_proof_utils.cpp util/params.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/proofs/proof_checker.h.node muz/base/dl_util.h.node muz/spacer/spacer_iuc_proof.h.node ast/proofs/proof_utils.h.node muz/spacer/spacer_proof_utils.h.node muz/spacer/spacer_util.h.node
	@echo src/muz/spacer/spacer_proof_utils.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_proof_utils$(OBJ_EXT) ../src/muz/spacer/spacer_proof_utils.cpp
muz/spacer/spacer_legacy_mev.h.node: ../src/muz/spacer/spacer_legacy_mev.h ast/ast.h.node ast/ast_pp.h.node util/obj_hashtable.h.node util/ref_vector.h.node util/trace.h.node util/vector.h.node ast/arith_decl_plugin.h.node ast/array_decl_plugin.h.node ast/bv_decl_plugin.h.node
	@echo done > muz/spacer/spacer_legacy_mev.h.node
muz/spacer/spacer_legacy_mbp$(OBJ_EXT): ../src/muz/spacer/spacer_legacy_mbp.cpp ast/array_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node muz/base/dl_util.h.node ast/for_each_expr.h.node smt/params/smt_params.h.node model/model.h.node util/ref_vector.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node util/util.h.node muz/spacer/spacer_manager.h.node muz/spacer/spacer_util.h.node ast/arith_decl_plugin.h.node ast/rewriter/expr_replacer.h.node model/model_smt2_pp.h.node ast/scoped_proof.h.node qe/qe_lite.h.node muz/spacer/spacer_qe_project.h.node model/model_pp.h.node ast/rewriter/expr_safe_replace.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node muz/spacer/spacer_legacy_mev.h.node
	@echo src/muz/spacer/spacer_legacy_mbp.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_legacy_mbp$(OBJ_EXT) ../src/muz/spacer/spacer_legacy_mbp.cpp
muz/spacer/spacer_json$(OBJ_EXT): ../src/muz/spacer/spacer_json.cpp muz/spacer/spacer_context.h.node muz/spacer/spacer_json.h.node muz/spacer/spacer_util.h.node
	@echo src/muz/spacer/spacer_json.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_json$(OBJ_EXT) ../src/muz/spacer/spacer_json.cpp
muz/spacer/spacer_dl_interface.h.node: ../src/muz/spacer/spacer_dl_interface.h util/lbool.h.node muz/base/dl_rule.h.node muz/base/dl_rule_set.h.node muz/base/dl_engine_base.h.node util/statistics.h.node
	@echo done > muz/spacer/spacer_dl_interface.h.node
muz/spacer/spacer_callback.h.node: ../src/muz/spacer/spacer_callback.h muz/spacer/spacer_context.h.node muz/base/dl_engine_base.h.node
	@echo done > muz/spacer/spacer_callback.h.node
muz/spacer/spacer_dl_interface$(OBJ_EXT): ../src/muz/spacer/spacer_dl_interface.cpp muz/base/dl_context.h.node muz/transforms/dl_mk_coi_filter.h.node muz/transforms/dl_mk_interp_tail_simplifier.h.node muz/transforms/dl_mk_subsumption_checker.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/base/dl_rule.h.node muz/base/dl_rule_transformer.h.node parsers/smt2/smt2parser.h.node muz/spacer/spacer_context.h.node muz/spacer/spacer_dl_interface.h.node muz/base/dl_rule_set.h.node muz/transforms/dl_mk_slice.h.node muz/transforms/dl_mk_unfold.h.node muz/transforms/dl_mk_coalesce.h.node model/model_smt2_pp.h.node ast/scoped_proof.h.node muz/transforms/dl_transforms.h.node muz/spacer/spacer_callback.h.node
	@echo src/muz/spacer/spacer_dl_interface.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_dl_interface$(OBJ_EXT) ../src/muz/spacer/spacer_dl_interface.cpp
muz/spacer/spacer_iuc_proof$(OBJ_EXT): ../src/muz/spacer/spacer_iuc_proof.cpp ast/ast_pp_dot.h.node muz/spacer/spacer_iuc_proof.h.node ast/for_each_expr.h.node ast/array_decl_plugin.h.node ast/proofs/proof_utils.h.node muz/spacer/spacer_proof_utils.h.node muz/spacer/spacer_util.h.node
	@echo src/muz/spacer/spacer_iuc_proof.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_iuc_proof$(OBJ_EXT) ../src/muz/spacer/spacer_iuc_proof.cpp
tactic/arith/propagate_ineqs_tactic.h.node: ../src/tactic/arith/propagate_ineqs_tactic.h util/params.h.node
	@echo done > tactic/arith/propagate_ineqs_tactic.h.node
muz/spacer/spacer_util$(OBJ_EXT): ../src/muz/spacer/spacer_util.cpp util/util.h.node ast/ast.h.node ast/occurs.h.node ast/ast_pp.h.node ast/scoped_proof.h.node ast/for_each_expr.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/expr_safe_replace.h.node ast/array_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/factor_equivs.h.node ast/rewriter/expr_replacer.h.node smt/params/smt_params.h.node model/model.h.node model/model_evaluator.h.node model/model_smt2_pp.h.node model/model_pp.h.node qe/qe_lite.h.node qe/qe_mbp.h.node qe/mbp/mbp_term_graph.h.node qe/mbp/mbp_plugin.h.node tactic/tactical.h.node tactic/core/propagate_values_tactic.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/arith/arith_bounds_tactic.h.node muz/base/dl_util.h.node muz/spacer/spacer_legacy_mev.h.node muz/spacer/spacer_qe_project.h.node muz/spacer/spacer_manager.h.node muz/spacer/spacer_util.h.node
	@echo src/muz/spacer/spacer_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_util$(OBJ_EXT) ../src/muz/spacer/spacer_util.cpp
muz/spacer/spacer_unsat_core_learner.h.node: ../src/muz/spacer/spacer_unsat_core_learner.h ast/ast.h.node muz/spacer/spacer_util.h.node muz/spacer/spacer_proof_utils.h.node muz/spacer/spacer_iuc_proof.h.node
	@echo done > muz/spacer/spacer_unsat_core_learner.h.node
muz/spacer/spacer_unsat_core_plugin.h.node: ../src/muz/spacer/spacer_unsat_core_plugin.h ast/ast.h.node util/min_cut.h.node
	@echo done > muz/spacer/spacer_unsat_core_plugin.h.node
muz/spacer/spacer_unsat_core_learner$(OBJ_EXT): ../src/muz/spacer/spacer_unsat_core_learner.cpp ast/for_each_expr.h.node ast/proofs/proof_utils.h.node muz/spacer/spacer_unsat_core_learner.h.node muz/spacer/spacer_unsat_core_plugin.h.node muz/spacer/spacer_iuc_proof.h.node muz/spacer/spacer_util.h.node
	@echo src/muz/spacer/spacer_unsat_core_learner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_unsat_core_learner$(OBJ_EXT) ../src/muz/spacer/spacer_unsat_core_learner.cpp
muz/spacer/spacer_legacy_mev$(OBJ_EXT): ../src/muz/spacer/spacer_legacy_mev.cpp ast/array_decl_plugin.h.node ast/ast_pp.h.node ast/rewriter/bool_rewriter.h.node muz/base/dl_util.h.node ast/for_each_expr.h.node smt/params/smt_params.h.node model/model.h.node util/ref_vector.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node util/util.h.node muz/spacer/spacer_manager.h.node muz/spacer/spacer_legacy_mev.h.node muz/spacer/spacer_util.h.node ast/arith_decl_plugin.h.node ast/rewriter/expr_replacer.h.node model/model_smt2_pp.h.node ast/scoped_proof.h.node qe/qe_lite.h.node muz/spacer/spacer_qe_project.h.node model/model_pp.h.node ast/rewriter/expr_safe_replace.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node
	@echo src/muz/spacer/spacer_legacy_mev.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_legacy_mev$(OBJ_EXT) ../src/muz/spacer/spacer_legacy_mev.cpp
muz/spacer/spacer_mbc.h.node: ../src/muz/spacer/spacer_mbc.h ast/ast.h.node util/obj_hashtable.h.node model/model.h.node
	@echo done > muz/spacer/spacer_mbc.h.node
muz/spacer/spacer_mbc$(OBJ_EXT): ../src/muz/spacer/spacer_mbc.cpp muz/spacer/spacer_mbc.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/th_rewriter.h.node ast/scoped_proof.h.node model/model_evaluator.h.node
	@echo src/muz/spacer/spacer_mbc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_mbc$(OBJ_EXT) ../src/muz/spacer/spacer_mbc.cpp
muz/spacer/spacer_iuc_solver$(OBJ_EXT): ../src/muz/spacer/spacer_iuc_solver.cpp muz/spacer/spacer_iuc_solver.h.node ast/ast.h.node muz/spacer/spacer_util.h.node ast/proofs/proof_utils.h.node muz/spacer/spacer_farkas_learner.h.node ast/rewriter/expr_replacer.h.node muz/spacer/spacer_unsat_core_learner.h.node muz/spacer/spacer_unsat_core_plugin.h.node muz/spacer/spacer_iuc_proof.h.node
	@echo src/muz/spacer/spacer_iuc_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_iuc_solver$(OBJ_EXT) ../src/muz/spacer/spacer_iuc_solver.cpp
muz/spacer/spacer_manager$(OBJ_EXT): ../src/muz/spacer/spacer_manager.cpp muz/spacer/spacer_manager.h.node ast/ast_smt2_pp.h.node ast/for_each_expr.h.node ast/has_free_vars.h.node ast/rewriter/expr_replacer.h.node ast/expr_abstract.h.node model/model2expr.h.node model/model_smt2_pp.h.node tactic/model_converter.h.node smt/smt_solver.h.node
	@echo src/muz/spacer/spacer_manager.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_manager$(OBJ_EXT) ../src/muz/spacer/spacer_manager.cpp
muz/spacer/spacer_generalizers$(OBJ_EXT): ../src/muz/spacer/spacer_generalizers.cpp muz/spacer/spacer_context.h.node muz/spacer/spacer_generalizers.h.node ast/ast_util.h.node ast/expr_abstract.h.node ast/rewriter/var_subst.h.node ast/for_each_expr.h.node ast/rewriter/factor_equivs.h.node ast/rewriter/expr_safe_replace.h.node ast/substitution/matcher.h.node ast/expr_functors.h.node smt/smt_solver.h.node qe/mbp/mbp_term_graph.h.node
	@echo src/muz/spacer/spacer_generalizers.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_generalizers$(OBJ_EXT) ../src/muz/spacer/spacer_generalizers.cpp
muz/spacer/spacer_callback$(OBJ_EXT): ../src/muz/spacer/spacer_callback.cpp muz/spacer/spacer_callback.h.node muz/spacer/spacer_context.h.node
	@echo src/muz/spacer/spacer_callback.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_callback$(OBJ_EXT) ../src/muz/spacer/spacer_callback.cpp
muz/spacer/spacer_sym_mux$(OBJ_EXT): ../src/muz/spacer/spacer_sym_mux.cpp ast/ast_pp.h.node ast/for_each_expr.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node muz/spacer/spacer_util.h.node muz/spacer/spacer_sym_mux.h.node
	@echo src/muz/spacer/spacer_sym_mux.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_sym_mux$(OBJ_EXT) ../src/muz/spacer/spacer_sym_mux.cpp
muz/spacer/spacer_arith_generalizers$(OBJ_EXT): ../src/muz/spacer/spacer_arith_generalizers.cpp ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node muz/spacer/spacer_generalizers.h.node smt/smt_solver.h.node
	@echo src/muz/spacer/spacer_arith_generalizers.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_arith_generalizers$(OBJ_EXT) ../src/muz/spacer/spacer_arith_generalizers.cpp
muz/spacer/spacer_mev_array.h.node: ../src/muz/spacer/spacer_mev_array.h ast/ast.h.node ast/rewriter/rewriter_types.h.node util/params.h.node ast/array_decl_plugin.h.node
	@echo done > muz/spacer/spacer_mev_array.h.node
muz/spacer/spacer_qe_project$(OBJ_EXT): ../src/muz/spacer/spacer_qe_project.cpp ast/arith_decl_plugin.h.node ast/ast_pp.h.node ast/expr_functors.h.node ast/expr_substitution.h.node ast/ast_util.h.node ast/is_variable_test.h.node ast/rewriter/expr_replacer.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/th_rewriter.h.node model/model_evaluator.h.node model/model_pp.h.node qe/qe.h.node qe/qe_lite.h.node muz/spacer/spacer_mev_array.h.node muz/spacer/spacer_qe_project.h.node
	@echo src/muz/spacer/spacer_qe_project.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_qe_project$(OBJ_EXT) ../src/muz/spacer/spacer_qe_project.cpp
muz/spacer/spacer_sat_answer.h.node: ../src/muz/spacer/spacer_sat_answer.h muz/spacer/spacer_context.h.node ast/ast.h.node util/obj_hashtable.h.node model/model.h.node solver/solver.h.node
	@echo done > muz/spacer/spacer_sat_answer.h.node
muz/spacer/spacer_sat_answer$(OBJ_EXT): ../src/muz/spacer/spacer_sat_answer.cpp muz/spacer/spacer_sat_answer.h.node muz/base/dl_context.h.node muz/base/dl_rule.h.node ast/scoped_proof.h.node smt/smt_solver.h.node
	@echo src/muz/spacer/spacer_sat_answer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_sat_answer$(OBJ_EXT) ../src/muz/spacer/spacer_sat_answer.cpp
muz/spacer/spacer_context$(OBJ_EXT): ../src/muz/spacer/spacer_context.cpp util/util.h.node util/timeit.h.node util/luby.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/rewriter/var_subst.h.node ast/ast_smt2_pp.h.node ast/ast_ll_pp.h.node ast/ast_util.h.node ast/proofs/proof_checker.h.node ast/for_each_expr.h.node ast/scoped_proof.h.node ast/rewriter/expr_safe_replace.h.node ast/expr_abstract.h.node model/model_smt2_pp.h.node tactic/core/blast_term_ite_tactic.h.node smt/tactic/unit_subsumption_tactic.h.node smt/smt_value_sort.h.node smt/smt_solver.h.node muz/base/dl_util.h.node muz/spacer/spacer_prop_solver.h.node muz/spacer/spacer_context.h.node muz/spacer/spacer_generalizers.h.node muz/base/dl_rule_set.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/spacer/spacer_qe_project.h.node muz/spacer/spacer_sat_answer.h.node
	@echo src/muz/spacer/spacer_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_context$(OBJ_EXT) ../src/muz/spacer/spacer_context.cpp
muz/spacer/spacer_sem_matcher.h.node: ../src/muz/spacer/spacer_sem_matcher.h ast/substitution/substitution.h.node ast/arith_decl_plugin.h.node util/hashtable.h.node
	@echo done > muz/spacer/spacer_sem_matcher.h.node
muz/spacer/spacer_quant_generalizer$(OBJ_EXT): ../src/muz/spacer/spacer_quant_generalizer.cpp muz/spacer/spacer_context.h.node muz/spacer/spacer_generalizers.h.node muz/spacer/spacer_manager.h.node ast/expr_abstract.h.node ast/rewriter/var_subst.h.node ast/for_each_expr.h.node ast/rewriter/factor_equivs.h.node ast/rewriter/expr_safe_replace.h.node ast/substitution/matcher.h.node ast/expr_functors.h.node muz/spacer/spacer_sem_matcher.h.node
	@echo src/muz/spacer/spacer_quant_generalizer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_quant_generalizer$(OBJ_EXT) ../src/muz/spacer/spacer_quant_generalizer.cpp
muz/spacer/spacer_unsat_core_plugin$(OBJ_EXT): ../src/muz/spacer/spacer_unsat_core_plugin.cpp ast/rewriter/bool_rewriter.h.node ast/arith_decl_plugin.h.node ast/proofs/proof_utils.h.node solver/solver.h.node smt/smt_farkas_util.h.node smt/smt_solver.h.node muz/spacer/spacer_matrix.h.node muz/spacer/spacer_unsat_core_plugin.h.node muz/spacer/spacer_unsat_core_learner.h.node muz/spacer/spacer_iuc_proof.h.node
	@echo src/muz/spacer/spacer_unsat_core_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_unsat_core_plugin$(OBJ_EXT) ../src/muz/spacer/spacer_unsat_core_plugin.cpp
muz/spacer/spacer_antiunify$(OBJ_EXT): ../src/muz/spacer/spacer_antiunify.cpp muz/spacer/spacer_antiunify.h.node ast/ast.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node ast/arith_decl_plugin.h.node ast/ast_util.h.node ast/expr_abstract.h.node
	@echo src/muz/spacer/spacer_antiunify.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_antiunify$(OBJ_EXT) ../src/muz/spacer/spacer_antiunify.cpp
muz/spacer/spacer_mev_array$(OBJ_EXT): ../src/muz/spacer/spacer_mev_array.cpp model/model.h.node model/model_evaluator_params.hpp.node ast/rewriter/rewriter_types.h.node model/model_evaluator.h.node muz/spacer/spacer_mev_array.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/arith_rewriter.h.node ast/rewriter/bv_rewriter.h.node ast/rewriter/datatype_rewriter.h.node ast/rewriter/array_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/ast_pp.h.node model/func_interp.h.node
	@echo src/muz/spacer/spacer_mev_array.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_mev_array$(OBJ_EXT) ../src/muz/spacer/spacer_mev_array.cpp
muz/spacer/spacer_farkas_learner$(OBJ_EXT): ../src/muz/spacer/spacer_farkas_learner.cpp ast/ast_smt2_pp.h.node ast/array_decl_plugin.h.node ast/rewriter/bool_rewriter.h.node ast/dl_decl_plugin.h.node ast/for_each_expr.h.node muz/base/dl_util.h.node ast/rewriter/rewriter.h.node ast/rewriter/rewriter_def.h.node muz/spacer/spacer_util.h.node muz/spacer/spacer_farkas_learner.h.node ast/rewriter/th_rewriter.h.node ast/ast_ll_pp.h.node ast/proofs/proof_utils.h.node ast/reg_decl_plugins.h.node smt/smt_farkas_util.h.node
	@echo src/muz/spacer/spacer_farkas_learner.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_farkas_learner$(OBJ_EXT) ../src/muz/spacer/spacer_farkas_learner.cpp
muz/spacer/spacer_sem_matcher$(OBJ_EXT): ../src/muz/spacer/spacer_sem_matcher.cpp muz/spacer/spacer_sem_matcher.h.node
	@echo src/muz/spacer/spacer_sem_matcher.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_sem_matcher$(OBJ_EXT) ../src/muz/spacer/spacer_sem_matcher.cpp
muz/spacer/spacer_prop_solver$(OBJ_EXT): ../src/muz/spacer/spacer_prop_solver.cpp ast/ast_smt2_pp.h.node ast/datatype_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/rewriter/expr_replacer.h.node smt/params/smt_params.h.node model/model.h.node model/model_pp.h.node muz/base/dl_util.h.node muz/spacer/spacer_util.h.node muz/spacer/spacer_farkas_learner.h.node muz/spacer/spacer_prop_solver.h.node model/model_evaluator.h.node muz/base/fp_params.hpp.node
	@echo src/muz/spacer/spacer_prop_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_prop_solver$(OBJ_EXT) ../src/muz/spacer/spacer_prop_solver.cpp
muz/spacer/spacer_pdr.h.node: ../src/muz/spacer/spacer_pdr.h muz/spacer/spacer_context.h.node
	@echo done > muz/spacer/spacer_pdr.h.node
muz/spacer/spacer_pdr$(OBJ_EXT): ../src/muz/spacer/spacer_pdr.cpp muz/spacer/spacer_pdr.h.node muz/base/dl_context.h.node muz/spacer/spacer_mbc.h.node
	@echo src/muz/spacer/spacer_pdr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_53) $(CXX_OUT_FLAG)muz/spacer/spacer_pdr$(OBJ_EXT) ../src/muz/spacer/spacer_pdr.cpp
muz/spacer/spacer$(LIB_EXT): muz/spacer/spacer_legacy_frames$(OBJ_EXT) muz/spacer/spacer_matrix$(OBJ_EXT) muz/spacer/spacer_proof_utils$(OBJ_EXT) muz/spacer/spacer_legacy_mbp$(OBJ_EXT) muz/spacer/spacer_json$(OBJ_EXT) muz/spacer/spacer_dl_interface$(OBJ_EXT) muz/spacer/spacer_iuc_proof$(OBJ_EXT) muz/spacer/spacer_util$(OBJ_EXT) muz/spacer/spacer_unsat_core_learner$(OBJ_EXT) muz/spacer/spacer_legacy_mev$(OBJ_EXT) muz/spacer/spacer_mbc$(OBJ_EXT) muz/spacer/spacer_iuc_solver$(OBJ_EXT) muz/spacer/spacer_manager$(OBJ_EXT) muz/spacer/spacer_generalizers$(OBJ_EXT) muz/spacer/spacer_callback$(OBJ_EXT) muz/spacer/spacer_sym_mux$(OBJ_EXT) muz/spacer/spacer_arith_generalizers$(OBJ_EXT) muz/spacer/spacer_qe_project$(OBJ_EXT) muz/spacer/spacer_sat_answer$(OBJ_EXT) muz/spacer/spacer_context$(OBJ_EXT) muz/spacer/spacer_quant_generalizer$(OBJ_EXT) muz/spacer/spacer_unsat_core_plugin$(OBJ_EXT) muz/spacer/spacer_antiunify$(OBJ_EXT) muz/spacer/spacer_mev_array$(OBJ_EXT) muz/spacer/spacer_farkas_learner$(OBJ_EXT) muz/spacer/spacer_sem_matcher$(OBJ_EXT) muz/spacer/spacer_prop_solver$(OBJ_EXT) muz/spacer/spacer_pdr$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/spacer/spacer$(LIB_EXT) muz/spacer/spacer_legacy_frames$(OBJ_EXT) muz/spacer/spacer_matrix$(OBJ_EXT) muz/spacer/spacer_proof_utils$(OBJ_EXT) muz/spacer/spacer_legacy_mbp$(OBJ_EXT) muz/spacer/spacer_json$(OBJ_EXT) muz/spacer/spacer_dl_interface$(OBJ_EXT) muz/spacer/spacer_iuc_proof$(OBJ_EXT) muz/spacer/spacer_util$(OBJ_EXT) muz/spacer/spacer_unsat_core_learner$(OBJ_EXT) muz/spacer/spacer_legacy_mev$(OBJ_EXT) muz/spacer/spacer_mbc$(OBJ_EXT) muz/spacer/spacer_iuc_solver$(OBJ_EXT) muz/spacer/spacer_manager$(OBJ_EXT) muz/spacer/spacer_generalizers$(OBJ_EXT) muz/spacer/spacer_callback$(OBJ_EXT) muz/spacer/spacer_sym_mux$(OBJ_EXT) muz/spacer/spacer_arith_generalizers$(OBJ_EXT) muz/spacer/spacer_qe_project$(OBJ_EXT) muz/spacer/spacer_sat_answer$(OBJ_EXT) muz/spacer/spacer_context$(OBJ_EXT) muz/spacer/spacer_quant_generalizer$(OBJ_EXT) muz/spacer/spacer_unsat_core_plugin$(OBJ_EXT) muz/spacer/spacer_antiunify$(OBJ_EXT) muz/spacer/spacer_mev_array$(OBJ_EXT) muz/spacer/spacer_farkas_learner$(OBJ_EXT) muz/spacer/spacer_sem_matcher$(OBJ_EXT) muz/spacer/spacer_prop_solver$(OBJ_EXT) muz/spacer/spacer_pdr$(OBJ_EXT)
spacer: muz/spacer/spacer$(LIB_EXT)

includes_54 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src
muz/clp/clp_context.h.node: ../src/muz/clp/clp_context.h ast/ast.h.node util/lbool.h.node util/statistics.h.node muz/base/dl_engine_base.h.node
	@echo done > muz/clp/clp_context.h.node
muz/clp/clp_context$(OBJ_EXT): ../src/muz/clp/clp_context.cpp muz/clp/clp_context.h.node muz/base/dl_context.h.node ast/substitution/unifier.h.node ast/rewriter/var_subst.h.node ast/substitution/substitution.h.node smt/smt_kernel.h.node muz/transforms/dl_transforms.h.node
	@echo src/muz/clp/clp_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_54) $(CXX_OUT_FLAG)muz/clp/clp_context$(OBJ_EXT) ../src/muz/clp/clp_context.cpp
muz/clp/clp$(LIB_EXT): muz/clp/clp_context$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/clp/clp$(LIB_EXT) muz/clp/clp_context$(OBJ_EXT)
clp: muz/clp/clp$(LIB_EXT)

includes_55 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src
muz/tab/tab_context.h.node: ../src/muz/tab/tab_context.h ast/ast.h.node util/lbool.h.node util/statistics.h.node muz/base/dl_engine_base.h.node
	@echo done > muz/tab/tab_context.h.node
muz/tab/tab_context$(OBJ_EXT): ../src/muz/tab/tab_context.cpp muz/tab/tab_context.h.node util/trail.h.node muz/base/dl_rule_set.h.node muz/base/dl_context.h.node muz/transforms/dl_mk_rule_inliner.h.node smt/smt_kernel.h.node qe/qe_lite.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/th_rewriter.h.node ast/datatype_decl_plugin.h.node ast/for_each_expr.h.node ast/substitution/matcher.h.node ast/scoped_proof.h.node muz/base/fp_params.hpp.node ast/ast_util.h.node
	@echo src/muz/tab/tab_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_55) $(CXX_OUT_FLAG)muz/tab/tab_context$(OBJ_EXT) ../src/muz/tab/tab_context.cpp
muz/tab/tab$(LIB_EXT): muz/tab/tab_context$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/tab/tab$(LIB_EXT) muz/tab/tab_context$(OBJ_EXT)
tab: muz/tab/tab$(LIB_EXT)

includes_56 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/rel -I../src
muz/ddnf/ddnf.h.node: ../src/muz/ddnf/ddnf.h ast/ast.h.node util/lbool.h.node util/statistics.h.node muz/base/dl_engine_base.h.node
	@echo done > muz/ddnf/ddnf.h.node
muz/ddnf/ddnf$(OBJ_EXT): ../src/muz/ddnf/ddnf.cpp muz/ddnf/ddnf.h.node muz/base/dl_rule_set.h.node muz/base/dl_context.h.node ast/scoped_proof.h.node ast/bv_decl_plugin.h.node muz/rel/tbv.h.node
	@echo src/muz/ddnf/ddnf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_56) $(CXX_OUT_FLAG)muz/ddnf/ddnf$(OBJ_EXT) ../src/muz/ddnf/ddnf.cpp
muz/ddnf/ddnf$(LIB_EXT): muz/ddnf/ddnf$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/ddnf/ddnf$(LIB_EXT) muz/ddnf/ddnf$(OBJ_EXT)
ddnf: muz/ddnf/ddnf$(LIB_EXT)

includes_57 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/tactic/fd_solver -I../src/sat/sat_solver -I../src/tactic/bv -I../src
muz/bmc/dl_bmc_engine.h.node: ../src/muz/bmc/dl_bmc_engine.h util/params.h.node util/statistics.h.node ast/bv_decl_plugin.h.node solver/solver.h.node
	@echo done > muz/bmc/dl_bmc_engine.h.node
muz/bmc/dl_bmc_engine$(OBJ_EXT): ../src/muz/bmc/dl_bmc_engine.cpp ast/datatype_decl_plugin.h.node ast/dl_decl_plugin.h.node ast/ast_smt_pp.h.node ast/well_sorted.h.node ast/rewriter/bool_rewriter.h.node ast/rewriter/rewriter_def.h.node ast/scoped_proof.h.node smt/smt_solver.h.node tactic/fd_solver/fd_solver.h.node tactic/tactic.h.node muz/base/dl_context.h.node muz/base/dl_rule_transformer.h.node muz/bmc/dl_bmc_engine.h.node muz/transforms/dl_mk_slice.h.node model/model_smt2_pp.h.node muz/transforms/dl_transforms.h.node muz/transforms/dl_mk_rule_inliner.h.node muz/base/fp_params.hpp.node
	@echo src/muz/bmc/dl_bmc_engine.cpp
	@$(CXX) $(CXXFLAGS) $(includes_57) $(CXX_OUT_FLAG)muz/bmc/dl_bmc_engine$(OBJ_EXT) ../src/muz/bmc/dl_bmc_engine.cpp
muz/bmc/bmc$(LIB_EXT): muz/bmc/dl_bmc_engine$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/bmc/bmc$(LIB_EXT) muz/bmc/dl_bmc_engine$(OBJ_EXT)
bmc: muz/bmc/bmc$(LIB_EXT)

includes_58 = -I../src/muz/base -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/aig -I../src/qe -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src/tactic/core -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/sat/sat_solver -I../src/tactic/bv -I../src/muz/ddnf -I../src/muz/spacer -I../src/smt/tactic -I../src
muz/fp/dl_cmds.h.node: ../src/muz/fp/dl_cmds.h ast/ast.h.node
	@echo done > muz/fp/dl_cmds.h.node
muz/fp/dl_register_engine.h.node: ../src/muz/fp/dl_register_engine.h muz/base/dl_context.h.node
	@echo done > muz/fp/dl_register_engine.h.node
muz/fp/dl_cmds$(OBJ_EXT): ../src/muz/fp/dl_cmds.cpp cmd_context/cmd_context.h.node muz/fp/dl_cmds.h.node muz/rel/dl_external_relation.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node ast/dl_decl_plugin.h.node muz/rel/dl_instruction.h.node muz/rel/dl_compiler.h.node muz/base/dl_rule.h.node ast/ast_pp.h.node cmd_context/parametric_cmd.h.node util/cancel_eh.h.node util/scoped_ctrl_c.h.node util/scoped_timer.h.node util/trail.h.node muz/base/fp_params.hpp.node
	@echo src/muz/fp/dl_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_58) $(CXX_OUT_FLAG)muz/fp/dl_cmds$(OBJ_EXT) ../src/muz/fp/dl_cmds.cpp
muz/fp/datalog_parser.h.node: ../src/muz/fp/datalog_parser.h ast/ast.h.node muz/base/dl_context.h.node
	@echo done > muz/fp/datalog_parser.h.node
muz/fp/datalog_parser$(OBJ_EXT): ../src/muz/fp/datalog_parser.cpp muz/fp/datalog_parser.h.node util/string_buffer.h.node util/str_hashtable.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node util/region.h.node util/warning.h.node
	@echo src/muz/fp/datalog_parser.cpp
	@$(CXX) $(CXXFLAGS) $(includes_58) $(CXX_OUT_FLAG)muz/fp/datalog_parser$(OBJ_EXT) ../src/muz/fp/datalog_parser.cpp
muz/fp/horn_tactic.h.node: ../src/muz/fp/horn_tactic.h util/params.h.node
	@echo done > muz/fp/horn_tactic.h.node
muz/fp/horn_tactic$(OBJ_EXT): ../src/muz/fp/horn_tactic.cpp ast/ast_util.h.node ast/rewriter/var_subst.h.node ast/rewriter/expr_replacer.h.node tactic/tactical.h.node tactic/model_converter.h.node tactic/proof_converter.h.node tactic/generic_model_converter.h.node muz/fp/horn_tactic.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node muz/base/dl_rule_transformer.h.node muz/transforms/dl_mk_slice.h.node muz/transforms/dl_transforms.h.node muz/base/fp_params.hpp.node
	@echo src/muz/fp/horn_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_58) $(CXX_OUT_FLAG)muz/fp/horn_tactic$(OBJ_EXT) ../src/muz/fp/horn_tactic.cpp
muz/fp/dl_register_engine$(OBJ_EXT): ../src/muz/fp/dl_register_engine.cpp muz/fp/dl_register_engine.h.node muz/bmc/dl_bmc_engine.h.node muz/clp/clp_context.h.node muz/tab/tab_context.h.node muz/rel/rel_context.h.node muz/ddnf/ddnf.h.node muz/spacer/spacer_dl_interface.h.node
	@echo src/muz/fp/dl_register_engine.cpp
	@$(CXX) $(CXXFLAGS) $(includes_58) $(CXX_OUT_FLAG)muz/fp/dl_register_engine$(OBJ_EXT) ../src/muz/fp/dl_register_engine.cpp
muz/fp/fp$(LIB_EXT): muz/fp/dl_cmds$(OBJ_EXT) muz/fp/datalog_parser$(OBJ_EXT) muz/fp/horn_tactic$(OBJ_EXT) muz/fp/dl_register_engine$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)muz/fp/fp$(LIB_EXT) muz/fp/dl_cmds$(OBJ_EXT) muz/fp/datalog_parser$(OBJ_EXT) muz/fp/horn_tactic$(OBJ_EXT) muz/fp/dl_register_engine$(OBJ_EXT)
fp: muz/fp/fp$(LIB_EXT)

includes_59 = -I../src/ast/normal_forms -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/tactic/core -I../src/tactic -I../src/model -I../src/ast/macros -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src/smt/tactic -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src
util/array_map.h.node: ../src/util/array_map.h util/vector.h.node util/optional.h.node
	@echo done > util/array_map.h.node
tactic/ufbv/ufbv_rewriter.h.node: ../src/tactic/ufbv/ufbv_rewriter.h ast/ast.h.node ast/substitution/substitution.h.node ast/rewriter/bool_rewriter.h.node util/obj_hashtable.h.node util/obj_pair_hashtable.h.node util/array_map.h.node
	@echo done > tactic/ufbv/ufbv_rewriter.h.node
tactic/ufbv/ufbv_rewriter$(OBJ_EXT): ../src/tactic/ufbv/ufbv_rewriter.cpp util/uint_set.h.node ast/ast_pp.h.node ast/for_each_expr.h.node ast/rewriter/var_subst.h.node tactic/ufbv/ufbv_rewriter.h.node
	@echo src/tactic/ufbv/ufbv_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_59) $(CXX_OUT_FLAG)tactic/ufbv/ufbv_rewriter$(OBJ_EXT) ../src/tactic/ufbv/ufbv_rewriter.cpp
tactic/ufbv/ufbv_rewriter_tactic.h.node: ../src/tactic/ufbv/ufbv_rewriter_tactic.h util/params.h.node
	@echo done > tactic/ufbv/ufbv_rewriter_tactic.h.node
tactic/ufbv/ufbv_rewriter_tactic$(OBJ_EXT): ../src/tactic/ufbv/ufbv_rewriter_tactic.cpp tactic/tactical.h.node tactic/ufbv/ufbv_rewriter.h.node tactic/ufbv/ufbv_rewriter_tactic.h.node
	@echo src/tactic/ufbv/ufbv_rewriter_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_59) $(CXX_OUT_FLAG)tactic/ufbv/ufbv_rewriter_tactic$(OBJ_EXT) ../src/tactic/ufbv/ufbv_rewriter_tactic.cpp
tactic/core/distribute_forall_tactic.h.node: ../src/tactic/core/distribute_forall_tactic.h util/params.h.node
	@echo done > tactic/core/distribute_forall_tactic.h.node
tactic/core/der_tactic.h.node: ../src/tactic/core/der_tactic.h
	@echo done > tactic/core/der_tactic.h.node
tactic/core/reduce_args_tactic.h.node: ../src/tactic/core/reduce_args_tactic.h util/params.h.node
	@echo done > tactic/core/reduce_args_tactic.h.node
smt/tactic/smt_tactic.h.node: ../src/smt/tactic/smt_tactic.h util/params.h.node ast/ast.h.node util/obj_hashtable.h.node tactic/goal.h.node
	@echo done > smt/tactic/smt_tactic.h.node
tactic/ufbv/macro_finder_tactic.h.node: ../src/tactic/ufbv/macro_finder_tactic.h util/params.h.node
	@echo done > tactic/ufbv/macro_finder_tactic.h.node
tactic/ufbv/quasi_macros_tactic.h.node: ../src/tactic/ufbv/quasi_macros_tactic.h util/params.h.node
	@echo done > tactic/ufbv/quasi_macros_tactic.h.node
tactic/ufbv/ufbv_tactic.h.node: ../src/tactic/ufbv/ufbv_tactic.h util/params.h.node
	@echo done > tactic/ufbv/ufbv_tactic.h.node
tactic/ufbv/ufbv_tactic$(OBJ_EXT): ../src/tactic/ufbv/ufbv_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/distribute_forall_tactic.h.node tactic/core/der_tactic.h.node tactic/core/reduce_args_tactic.h.node smt/tactic/smt_tactic.h.node tactic/core/nnf_tactic.h.node tactic/ufbv/macro_finder_tactic.h.node tactic/ufbv/ufbv_rewriter_tactic.h.node tactic/ufbv/quasi_macros_tactic.h.node tactic/ufbv/ufbv_tactic.h.node
	@echo src/tactic/ufbv/ufbv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_59) $(CXX_OUT_FLAG)tactic/ufbv/ufbv_tactic$(OBJ_EXT) ../src/tactic/ufbv/ufbv_tactic.cpp
tactic/ufbv/macro_finder_tactic$(OBJ_EXT): ../src/tactic/ufbv/macro_finder_tactic.cpp tactic/tactical.h.node ast/macros/macro_manager.h.node ast/macros/macro_finder.h.node tactic/generic_model_converter.h.node tactic/ufbv/macro_finder_tactic.h.node
	@echo src/tactic/ufbv/macro_finder_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_59) $(CXX_OUT_FLAG)tactic/ufbv/macro_finder_tactic$(OBJ_EXT) ../src/tactic/ufbv/macro_finder_tactic.cpp
tactic/ufbv/quasi_macros_tactic$(OBJ_EXT): ../src/tactic/ufbv/quasi_macros_tactic.cpp tactic/tactical.h.node tactic/generic_model_converter.h.node ast/macros/macro_manager.h.node ast/macros/macro_finder.h.node ast/macros/quasi_macros.h.node tactic/ufbv/quasi_macros_tactic.h.node
	@echo src/tactic/ufbv/quasi_macros_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_59) $(CXX_OUT_FLAG)tactic/ufbv/quasi_macros_tactic$(OBJ_EXT) ../src/tactic/ufbv/quasi_macros_tactic.cpp
tactic/ufbv/ufbv_tactic$(LIB_EXT): tactic/ufbv/ufbv_rewriter$(OBJ_EXT) tactic/ufbv/ufbv_rewriter_tactic$(OBJ_EXT) tactic/ufbv/ufbv_tactic$(OBJ_EXT) tactic/ufbv/macro_finder_tactic$(OBJ_EXT) tactic/ufbv/quasi_macros_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/ufbv/ufbv_tactic$(LIB_EXT) tactic/ufbv/ufbv_rewriter$(OBJ_EXT) tactic/ufbv/ufbv_rewriter_tactic$(OBJ_EXT) tactic/ufbv/ufbv_tactic$(OBJ_EXT) tactic/ufbv/macro_finder_tactic$(OBJ_EXT) tactic/ufbv/quasi_macros_tactic$(OBJ_EXT)
ufbv_tactic: tactic/ufbv/ufbv_tactic$(LIB_EXT)

includes_60 = -I../src/ackermannization -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src/sat/sat_solver -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/nlsat/tactic -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src
tactic/smtlogics/smt_tactic_select$(OBJ_EXT): ../src/tactic/smtlogics/smt_tactic_select.cpp smt/tactic/smt_tactic.h.node sat/tactic/sat_tactic.h.node sat/sat_params.hpp.node
	@echo src/tactic/smtlogics/smt_tactic_select.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/smt_tactic_select$(OBJ_EXT) ../src/tactic/smtlogics/smt_tactic_select.cpp
qe/qe_tactic.h.node: ../src/qe/qe_tactic.h util/params.h.node
	@echo done > qe/qe_tactic.h.node
tactic/arith/probe_arith.h.node: ../src/tactic/arith/probe_arith.h
	@echo done > tactic/arith/probe_arith.h.node
tactic/smtlogics/quant_tactics$(OBJ_EXT): ../src/tactic/smtlogics/quant_tactics.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node qe/qe_tactic.h.node qe/qe_lite.h.node qe/qsat.h.node tactic/core/ctx_simplify_tactic.h.node smt/tactic/smt_tactic.h.node tactic/core/elim_term_ite_tactic.h.node tactic/arith/probe_arith.h.node
	@echo src/tactic/smtlogics/quant_tactics.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/quant_tactics$(OBJ_EXT) ../src/tactic/smtlogics/quant_tactics.cpp
tactic/arith/nla2bv_tactic.h.node: ../src/tactic/arith/nla2bv_tactic.h util/params.h.node
	@echo done > tactic/arith/nla2bv_tactic.h.node
tactic/arith/lia2card_tactic.h.node: ../src/tactic/arith/lia2card_tactic.h util/params.h.node
	@echo done > tactic/arith/lia2card_tactic.h.node
tactic/core/cofactor_term_ite_tactic.h.node: ../src/tactic/core/cofactor_term_ite_tactic.h util/params.h.node
	@echo done > tactic/core/cofactor_term_ite_tactic.h.node
nlsat/tactic/qfnra_nlsat_tactic.h.node: ../src/nlsat/tactic/qfnra_nlsat_tactic.h util/params.h.node
	@echo done > nlsat/tactic/qfnra_nlsat_tactic.h.node
tactic/smtlogics/qfnia_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfnia_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node smt/tactic/smt_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node sat/tactic/sat_tactic.h.node tactic/arith/nla2bv_tactic.h.node tactic/arith/lia2card_tactic.h.node tactic/arith/card2bv_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/core/cofactor_term_ite_tactic.h.node nlsat/tactic/qfnra_nlsat_tactic.h.node
	@echo src/tactic/smtlogics/qfnia_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfnia_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfnia_tactic.cpp
tactic/smtlogics/qfbv_tactic.h.node: ../src/tactic/smtlogics/qfbv_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfbv_tactic.h.node
tactic/smtlogics/qfufbv_tactic_params.hpp.node: ../src/tactic/smtlogics/qfufbv_tactic_params.hpp util/params.h.node util/gparams.h.node
	@echo done > tactic/smtlogics/qfufbv_tactic_params.hpp.node
tactic/smtlogics/qfufbv_ackr_model_converter.h.node: ../src/tactic/smtlogics/qfufbv_ackr_model_converter.h tactic/model_converter.h.node ackermannization/ackr_info.h.node
	@echo done > tactic/smtlogics/qfufbv_ackr_model_converter.h.node
tactic/smtlogics/qfaufbv_tactic.h.node: ../src/tactic/smtlogics/qfaufbv_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfaufbv_tactic.h.node
tactic/smtlogics/smt_tactic_select.h.node: ../src/tactic/smtlogics/smt_tactic_select.h util/params.h.node
	@echo done > tactic/smtlogics/smt_tactic_select.h.node
tactic/smtlogics/qfufbv_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfufbv_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node smt/tactic/smt_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/core/reduce_args_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/qfufbv_tactic_params.hpp.node model/model_smt2_pp.h.node ackermannization/lackr.h.node ackermannization/ackermannization_params.hpp.node tactic/smtlogics/qfufbv_ackr_model_converter.h.node sat/sat_solver/inc_sat_solver.h.node tactic/smtlogics/qfaufbv_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/smt_tactic_select.h.node solver/tactic2solver.h.node tactic/bv/bv_bound_chk_tactic.h.node ackermannization/ackermannize_bv_tactic.h.node
	@echo src/tactic/smtlogics/qfufbv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfufbv_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfufbv_tactic.cpp
tactic/bv/bv1_blaster_tactic.h.node: ../src/tactic/bv/bv1_blaster_tactic.h util/params.h.node
	@echo done > tactic/bv/bv1_blaster_tactic.h.node
tactic/smtlogics/qfbv_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfbv_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/bv1_blaster_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/aig/aig_tactic.h.node smt/tactic/smt_tactic.h.node sat/tactic/sat_tactic.h.node sat/sat_solver/inc_sat_solver.h.node ackermannization/ackermannize_bv_tactic.h.node tactic/smtlogics/smt_tactic_select.h.node
	@echo src/tactic/smtlogics/qfbv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfbv_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfbv_tactic.cpp
tactic/smtlogics/qfnra_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfnra_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/arith/nla2bv_tactic.h.node smt/tactic/smt_tactic.h.node nlsat/tactic/qfnra_nlsat_tactic.h.node
	@echo src/tactic/smtlogics/qfnra_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfnra_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfnra_tactic.cpp
tactic/arith/normalize_bounds_tactic.h.node: ../src/tactic/arith/normalize_bounds_tactic.h util/params.h.node
	@echo done > tactic/arith/normalize_bounds_tactic.h.node
tactic/arith/fix_dl_var_tactic.h.node: ../src/tactic/arith/fix_dl_var_tactic.h util/params.h.node
	@echo done > tactic/arith/fix_dl_var_tactic.h.node
tactic/arith/lia2pb_tactic.h.node: ../src/tactic/arith/lia2pb_tactic.h util/params.h.node
	@echo done > tactic/arith/lia2pb_tactic.h.node
tactic/arith/diff_neq_tactic.h.node: ../src/tactic/arith/diff_neq_tactic.h util/params.h.node
	@echo done > tactic/arith/diff_neq_tactic.h.node
tactic/smtlogics/qfidl_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfidl_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/arith/normalize_bounds_tactic.h.node tactic/arith/fix_dl_var_tactic.h.node smt/tactic/smt_tactic.h.node tactic/arith/lia2pb_tactic.h.node tactic/arith/pb2bv_tactic.h.node tactic/arith/diff_neq_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/aig/aig_tactic.h.node sat/tactic/sat_tactic.h.node
	@echo src/tactic/smtlogics/qfidl_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfidl_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfidl_tactic.cpp
tactic/smtlogics/qfufbv_ackr_model_converter$(OBJ_EXT): ../src/tactic/smtlogics/qfufbv_ackr_model_converter.cpp tactic/smtlogics/qfufbv_ackr_model_converter.h.node ackermannization/ackr_model_converter.h.node
	@echo src/tactic/smtlogics/qfufbv_ackr_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfufbv_ackr_model_converter$(OBJ_EXT) ../src/tactic/smtlogics/qfufbv_ackr_model_converter.cpp
tactic/smtlogics/nra_tactic$(OBJ_EXT): ../src/tactic/smtlogics/nra_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/nnf_tactic.h.node tactic/arith/probe_arith.h.node smt/tactic/smt_tactic.h.node qe/qe_tactic.h.node qe/nlqsat.h.node qe/qe_lite.h.node nlsat/tactic/qfnra_nlsat_tactic.h.node
	@echo src/tactic/smtlogics/nra_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/nra_tactic$(OBJ_EXT) ../src/tactic/smtlogics/nra_tactic.cpp
tactic/smtlogics/qfaufbv_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfaufbv_tactic.cpp tactic/core/solve_eqs_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/smt_tactic_select.h.node ackermannization/ackermannize_bv_tactic.h.node smt/tactic/smt_tactic.h.node
	@echo src/tactic/smtlogics/qfaufbv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfaufbv_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfaufbv_tactic.cpp
tactic/arith/add_bounds_tactic.h.node: ../src/tactic/arith/add_bounds_tactic.h util/params.h.node
	@echo done > tactic/arith/add_bounds_tactic.h.node
tactic/smtlogics/qflia_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qflia_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/arith/normalize_bounds_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node smt/tactic/smt_tactic.h.node tactic/arith/add_bounds_tactic.h.node tactic/arith/pb2bv_tactic.h.node tactic/arith/lia2pb_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/aig/aig_tactic.h.node sat/tactic/sat_tactic.h.node tactic/arith/bound_manager.h.node tactic/arith/probe_arith.h.node
	@echo src/tactic/smtlogics/qflia_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qflia_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qflia_tactic.cpp
tactic/arith/recover_01_tactic.h.node: ../src/tactic/arith/recover_01_tactic.h util/params.h.node
	@echo done > tactic/arith/recover_01_tactic.h.node
tactic/smtlogics/qflra_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qflra_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node smt/tactic/smt_tactic.h.node tactic/arith/recover_01_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/arith/probe_arith.h.node
	@echo src/tactic/smtlogics/qflra_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qflra_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qflra_tactic.cpp
tactic/core/symmetry_reduce_tactic.h.node: ../src/tactic/core/symmetry_reduce_tactic.h util/params.h.node
	@echo done > tactic/core/symmetry_reduce_tactic.h.node
tactic/smtlogics/qfuf_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfuf_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/symmetry_reduce_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/propagate_values_tactic.h.node smt/tactic/smt_tactic.h.node tactic/smtlogics/smt_tactic_select.h.node
	@echo src/tactic/smtlogics/qfuf_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfuf_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfuf_tactic.cpp
tactic/smtlogics/qfauflia_tactic$(OBJ_EXT): ../src/tactic/smtlogics/qfauflia_tactic.cpp tactic/tactical.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node smt/tactic/smt_tactic.h.node
	@echo src/tactic/smtlogics/qfauflia_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_60) $(CXX_OUT_FLAG)tactic/smtlogics/qfauflia_tactic$(OBJ_EXT) ../src/tactic/smtlogics/qfauflia_tactic.cpp
tactic/smtlogics/smtlogic_tactics$(LIB_EXT): tactic/smtlogics/smt_tactic_select$(OBJ_EXT) tactic/smtlogics/quant_tactics$(OBJ_EXT) tactic/smtlogics/qfnia_tactic$(OBJ_EXT) tactic/smtlogics/qfufbv_tactic$(OBJ_EXT) tactic/smtlogics/qfbv_tactic$(OBJ_EXT) tactic/smtlogics/qfnra_tactic$(OBJ_EXT) tactic/smtlogics/qfidl_tactic$(OBJ_EXT) tactic/smtlogics/qfufbv_ackr_model_converter$(OBJ_EXT) tactic/smtlogics/nra_tactic$(OBJ_EXT) tactic/smtlogics/qfaufbv_tactic$(OBJ_EXT) tactic/smtlogics/qflia_tactic$(OBJ_EXT) tactic/smtlogics/qflra_tactic$(OBJ_EXT) tactic/smtlogics/qfuf_tactic$(OBJ_EXT) tactic/smtlogics/qfauflia_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/smtlogics/smt_tactic_select$(OBJ_EXT) tactic/smtlogics/quant_tactics$(OBJ_EXT) tactic/smtlogics/qfnia_tactic$(OBJ_EXT) tactic/smtlogics/qfufbv_tactic$(OBJ_EXT) tactic/smtlogics/qfbv_tactic$(OBJ_EXT) tactic/smtlogics/qfnra_tactic$(OBJ_EXT) tactic/smtlogics/qfidl_tactic$(OBJ_EXT) tactic/smtlogics/qfufbv_ackr_model_converter$(OBJ_EXT) tactic/smtlogics/nra_tactic$(OBJ_EXT) tactic/smtlogics/qfaufbv_tactic$(OBJ_EXT) tactic/smtlogics/qflia_tactic$(OBJ_EXT) tactic/smtlogics/qflra_tactic$(OBJ_EXT) tactic/smtlogics/qfuf_tactic$(OBJ_EXT) tactic/smtlogics/qfauflia_tactic$(OBJ_EXT)
smtlogic_tactics: tactic/smtlogics/smtlogic_tactics$(LIB_EXT)

includes_61 = -I../src/ast/fpa -I../src/ast -I../src/util -I../src/math/polynomial -I../src/ast/rewriter -I../src/math/automata -I../src/params -I../src/model -I../src/ast/macros -I../src/tactic/core -I../src/tactic -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/solver -I../src/ast/proofs -I../src/parsers/util -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/sat/tactic -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/tactic/arith -I../src/tactic/smtlogics -I../src/ackermannization -I../src/sat/sat_solver -I../src/tactic/aig -I../src/nlsat/tactic -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src
tactic/fpa/fpa2bv_tactic.h.node: ../src/tactic/fpa/fpa2bv_tactic.h util/params.h.node
	@echo done > tactic/fpa/fpa2bv_tactic.h.node
tactic/fpa/fpa2bv_model_converter.h.node: ../src/tactic/fpa/fpa2bv_model_converter.h ast/fpa/fpa2bv_converter.h.node tactic/model_converter.h.node ast/fpa/bv2fpa_converter.h.node
	@echo done > tactic/fpa/fpa2bv_model_converter.h.node
tactic/fpa/fpa2bv_tactic$(OBJ_EXT): ../src/tactic/fpa/fpa2bv_tactic.cpp tactic/tactical.h.node ast/fpa/fpa2bv_rewriter.h.node tactic/core/simplify_tactic.h.node tactic/fpa/fpa2bv_tactic.h.node tactic/fpa/fpa2bv_model_converter.h.node
	@echo src/tactic/fpa/fpa2bv_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_61) $(CXX_OUT_FLAG)tactic/fpa/fpa2bv_tactic$(OBJ_EXT) ../src/tactic/fpa/fpa2bv_tactic.cpp
tactic/fpa/qffp_tactic.h.node: ../src/tactic/fpa/qffp_tactic.h util/params.h.node
	@echo done > tactic/fpa/qffp_tactic.h.node
tactic/fpa/qffplra_tactic.h.node: ../src/tactic/fpa/qffplra_tactic.h util/params.h.node
	@echo done > tactic/fpa/qffplra_tactic.h.node
tactic/fpa/qffplra_tactic$(OBJ_EXT): ../src/tactic/fpa/qffplra_tactic.cpp tactic/tactical.h.node tactic/fpa/qffp_tactic.h.node tactic/fpa/qffplra_tactic.h.node
	@echo src/tactic/fpa/qffplra_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_61) $(CXX_OUT_FLAG)tactic/fpa/qffplra_tactic$(OBJ_EXT) ../src/tactic/fpa/qffplra_tactic.cpp
tactic/fpa/fpa2bv_model_converter$(OBJ_EXT): ../src/tactic/fpa/fpa2bv_model_converter.cpp ast/ast_smt2_pp.h.node ast/rewriter/fpa_rewriter.h.node tactic/fpa/fpa2bv_model_converter.h.node
	@echo src/tactic/fpa/fpa2bv_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_61) $(CXX_OUT_FLAG)tactic/fpa/fpa2bv_model_converter$(OBJ_EXT) ../src/tactic/fpa/fpa2bv_model_converter.cpp
tactic/smtlogics/qfnra_tactic.h.node: ../src/tactic/smtlogics/qfnra_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfnra_tactic.h.node
tactic/fpa/qffp_tactic$(OBJ_EXT): ../src/tactic/fpa/qffp_tactic.cpp tactic/tactical.h.node tactic/fpa/fpa2bv_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/arith/probe_arith.h.node tactic/bv/bit_blaster_tactic.h.node tactic/smtlogics/qfnra_tactic.h.node sat/tactic/sat_tactic.h.node sat/sat_solver/inc_sat_solver.h.node smt/tactic/smt_tactic.h.node ackermannization/ackermannize_bv_tactic.h.node tactic/fpa/qffp_tactic.h.node
	@echo src/tactic/fpa/qffp_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_61) $(CXX_OUT_FLAG)tactic/fpa/qffp_tactic$(OBJ_EXT) ../src/tactic/fpa/qffp_tactic.cpp
tactic/fpa/fpa_tactics$(LIB_EXT): tactic/fpa/fpa2bv_tactic$(OBJ_EXT) tactic/fpa/qffplra_tactic$(OBJ_EXT) tactic/fpa/fpa2bv_model_converter$(OBJ_EXT) tactic/fpa/qffp_tactic$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/fpa/fpa_tactics$(LIB_EXT) tactic/fpa/fpa2bv_tactic$(OBJ_EXT) tactic/fpa/qffplra_tactic$(OBJ_EXT) tactic/fpa/fpa2bv_model_converter$(OBJ_EXT) tactic/fpa/qffp_tactic$(OBJ_EXT)
fpa_tactics: tactic/fpa/fpa_tactics$(LIB_EXT)

includes_62 = -I../src/tactic/smtlogics -I../src/ackermannization -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src/sat/sat_solver -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/nlsat/tactic -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src/tactic/ufbv -I../src/tactic/fpa -I../src/tactic/sls -I../src/math/subpaving/tactic -I../src/math/subpaving -I../src
tactic/portfolio/solver2lookahead$(OBJ_EXT): ../src/tactic/portfolio/solver2lookahead.cpp sat/sat_solver/inc_sat_solver.h.node solver/solver.h.node
	@echo src/tactic/portfolio/solver2lookahead.cpp
	@$(CXX) $(CXXFLAGS) $(includes_62) $(CXX_OUT_FLAG)tactic/portfolio/solver2lookahead$(OBJ_EXT) ../src/tactic/portfolio/solver2lookahead.cpp
tactic/portfolio/default_tactic.h.node: ../src/tactic/portfolio/default_tactic.h util/params.h.node
	@echo done > tactic/portfolio/default_tactic.h.node
tactic/smtlogics/qflia_tactic.h.node: ../src/tactic/smtlogics/qflia_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qflia_tactic.h.node
tactic/smtlogics/qflra_tactic.h.node: ../src/tactic/smtlogics/qflra_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qflra_tactic.h.node
tactic/smtlogics/qfnia_tactic.h.node: ../src/tactic/smtlogics/qfnia_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfnia_tactic.h.node
tactic/smtlogics/nra_tactic.h.node: ../src/tactic/smtlogics/nra_tactic.h
	@echo done > tactic/smtlogics/nra_tactic.h.node
tactic/smtlogics/quant_tactics.h.node: ../src/tactic/smtlogics/quant_tactics.h util/params.h.node
	@echo done > tactic/smtlogics/quant_tactics.h.node
tactic/smtlogics/qfauflia_tactic.h.node: ../src/tactic/smtlogics/qfauflia_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfauflia_tactic.h.node
tactic/portfolio/default_tactic$(OBJ_EXT): ../src/tactic/portfolio/default_tactic.cpp tactic/portfolio/default_tactic.h.node tactic/core/simplify_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node smt/tactic/smt_tactic.h.node tactic/smtlogics/qflia_tactic.h.node tactic/smtlogics/qflra_tactic.h.node tactic/smtlogics/qfnia_tactic.h.node tactic/smtlogics/qfnra_tactic.h.node tactic/smtlogics/nra_tactic.h.node tactic/arith/probe_arith.h.node tactic/smtlogics/quant_tactics.h.node tactic/fpa/qffp_tactic.h.node tactic/fpa/qffplra_tactic.h.node tactic/smtlogics/qfaufbv_tactic.h.node tactic/smtlogics/qfauflia_tactic.h.node tactic/fd_solver/fd_solver.h.node
	@echo src/tactic/portfolio/default_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_62) $(CXX_OUT_FLAG)tactic/portfolio/default_tactic$(OBJ_EXT) ../src/tactic/portfolio/default_tactic.cpp
solver/combined_solver.h.node: ../src/solver/combined_solver.h util/params.h.node
	@echo done > solver/combined_solver.h.node
tactic/smtlogics/qfuf_tactic.h.node: ../src/tactic/smtlogics/qfuf_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfuf_tactic.h.node
tactic/smtlogics/qfufbv_tactic.h.node: ../src/tactic/smtlogics/qfufbv_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfufbv_tactic.h.node
tactic/smtlogics/qfidl_tactic.h.node: ../src/tactic/smtlogics/qfidl_tactic.h util/params.h.node
	@echo done > tactic/smtlogics/qfidl_tactic.h.node
tactic/fd_solver/smtfd_solver.h.node: ../src/tactic/fd_solver/smtfd_solver.h ast/ast.h.node util/params.h.node
	@echo done > tactic/fd_solver/smtfd_solver.h.node
tactic/portfolio/smt_strategic_solver$(OBJ_EXT): ../src/tactic/portfolio/smt_strategic_solver.cpp cmd_context/cmd_context.h.node solver/combined_solver.h.node solver/tactic2solver.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/qflia_tactic.h.node tactic/smtlogics/qfnia_tactic.h.node tactic/smtlogics/qfnra_tactic.h.node tactic/smtlogics/qfuf_tactic.h.node tactic/smtlogics/qflra_tactic.h.node tactic/smtlogics/quant_tactics.h.node tactic/smtlogics/qfauflia_tactic.h.node tactic/smtlogics/qfaufbv_tactic.h.node tactic/smtlogics/qfufbv_tactic.h.node tactic/smtlogics/qfidl_tactic.h.node tactic/smtlogics/nra_tactic.h.node tactic/portfolio/default_tactic.h.node tactic/fd_solver/fd_solver.h.node tactic/fd_solver/smtfd_solver.h.node tactic/ufbv/ufbv_tactic.h.node tactic/fpa/qffp_tactic.h.node muz/fp/horn_tactic.h.node smt/smt_solver.h.node sat/sat_solver/inc_sat_solver.h.node ast/rewriter/bv_rewriter.h.node solver/solver2tactic.h.node solver/parallel_tactic.h.node solver/parallel_params.hpp.node tactic/tactic_params.hpp.node parsers/smt2/smt2parser.h.node
	@echo src/tactic/portfolio/smt_strategic_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_62) $(CXX_OUT_FLAG)tactic/portfolio/smt_strategic_solver$(OBJ_EXT) ../src/tactic/portfolio/smt_strategic_solver.cpp
tactic/portfolio/portfolio$(LIB_EXT): tactic/portfolio/solver2lookahead$(OBJ_EXT) tactic/portfolio/default_tactic$(OBJ_EXT) tactic/portfolio/smt_strategic_solver$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)tactic/portfolio/portfolio$(LIB_EXT) tactic/portfolio/solver2lookahead$(OBJ_EXT) tactic/portfolio/default_tactic$(OBJ_EXT) tactic/portfolio/smt_strategic_solver$(OBJ_EXT)
portfolio: tactic/portfolio/portfolio$(LIB_EXT)

includes_63 = -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/ast/normal_forms -I../src/cmd_context -I../src/solver -I../src/model -I../src/tactic -I../src/ast/proofs -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/interval -I../src/math/simplex -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/tactic/smtlogics -I../src/ackermannization -I../src/sat/sat_solver -I../src/tactic/core -I../src/tactic/aig -I../src/tactic/bv -I../src/tactic/arith -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/qe/mbp -I../src/nlsat/tactic -I../src/smt/tactic -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src/tactic/sls -I../src
opt/opt_solver.h.node: ../src/opt/opt_solver.h util/inf_rational.h.node util/inf_eps_rational.h.node ast/ast.h.node util/params.h.node solver/solver_na2as.h.node smt/smt_kernel.h.node smt/params/smt_params.h.node smt/smt_types.h.node smt/theory_opt.h.node tactic/generic_model_converter.h.node
	@echo done > opt/opt_solver.h.node
opt/maxsmt.h.node: ../src/opt/maxsmt.h ast/ast.h.node util/params.h.node solver/solver.h.node util/statistics.h.node smt/smt_context.h.node smt/smt_theory.h.node smt/theory_wmaxsat.h.node opt/opt_solver.h.node
	@echo done > opt/maxsmt.h.node
opt/maxres.h.node: ../src/opt/maxres.h
	@echo done > opt/maxres.h.node
opt/maxlex.h.node: ../src/opt/maxlex.h
	@echo done > opt/maxlex.h.node
opt/wmax.h.node: ../src/opt/wmax.h opt/maxsmt.h.node
	@echo done > opt/wmax.h.node
opt/opt_params.hpp.node: ../src/opt/opt_params.hpp util/params.h.node util/gparams.h.node
	@echo done > opt/opt_params.hpp.node
opt/opt_pareto.h.node: ../src/opt/opt_pareto.h solver/solver.h.node model/model.h.node
	@echo done > opt/opt_pareto.h.node
opt/optsmt.h.node: ../src/opt/optsmt.h opt/opt_solver.h.node
	@echo done > opt/optsmt.h.node
opt/opt_context.h.node: ../src/opt/opt_context.h ast/ast.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node tactic/model_converter.h.node tactic/tactic.h.node qe/qsat.h.node opt/opt_solver.h.node opt/opt_pareto.h.node opt/optsmt.h.node opt/maxsmt.h.node cmd_context/cmd_context.h.node
	@echo done > opt/opt_context.h.node
opt/maxsmt$(OBJ_EXT): ../src/opt/maxsmt.cpp util/uint_set.h.node ast/ast_pp.h.node ast/ast_util.h.node ast/pb_decl_plugin.h.node opt/maxsmt.h.node opt/maxres.h.node opt/maxlex.h.node opt/wmax.h.node opt/opt_params.hpp.node opt/opt_context.h.node smt/theory_wmaxsat.h.node smt/theory_pb.h.node
	@echo src/opt/maxsmt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/maxsmt$(OBJ_EXT) ../src/opt/maxsmt.cpp
opt/opt_solver$(OBJ_EXT): ../src/opt/opt_solver.cpp ast/reg_decl_plugins.h.node opt/opt_solver.h.node smt/smt_context.h.node smt/theory_arith.h.node smt/theory_diff_logic.h.node smt/theory_dense_diff_logic.h.node smt/theory_pb.h.node smt/theory_lra.h.node ast/ast_pp.h.node ast/ast_smt_pp.h.node ast/pp_params.hpp.node opt/opt_params.hpp.node model/model_smt2_pp.h.node util/stopwatch.h.node
	@echo src/opt/opt_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/opt_solver$(OBJ_EXT) ../src/opt/opt_solver.cpp
opt/maxlex$(OBJ_EXT): ../src/opt/maxlex.cpp ast/ast_pp.h.node opt/opt_context.h.node opt/maxsmt.h.node opt/maxlex.h.node
	@echo src/opt/maxlex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/maxlex$(OBJ_EXT) ../src/opt/maxlex.cpp
opt/sortmax$(OBJ_EXT): ../src/opt/sortmax.cpp opt/maxsmt.h.node util/uint_set.h.node ast/ast_pp.h.node model/model_smt2_pp.h.node smt/smt_theory.h.node smt/smt_context.h.node opt/opt_context.h.node util/sorting_network.h.node tactic/generic_model_converter.h.node
	@echo src/opt/sortmax.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/sortmax$(OBJ_EXT) ../src/opt/sortmax.cpp
opt/wmax$(OBJ_EXT): ../src/opt/wmax.cpp opt/wmax.h.node util/uint_set.h.node ast/ast_pp.h.node model/model_smt2_pp.h.node smt/smt_theory.h.node smt/smt_context.h.node smt/theory_wmaxsat.h.node opt/opt_context.h.node
	@echo src/opt/wmax.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/wmax$(OBJ_EXT) ../src/opt/wmax.cpp
opt/optsmt$(OBJ_EXT): ../src/opt/optsmt.cpp opt/optsmt.h.node opt/opt_solver.h.node opt/opt_context.h.node ast/arith_decl_plugin.h.node smt/theory_arith.h.node ast/ast_pp.h.node ast/ast_util.h.node model/model_pp.h.node ast/rewriter/th_rewriter.h.node opt/opt_params.hpp.node
	@echo src/opt/optsmt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/optsmt$(OBJ_EXT) ../src/opt/optsmt.cpp
opt/opt_pareto$(OBJ_EXT): ../src/opt/opt_pareto.cpp opt/opt_pareto.h.node ast/ast_pp.h.node ast/ast_util.h.node model/model_smt2_pp.h.node
	@echo src/opt/opt_pareto.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/opt_pareto$(OBJ_EXT) ../src/opt/opt_pareto.cpp
opt/opt_parse.h.node: ../src/opt/opt_parse.h
	@echo done > opt/opt_parse.h.node
opt/opt_parse$(OBJ_EXT): ../src/opt/opt_parse.cpp opt/opt_context.h.node opt/opt_parse.h.node
	@echo src/opt/opt_parse.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/opt_parse$(OBJ_EXT) ../src/opt/opt_parse.cpp
opt/opt_cmds.h.node: ../src/opt/opt_cmds.h ast/ast.h.node opt/opt_context.h.node
	@echo done > opt/opt_cmds.h.node
opt/opt_cmds$(OBJ_EXT): ../src/opt/opt_cmds.cpp opt/opt_cmds.h.node cmd_context/cmd_context.h.node ast/ast_pp.h.node opt/opt_context.h.node util/cancel_eh.h.node util/scoped_ctrl_c.h.node util/scoped_timer.h.node cmd_context/parametric_cmd.h.node opt/opt_params.hpp.node model/model_smt2_pp.h.node
	@echo src/opt/opt_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/opt_cmds$(OBJ_EXT) ../src/opt/opt_cmds.cpp
opt/pb_sls.h.node: ../src/opt/pb_sls.h ast/pb_decl_plugin.h.node model/model.h.node util/lbool.h.node util/params.h.node util/statistics.h.node
	@echo done > opt/pb_sls.h.node
opt/pb_sls$(OBJ_EXT): ../src/opt/pb_sls.cpp opt/pb_sls.h.node smt/smt_literal.h.node ast/ast_pp.h.node ast/rewriter/th_rewriter.h.node sat/sat_types.h.node
	@echo src/opt/pb_sls.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/pb_sls$(OBJ_EXT) ../src/opt/pb_sls.cpp
opt/maxres$(OBJ_EXT): ../src/opt/maxres.cpp ast/ast_pp.h.node ast/pb_decl_plugin.h.node ast/ast_util.h.node model/model_smt2_pp.h.node solver/solver.h.node solver/mus.h.node sat/sat_solver/inc_sat_solver.h.node smt/smt_solver.h.node opt/opt_context.h.node opt/opt_params.hpp.node opt/maxsmt.h.node opt/maxres.h.node
	@echo src/opt/maxres.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/maxres$(OBJ_EXT) ../src/opt/maxres.cpp
tactic/arith/eq2bv_tactic.h.node: ../src/tactic/arith/eq2bv_tactic.h util/params.h.node
	@echo done > tactic/arith/eq2bv_tactic.h.node
opt/opt_context$(OBJ_EXT): ../src/opt/opt_context.cpp util/gparams.h.node ast/for_each_expr.h.node ast/ast_pp.h.node ast/bv_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/ast_smt_pp.h.node ast/ast_pp_util.h.node model/model_smt2_pp.h.node tactic/goal.h.node tactic/tactic.h.node tactic/arith/lia2card_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/tactical.h.node tactic/arith/card2bv_tactic.h.node tactic/arith/eq2bv_tactic.h.node tactic/bv/dt2bv_tactic.h.node tactic/generic_model_converter.h.node ackermannization/ackermannize_bv_tactic.h.node sat/sat_solver/inc_sat_solver.h.node qe/qsat.h.node opt/opt_context.h.node opt/opt_solver.h.node opt/opt_params.hpp.node
	@echo src/opt/opt_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_63) $(CXX_OUT_FLAG)opt/opt_context$(OBJ_EXT) ../src/opt/opt_context.cpp
opt/opt$(LIB_EXT): opt/maxsmt$(OBJ_EXT) opt/opt_solver$(OBJ_EXT) opt/maxlex$(OBJ_EXT) opt/sortmax$(OBJ_EXT) opt/wmax$(OBJ_EXT) opt/optsmt$(OBJ_EXT) opt/opt_pareto$(OBJ_EXT) opt/opt_parse$(OBJ_EXT) opt/opt_cmds$(OBJ_EXT) opt/pb_sls$(OBJ_EXT) opt/maxres$(OBJ_EXT) opt/opt_context$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)opt/opt$(LIB_EXT) opt/maxsmt$(OBJ_EXT) opt/opt_solver$(OBJ_EXT) opt/maxlex$(OBJ_EXT) opt/sortmax$(OBJ_EXT) opt/wmax$(OBJ_EXT) opt/optsmt$(OBJ_EXT) opt/opt_pareto$(OBJ_EXT) opt/opt_parse$(OBJ_EXT) opt/opt_cmds$(OBJ_EXT) opt/pb_sls$(OBJ_EXT) opt/maxres$(OBJ_EXT) opt/opt_context$(OBJ_EXT)
opt: opt/opt$(LIB_EXT)

includes_64 = -I../src/tactic/portfolio -I../src/tactic/smtlogics -I../src/ackermannization -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src/sat/sat_solver -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/nlsat/tactic -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src/tactic/ufbv -I../src/tactic/fpa -I../src/tactic/sls -I../src/math/subpaving/tactic -I../src/math/subpaving -I../src/math/realclosure -I../src/opt -I../src
api/z3_macros.h.node: ../src/api/z3_macros.h
	@echo done > api/z3_macros.h.node
api/z3_api.h.node: ../src/api/z3_api.h
	@echo done > api/z3_api.h.node
api/z3_ast_containers.h.node: ../src/api/z3_ast_containers.h
	@echo done > api/z3_ast_containers.h.node
api/z3_algebraic.h.node: ../src/api/z3_algebraic.h
	@echo done > api/z3_algebraic.h.node
api/z3_polynomial.h.node: ../src/api/z3_polynomial.h
	@echo done > api/z3_polynomial.h.node
api/z3_rcf.h.node: ../src/api/z3_rcf.h
	@echo done > api/z3_rcf.h.node
api/z3_fixedpoint.h.node: ../src/api/z3_fixedpoint.h
	@echo done > api/z3_fixedpoint.h.node
api/z3_optimization.h.node: ../src/api/z3_optimization.h
	@echo done > api/z3_optimization.h.node
api/z3_fpa.h.node: ../src/api/z3_fpa.h
	@echo done > api/z3_fpa.h.node
api/z3_spacer.h.node: ../src/api/z3_spacer.h
	@echo done > api/z3_spacer.h.node
api/z3.h.node: ../src/api/z3.h api/z3_macros.h.node api/z3_api.h.node api/z3_ast_containers.h.node api/z3_algebraic.h.node api/z3_polynomial.h.node api/z3_rcf.h.node api/z3_fixedpoint.h.node api/z3_optimization.h.node api/z3_fpa.h.node api/z3_spacer.h.node
	@echo done > api/z3.h.node
api/z3_replayer.h.node: ../src/api/z3_replayer.h api/z3.h.node util/z3_exception.h.node
	@echo done > api/z3_replayer.h.node
util/stream_buffer.h.node: ../src/util/stream_buffer.h
	@echo done > util/stream_buffer.h.node
api/z3_replayer$(OBJ_EXT): ../src/api/z3_replayer.cpp util/vector.h.node util/map.h.node api/z3_replayer.h.node util/stream_buffer.h.node util/symbol.h.node util/trace.h.node
	@echo src/api/z3_replayer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/z3_replayer$(OBJ_EXT) ../src/api/z3_replayer.cpp
api/api_log_macros.h.node: ../src/api/api_log_macros.h api/z3.h.node
	@echo done > api/api_log_macros.h.node
api/api_util.h.node: ../src/api/api_util.h util/params.h.node util/lbool.h.node ast/ast.h.node
	@echo done > api/api_util.h.node
api/api_polynomial.h.node: ../src/api/api_polynomial.h math/polynomial/polynomial.h.node
	@echo done > api/api_polynomial.h.node
api/api_context.h.node: ../src/api/api_context.h util/hashtable.h.node util/mutex.h.node util/event_handler.h.node ast/ast.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/dl_decl_plugin.h.node ast/fpa_decl_plugin.h.node ast/recfun_decl_plugin.h.node ast/special_relations_decl_plugin.h.node ast/rewriter/seq_rewriter.h.node smt/params/smt_params.h.node smt/smt_kernel.h.node smt/smt_solver.h.node cmd_context/tactic_manager.h.node params/context_params.h.node cmd_context/cmd_context.h.node solver/solver.h.node api/z3.h.node api/api_util.h.node api/api_polynomial.h.node
	@echo done > api/api_context.h.node
api/api_ast_vector.h.node: ../src/api/api_ast_vector.h api/api_util.h.node
	@echo done > api/api_ast_vector.h.node
api/api_ast_vector$(OBJ_EXT): ../src/api/api_ast_vector.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_ast_vector.h.node ast/ast_translation.h.node ast/ast_smt2_pp.h.node
	@echo src/api/api_ast_vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_ast_vector$(OBJ_EXT) ../src/api/api_ast_vector.cpp
api/z3_logger.h.node: ../src/api/z3_logger.h util/symbol.h.node
	@echo done > api/z3_logger.h.node
api/api_log_macros$(OBJ_EXT): ../src/api/api_log_macros.cpp api/z3.h.node api/api_log_macros.h.node api/z3_logger.h.node
	@echo src/api/api_log_macros.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_log_macros$(OBJ_EXT) ../src/api/api_log_macros.cpp
api/api_context$(OBJ_EXT): ../src/api/api_context.cpp util/z3_version.h.node api/api_context.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node api/api_log_macros.h.node api/api_util.h.node ast/reg_decl_plugins.h.node math/realclosure/realclosure.h.node
	@echo src/api/api_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_context$(OBJ_EXT) ../src/api/api_context.cpp
api/api_stats.h.node: ../src/api/api_stats.h api/api_util.h.node util/statistics.h.node
	@echo done > api/api_stats.h.node
api/api_stats$(OBJ_EXT): ../src/api/api_stats.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_stats.h.node
	@echo src/api/api_stats.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_stats$(OBJ_EXT) ../src/api/api_stats.cpp
api/api_rcf$(OBJ_EXT): ../src/api/api_rcf.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node math/realclosure/realclosure.h.node
	@echo src/api/api_rcf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_rcf$(OBJ_EXT) ../src/api/api_rcf.cpp
api/api_goal.h.node: ../src/api/api_goal.h api/api_util.h.node tactic/goal.h.node
	@echo done > api/api_goal.h.node
api/api_tactic.h.node: ../src/api/api_tactic.h api/api_goal.h.node tactic/tactical.h.node
	@echo done > api/api_tactic.h.node
api/api_model.h.node: ../src/api/api_model.h api/api_util.h.node model/model.h.node
	@echo done > api/api_model.h.node
api/api_tactic$(OBJ_EXT): ../src/api/api_tactic.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_tactic.h.node api/api_model.h.node util/scoped_ctrl_c.h.node util/cancel_eh.h.node util/scoped_timer.h.node
	@echo src/api/api_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_tactic$(OBJ_EXT) ../src/api/api_tactic.cpp
api/api_special_relations$(OBJ_EXT): ../src/api/api_special_relations.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/ast_pp.h.node ast/special_relations_decl_plugin.h.node
	@echo src/api/api_special_relations.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_special_relations$(OBJ_EXT) ../src/api/api_special_relations.cpp
api/api_commands$(OBJ_EXT): ../src/api/api_commands.cpp api/z3.h.node api/z3_replayer.h.node
	@echo src/api/api_commands.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_commands$(OBJ_EXT) ../src/api/api_commands.cpp
api/api_ast$(OBJ_EXT): ../src/api/api_ast.cpp api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/well_sorted.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/array_decl_plugin.h.node ast/pb_decl_plugin.h.node ast/ast_translation.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/ast_smt_pp.h.node ast/ast_smt2_pp.h.node ast/rewriter/th_rewriter.h.node ast/rewriter/var_subst.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/recfun_replace.h.node ast/rewriter/seq_rewriter.h.node ast/pp.h.node util/scoped_ctrl_c.h.node util/cancel_eh.h.node util/scoped_timer.h.node ast/pp_params.hpp.node ast/expr_abstract.h.node
	@echo src/api/api_ast.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_ast$(OBJ_EXT) ../src/api/api_ast.cpp
api/api_pb$(OBJ_EXT): ../src/api/api_pb.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/pb_decl_plugin.h.node
	@echo src/api/api_pb.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_pb$(OBJ_EXT) ../src/api/api_pb.cpp
api/api_ast_map.h.node: ../src/api/api_ast_map.h api/api_util.h.node util/obj_hashtable.h.node
	@echo done > api/api_ast_map.h.node
api/api_ast_map$(OBJ_EXT): ../src/api/api_ast_map.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_ast_map.h.node api/api_ast_vector.h.node ast/ast_smt2_pp.h.node util/dec_ref_util.h.node
	@echo src/api/api_ast_map.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_ast_map$(OBJ_EXT) ../src/api/api_ast_map.cpp
api/api_config_params$(OBJ_EXT): ../src/api/api_config_params.cpp api/z3.h.node api/api_context.h.node ast/pp.h.node api/api_log_macros.h.node api/api_util.h.node cmd_context/cmd_context.h.node util/symbol.h.node util/gparams.h.node util/env_params.h.node params/context_params.h.node
	@echo src/api/api_config_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_config_params$(OBJ_EXT) ../src/api/api_config_params.cpp
api/api_log$(OBJ_EXT): ../src/api/api_log.cpp api/z3.h.node api/api_log_macros.h.node util/util.h.node util/z3_version.h.node util/mutex.h.node
	@echo src/api/api_log.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_log$(OBJ_EXT) ../src/api/api_log.cpp
api/api_quant$(OBJ_EXT): ../src/api/api_quant.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node parsers/util/pattern_validation.h.node ast/expr_abstract.h.node
	@echo src/api/api_quant.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_quant$(OBJ_EXT) ../src/api/api_quant.cpp
api/api_numeral$(OBJ_EXT): ../src/api/api_numeral.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node math/polynomial/algebraic_numbers.h.node ast/fpa_decl_plugin.h.node
	@echo src/api/api_numeral.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_numeral$(OBJ_EXT) ../src/api/api_numeral.cpp
api/api_datalog.h.node: ../src/api/api_datalog.h api/z3.h.node ast/ast.h.node smt/params/smt_params.h.node smt/smt_kernel.h.node api/api_util.h.node
	@echo done > api/api_datalog.h.node
api/api_datalog$(OBJ_EXT): ../src/api/api_datalog.cpp api/api_datalog.h.node api/api_context.h.node api/api_util.h.node ast/ast_pp.h.node api/api_ast_vector.h.node api/api_log_macros.h.node api/api_stats.h.node muz/fp/datalog_parser.h.node util/cancel_eh.h.node util/scoped_ctrl_c.h.node util/scoped_timer.h.node muz/fp/dl_cmds.h.node cmd_context/cmd_context.h.node parsers/smt2/smt2parser.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node muz/rel/dl_external_relation.h.node ast/dl_decl_plugin.h.node muz/rel/rel_context.h.node
	@echo src/api/api_datalog.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_datalog$(OBJ_EXT) ../src/api/api_datalog.cpp
api/api_array$(OBJ_EXT): ../src/api/api_array.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/array_decl_plugin.h.node
	@echo src/api/api_array.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_array$(OBJ_EXT) ../src/api/api_array.cpp
util/file_path.h.node: ../src/util/file_path.h
	@echo done > util/file_path.h.node
api/api_solver.h.node: ../src/api/api_solver.h util/mutex.h.node api/api_util.h.node solver/solver.h.node
	@echo done > api/api_solver.h.node
api/api_solver$(OBJ_EXT): ../src/api/api_solver.cpp util/scoped_ctrl_c.h.node util/cancel_eh.h.node util/file_path.h.node util/scoped_timer.h.node util/file_path.h.node ast/ast_pp.h.node api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_tactic.h.node api/api_solver.h.node api/api_model.h.node api/api_stats.h.node api/api_ast_vector.h.node model/model_params.hpp.node smt/smt_solver.h.node smt/smt_implied_equalities.h.node solver/smt_logics.h.node solver/tactic2solver.h.node solver/solver_params.hpp.node cmd_context/cmd_context.h.node parsers/smt2/smt2parser.h.node sat/dimacs.h.node sat/sat_solver.h.node sat/tactic/goal2sat.h.node
	@echo src/api/api_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_solver$(OBJ_EXT) ../src/api/api_solver.cpp
api/api_qe$(OBJ_EXT): ../src/api/api_qe.cpp ast/expr_map.h.node api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node api/api_model.h.node api/api_ast_map.h.node api/api_ast_vector.h.node qe/qe_lite.h.node muz/spacer/spacer_util.h.node
	@echo src/api/api_qe.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_qe$(OBJ_EXT) ../src/api/api_qe.cpp
api/api_parsers$(OBJ_EXT): ../src/api/api_parsers.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node api/api_ast_vector.h.node cmd_context/cmd_context.h.node smt/smt_solver.h.node parsers/smt2/smt2parser.h.node solver/solver_na2as.h.node
	@echo src/api/api_parsers.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_parsers$(OBJ_EXT) ../src/api/api_parsers.cpp
api/api_seq$(OBJ_EXT): ../src/api/api_seq.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/ast_pp.h.node
	@echo src/api/api_seq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_seq$(OBJ_EXT) ../src/api/api_seq.cpp
api/api_goal$(OBJ_EXT): ../src/api/api_goal.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_goal.h.node ast/ast_translation.h.node api/api_model.h.node
	@echo src/api/api_goal.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_goal$(OBJ_EXT) ../src/api/api_goal.cpp
api/api_opt$(OBJ_EXT): ../src/api/api_opt.cpp util/cancel_eh.h.node util/scoped_timer.h.node util/scoped_ctrl_c.h.node util/file_path.h.node parsers/smt2/smt2parser.h.node model/model_params.hpp.node opt/opt_context.h.node opt/opt_cmds.h.node opt/opt_parse.h.node api/z3.h.node api/api_log_macros.h.node api/api_stats.h.node api/api_context.h.node api/api_util.h.node api/api_model.h.node api/api_ast_vector.h.node
	@echo src/api/api_opt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_opt$(OBJ_EXT) ../src/api/api_opt.cpp
api/api_algebraic$(OBJ_EXT): ../src/api/api_algebraic.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_ast_vector.h.node math/polynomial/algebraic_numbers.h.node ast/expr2polynomial.h.node util/cancel_eh.h.node util/scoped_timer.h.node
	@echo src/api/api_algebraic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_algebraic$(OBJ_EXT) ../src/api/api_algebraic.cpp
api/api_arith$(OBJ_EXT): ../src/api/api_arith.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/arith_decl_plugin.h.node math/polynomial/algebraic_numbers.h.node
	@echo src/api/api_arith.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_arith$(OBJ_EXT) ../src/api/api_arith.cpp
api/api_datatype$(OBJ_EXT): ../src/api/api_datatype.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/datatype_decl_plugin.h.node
	@echo src/api/api_datatype.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_datatype$(OBJ_EXT) ../src/api/api_datatype.cpp
api/api_model$(OBJ_EXT): ../src/api/api_model.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_model.h.node api/api_ast_vector.h.node ast/array_decl_plugin.h.node model/model.h.node model/model_v2_pp.h.node model/model_smt2_pp.h.node model/model_params.hpp.node model/model_evaluator_params.hpp.node
	@echo src/api/api_model.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_model$(OBJ_EXT) ../src/api/api_model.cpp
api/api_params$(OBJ_EXT): ../src/api/api_params.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node util/params.h.node
	@echo src/api/api_params.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_params$(OBJ_EXT) ../src/api/api_params.cpp
api/api_fpa$(OBJ_EXT): ../src/api/api_fpa.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node ast/fpa_decl_plugin.h.node
	@echo src/api/api_fpa.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_fpa$(OBJ_EXT) ../src/api/api_fpa.cpp
api/api_polynomial$(OBJ_EXT): ../src/api/api_polynomial.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_polynomial.h.node api/api_ast_vector.h.node ast/expr2polynomial.h.node util/cancel_eh.h.node util/scoped_timer.h.node ast/expr2var.h.node
	@echo src/api/api_polynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_polynomial$(OBJ_EXT) ../src/api/api_polynomial.cpp
api/api_bv$(OBJ_EXT): ../src/api/api_bv.cpp api/z3.h.node api/api_log_macros.h.node api/api_context.h.node api/api_util.h.node ast/bv_decl_plugin.h.node
	@echo src/api/api_bv.cpp
	@$(CXX) $(CXXFLAGS) $(includes_64) $(CXX_OUT_FLAG)api/api_bv$(OBJ_EXT) ../src/api/api_bv.cpp
api/api$(LIB_EXT): api/z3_replayer$(OBJ_EXT) api/api_ast_vector$(OBJ_EXT) api/api_log_macros$(OBJ_EXT) api/api_context$(OBJ_EXT) api/api_stats$(OBJ_EXT) api/api_rcf$(OBJ_EXT) api/api_tactic$(OBJ_EXT) api/api_special_relations$(OBJ_EXT) api/api_commands$(OBJ_EXT) api/api_ast$(OBJ_EXT) api/api_pb$(OBJ_EXT) api/api_ast_map$(OBJ_EXT) api/api_config_params$(OBJ_EXT) api/api_log$(OBJ_EXT) api/api_quant$(OBJ_EXT) api/api_numeral$(OBJ_EXT) api/api_datalog$(OBJ_EXT) api/api_array$(OBJ_EXT) api/api_solver$(OBJ_EXT) api/api_qe$(OBJ_EXT) api/api_parsers$(OBJ_EXT) api/api_seq$(OBJ_EXT) api/api_goal$(OBJ_EXT) api/api_opt$(OBJ_EXT) api/api_algebraic$(OBJ_EXT) api/api_arith$(OBJ_EXT) api/api_datatype$(OBJ_EXT) api/api_model$(OBJ_EXT) api/api_params$(OBJ_EXT) api/api_fpa$(OBJ_EXT) api/api_polynomial$(OBJ_EXT) api/api_bv$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)api/api$(LIB_EXT) api/z3_replayer$(OBJ_EXT) api/api_ast_vector$(OBJ_EXT) api/api_log_macros$(OBJ_EXT) api/api_context$(OBJ_EXT) api/api_stats$(OBJ_EXT) api/api_rcf$(OBJ_EXT) api/api_tactic$(OBJ_EXT) api/api_special_relations$(OBJ_EXT) api/api_commands$(OBJ_EXT) api/api_ast$(OBJ_EXT) api/api_pb$(OBJ_EXT) api/api_ast_map$(OBJ_EXT) api/api_config_params$(OBJ_EXT) api/api_log$(OBJ_EXT) api/api_quant$(OBJ_EXT) api/api_numeral$(OBJ_EXT) api/api_datalog$(OBJ_EXT) api/api_array$(OBJ_EXT) api/api_solver$(OBJ_EXT) api/api_qe$(OBJ_EXT) api/api_parsers$(OBJ_EXT) api/api_seq$(OBJ_EXT) api/api_goal$(OBJ_EXT) api/api_opt$(OBJ_EXT) api/api_algebraic$(OBJ_EXT) api/api_arith$(OBJ_EXT) api/api_datatype$(OBJ_EXT) api/api_model$(OBJ_EXT) api/api_params$(OBJ_EXT) api/api_fpa$(OBJ_EXT) api/api_polynomial$(OBJ_EXT) api/api_bv$(OBJ_EXT)
api: api/api$(LIB_EXT)

includes_65 = -I../src/cmd_context -I../src/solver -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/tactic -I../src/ast/proofs -I../src/math/subpaving/tactic -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/parsers/util -I../src/math/subpaving -I../src/math/interval -I../src/qe -I../src/smt -I../src/ast/rewriter/bit_blaster -I../src/smt/proto_model -I../src/smt/params -I../src/ast/substitution -I../src/math/grobner -I../src/math/dd -I../src/math/simplex -I../src/ast/fpa -I../src/math/lp -I../src/nlsat -I../src/sat -I../src/qe/mbp -I../src/nlsat/tactic -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/tactic/arith -I../src
math/polynomial/polynomial_var2value.h.node: ../src/math/polynomial/polynomial_var2value.h math/polynomial/polynomial.h.node util/scoped_numeral_vector.h.node
	@echo done > math/polynomial/polynomial_var2value.h.node
cmd_context/extra_cmds/polynomial_cmds$(OBJ_EXT): ../src/cmd_context/extra_cmds/polynomial_cmds.cpp cmd_context/cmd_context.h.node cmd_context/cmd_util.h.node util/scoped_timer.h.node util/scoped_ctrl_c.h.node util/cancel_eh.h.node ast/ast_smt2_pp.h.node ast/expr2polynomial.h.node cmd_context/parametric_cmd.h.node util/mpq.h.node math/polynomial/algebraic_numbers.h.node math/polynomial/polynomial_var2value.h.node ast/expr2var.h.node ast/pp.h.node ast/pp_params.hpp.node
	@echo src/cmd_context/extra_cmds/polynomial_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_65) $(CXX_OUT_FLAG)cmd_context/extra_cmds/polynomial_cmds$(OBJ_EXT) ../src/cmd_context/extra_cmds/polynomial_cmds.cpp
cmd_context/extra_cmds/dbg_cmds$(OBJ_EXT): ../src/cmd_context/extra_cmds/dbg_cmds.cpp cmd_context/cmd_context.h.node cmd_context/cmd_util.h.node ast/rewriter/rewriter.h.node ast/shared_occs.h.node ast/for_each_expr.h.node ast/rewriter/rewriter.h.node ast/rewriter/bool_rewriter.h.node ast/ast_lt.h.node cmd_context/simplify_cmd.h.node ast/ast_smt2_pp.h.node tactic/arith/bound_manager.h.node ast/used_vars.h.node ast/rewriter/var_subst.h.node ast/ast_util.h.node util/gparams.h.node qe/qe_mbp.h.node qe/qe_mbi.h.node qe/mbp/mbp_term_graph.h.node
	@echo src/cmd_context/extra_cmds/dbg_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_65) $(CXX_OUT_FLAG)cmd_context/extra_cmds/dbg_cmds$(OBJ_EXT) ../src/cmd_context/extra_cmds/dbg_cmds.cpp
cmd_context/extra_cmds/subpaving_cmds$(OBJ_EXT): ../src/cmd_context/extra_cmds/subpaving_cmds.cpp cmd_context/cmd_context.h.node cmd_context/cmd_util.h.node math/subpaving/tactic/expr2subpaving.h.node ast/rewriter/th_rewriter.h.node ast/ast_smt2_pp.h.node ast/expr2var.h.node
	@echo src/cmd_context/extra_cmds/subpaving_cmds.cpp
	@$(CXX) $(CXXFLAGS) $(includes_65) $(CXX_OUT_FLAG)cmd_context/extra_cmds/subpaving_cmds$(OBJ_EXT) ../src/cmd_context/extra_cmds/subpaving_cmds.cpp
cmd_context/extra_cmds/extra_cmds$(LIB_EXT): cmd_context/extra_cmds/polynomial_cmds$(OBJ_EXT) cmd_context/extra_cmds/dbg_cmds$(OBJ_EXT) cmd_context/extra_cmds/subpaving_cmds$(OBJ_EXT)
	@$(AR) $(AR_FLAGS) $(AR_OUTFLAG)cmd_context/extra_cmds/extra_cmds$(LIB_EXT) cmd_context/extra_cmds/polynomial_cmds$(OBJ_EXT) cmd_context/extra_cmds/dbg_cmds$(OBJ_EXT) cmd_context/extra_cmds/subpaving_cmds$(OBJ_EXT)
extra_cmds: cmd_context/extra_cmds/extra_cmds$(LIB_EXT)

includes_66 = -I../src/api -I../src/tactic/portfolio -I../src/tactic/smtlogics -I../src/ackermannization -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src/sat/sat_solver -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/nlsat/tactic -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src/tactic/ufbv -I../src/tactic/fpa -I../src/tactic/sls -I../src/math/subpaving/tactic -I../src/math/subpaving -I../src/math/realclosure -I../src/opt -I../src/cmd_context/extra_cmds -I../src
shell/opt_frontend.h.node: ../src/shell/opt_frontend.h
	@echo done > shell/opt_frontend.h.node
shell/opt_frontend$(OBJ_EXT): ../src/shell/opt_frontend.cpp util/gparams.h.node util/timeout.h.node util/cancel_eh.h.node util/scoped_timer.h.node util/mutex.h.node ast/ast_util.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node model/model_smt2_pp.h.node opt/opt_context.h.node shell/opt_frontend.h.node opt/opt_parse.h.node
	@echo src/shell/opt_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/opt_frontend$(OBJ_EXT) ../src/shell/opt_frontend.cpp
shell/datalog_frontend.h.node: ../src/shell/datalog_frontend.h
	@echo done > shell/datalog_frontend.h.node
shell/datalog_frontend$(OBJ_EXT): ../src/shell/datalog_frontend.cpp util/stopwatch.h.node smt/params/smt_params.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node muz/rel/dl_compiler.h.node muz/transforms/dl_mk_filter_rules.h.node muz/rel/dl_finite_product_relation.h.node muz/base/dl_context.h.node muz/rel/rel_context.h.node muz/fp/dl_register_engine.h.node muz/fp/datalog_parser.h.node shell/datalog_frontend.h.node util/timeout.h.node util/mutex.h.node
	@echo src/shell/datalog_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/datalog_frontend$(OBJ_EXT) ../src/shell/datalog_frontend.cpp
math/subpaving/tactic/subpaving_tactic.h.node: ../src/math/subpaving/tactic/subpaving_tactic.h util/params.h.node
	@echo done > math/subpaving/tactic/subpaving_tactic.h.node
shell/install_tactic$(OBJ_EXT): ../src/shell/install_tactic.cpp tactic/tactic.h.node cmd_context/tactic_cmds.h.node cmd_context/cmd_context.h.node ackermannization/ackermannize_bv_tactic.h.node ackermannization/ackr_bound_probe.h.node math/subpaving/tactic/subpaving_tactic.h.node muz/fp/horn_tactic.h.node nlsat/tactic/nlsat_tactic.h.node nlsat/tactic/qfnra_nlsat_tactic.h.node qe/nlqsat.h.node qe/qe_lite.h.node qe/qe_tactic.h.node qe/qsat.h.node sat/sat_solver/inc_sat_solver.h.node sat/tactic/sat_tactic.h.node smt/tactic/ctx_solver_simplify_tactic.h.node smt/tactic/smt_tactic.h.node smt/tactic/unit_subsumption_tactic.h.node tactic/aig/aig_tactic.h.node tactic/arith/add_bounds_tactic.h.node tactic/arith/card2bv_tactic.h.node tactic/arith/degree_shift_tactic.h.node tactic/arith/diff_neq_tactic.h.node tactic/arith/eq2bv_tactic.h.node tactic/arith/factor_tactic.h.node tactic/arith/fix_dl_var_tactic.h.node tactic/arith/fm_tactic.h.node tactic/arith/lia2card_tactic.h.node tactic/arith/lia2pb_tactic.h.node tactic/arith/nla2bv_tactic.h.node tactic/arith/normalize_bounds_tactic.h.node tactic/arith/pb2bv_tactic.h.node tactic/arith/probe_arith.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/arith/purify_arith_tactic.h.node tactic/arith/recover_01_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/bv1_blaster_tactic.h.node tactic/bv/bv_bound_chk_tactic.h.node tactic/bv/bv_bounds_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/bv/bvarray2uf_tactic.h.node tactic/bv/dt2bv_tactic.h.node tactic/bv/elim_small_bv_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/core/blast_term_ite_tactic.h.node tactic/core/cofactor_term_ite_tactic.h.node tactic/core/collect_statistics_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/core/der_tactic.h.node tactic/core/distribute_forall_tactic.h.node tactic/core/dom_simplify_tactic.h.node tactic/core/elim_term_ite_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/core/injectivity_tactic.h.node tactic/core/nnf_tactic.h.node tactic/core/occf_tactic.h.node tactic/core/pb_preprocess_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/reduce_args_tactic.h.node tactic/core/reduce_invertible_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/special_relations_tactic.h.node tactic/core/split_clause_tactic.h.node tactic/core/symmetry_reduce_tactic.h.node tactic/core/tseitin_cnf_tactic.h.node tactic/fd_solver/fd_solver.h.node tactic/fd_solver/smtfd_solver.h.node tactic/fpa/fpa2bv_tactic.h.node tactic/fpa/qffp_tactic.h.node tactic/fpa/qffplra_tactic.h.node tactic/portfolio/default_tactic.h.node tactic/probe.h.node tactic/sine_filter.h.node tactic/sls/sls_tactic.h.node tactic/smtlogics/nra_tactic.h.node tactic/smtlogics/qfaufbv_tactic.h.node tactic/smtlogics/qfauflia_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/qfidl_tactic.h.node tactic/smtlogics/qflia_tactic.h.node tactic/smtlogics/qflra_tactic.h.node tactic/smtlogics/qfnia_tactic.h.node tactic/smtlogics/qfnra_tactic.h.node tactic/smtlogics/qfuf_tactic.h.node tactic/smtlogics/qfufbv_tactic.h.node tactic/smtlogics/quant_tactics.h.node tactic/tactic.h.node tactic/ufbv/macro_finder_tactic.h.node tactic/ufbv/quasi_macros_tactic.h.node tactic/ufbv/ufbv_rewriter_tactic.h.node tactic/ufbv/ufbv_tactic.h.node
	@echo src/shell/install_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/install_tactic$(OBJ_EXT) ../src/shell/install_tactic.cpp
shell/drat_frontend.h.node: ../src/shell/drat_frontend.h
	@echo done > shell/drat_frontend.h.node
shell/drat_frontend$(OBJ_EXT): ../src/shell/drat_frontend.cpp ast/bv_decl_plugin.h.node util/memory_manager.h.node util/statistics.h.node sat/dimacs.h.node sat/sat_solver.h.node sat/sat_drat.h.node smt/smt_solver.h.node shell/drat_frontend.h.node parsers/smt2/smt2parser.h.node cmd_context/cmd_context.h.node
	@echo src/shell/drat_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/drat_frontend$(OBJ_EXT) ../src/shell/drat_frontend.cpp
shell/gparams_register_modules$(OBJ_EXT): ../src/shell/gparams_register_modules.cpp util/gparams.h.node ackermannization/ackermannization_params.hpp.node ackermannization/ackermannize_bv_tactic_params.hpp.node ast/normal_forms/nnf.h.node ast/normal_forms/nnf_params.hpp.node ast/pp_params.hpp.node math/polynomial/algebraic_params.hpp.node math/realclosure/rcf_params.hpp.node model/model_evaluator_params.hpp.node model/model_params.hpp.node muz/base/fp_params.hpp.node nlsat/nlsat_params.hpp.node opt/opt_params.hpp.node params/arith_rewriter_params.hpp.node params/array_rewriter_params.hpp.node params/bool_rewriter_params.hpp.node params/bv_rewriter_params.hpp.node params/context_params.h.node params/fpa2bv_rewriter_params.hpp.node params/fpa_rewriter_params.hpp.node params/pattern_inference_params_helper.hpp.node params/poly_rewriter_params.hpp.node params/rewriter_params.hpp.node params/seq_rewriter_params.hpp.node parsers/util/parser_params.hpp.node sat/sat_asymm_branch_params.hpp.node sat/sat_params.hpp.node sat/sat_scc_params.hpp.node sat/sat_simplifier_params.hpp.node smt/params/smt_params_helper.hpp.node solver/combined_solver_params.hpp.node solver/parallel_params.hpp.node solver/solver_params.hpp.node tactic/sls/sls_params.hpp.node tactic/smtlogics/qfufbv_tactic_params.hpp.node tactic/tactic_params.hpp.node util/env_params.h.node
	@echo src/shell/gparams_register_modules.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/gparams_register_modules$(OBJ_EXT) ../src/shell/gparams_register_modules.cpp
shell/mem_initializer$(OBJ_EXT): ../src/shell/mem_initializer.cpp util/debug.h.node util/gparams.h.node util/prime_generator.h.node util/rational.h.node util/rlimit.h.node util/scoped_timer.h.node util/symbol.h.node util/trace.h.node
	@echo src/shell/mem_initializer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/mem_initializer$(OBJ_EXT) ../src/shell/mem_initializer.cpp
shell/z3_log_frontend$(OBJ_EXT): ../src/shell/z3_log_frontend.cpp util/util.h.node util/error_codes.h.node api/z3_replayer.h.node
	@echo src/shell/z3_log_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/z3_log_frontend$(OBJ_EXT) ../src/shell/z3_log_frontend.cpp
math/lp/mps_reader.h.node: ../src/math/lp/mps_reader.h util/vector.h.node math/lp/lp_primal_simplex.h.node math/lp/lp_dual_simplex.h.node math/lp/lar_solver.h.node math/lp/lp_utils.h.node math/lp/lp_solver.h.node
	@echo done > math/lp/mps_reader.h.node
shell/lp_frontend$(OBJ_EXT): ../src/shell/lp_frontend.cpp math/lp/lp_settings.h.node math/lp/mps_reader.h.node util/timeout.h.node util/cancel_eh.h.node util/scoped_timer.h.node util/rlimit.h.node util/gparams.h.node util/mutex.h.node smt/params/smt_params_helper.hpp.node
	@echo src/shell/lp_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/lp_frontend$(OBJ_EXT) ../src/shell/lp_frontend.cpp
shell/smtlib_frontend.h.node: ../src/shell/smtlib_frontend.h
	@echo done > shell/smtlib_frontend.h.node
shell/z3_log_frontend.h.node: ../src/shell/z3_log_frontend.h
	@echo done > shell/z3_log_frontend.h.node
shell/dimacs_frontend.h.node: ../src/shell/dimacs_frontend.h
	@echo done > shell/dimacs_frontend.h.node
shell/lp_frontend.h.node: ../src/shell/lp_frontend.h
	@echo done > shell/lp_frontend.h.node
shell/main$(OBJ_EXT): ../src/shell/main.cpp util/memory_manager.h.node util/trace.h.node util/debug.h.node util/util.h.node ast/pp.h.node shell/smtlib_frontend.h.node shell/z3_log_frontend.h.node util/warning.h.node util/z3_version.h.node shell/dimacs_frontend.h.node shell/datalog_frontend.h.node shell/opt_frontend.h.node util/timeout.h.node util/z3_exception.h.node util/error_codes.h.node util/file_path.h.node util/gparams.h.node util/env_params.h.node util/file_path.h.node shell/lp_frontend.h.node shell/drat_frontend.h.node
	@echo src/shell/main.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/main$(OBJ_EXT) ../src/shell/main.cpp
cmd_context/extra_cmds/dbg_cmds.h.node: ../src/cmd_context/extra_cmds/dbg_cmds.h
	@echo done > cmd_context/extra_cmds/dbg_cmds.h.node
cmd_context/extra_cmds/polynomial_cmds.h.node: ../src/cmd_context/extra_cmds/polynomial_cmds.h
	@echo done > cmd_context/extra_cmds/polynomial_cmds.h.node
cmd_context/extra_cmds/subpaving_cmds.h.node: ../src/cmd_context/extra_cmds/subpaving_cmds.h
	@echo done > cmd_context/extra_cmds/subpaving_cmds.h.node
shell/smtlib_frontend$(OBJ_EXT): ../src/shell/smtlib_frontend.cpp util/timeout.h.node util/mutex.h.node parsers/smt2/smt2parser.h.node muz/fp/dl_cmds.h.node cmd_context/extra_cmds/dbg_cmds.h.node opt/opt_cmds.h.node cmd_context/extra_cmds/polynomial_cmds.h.node cmd_context/extra_cmds/subpaving_cmds.h.node smt/smt2_extra_cmds.h.node smt/smt_solver.h.node
	@echo src/shell/smtlib_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/smtlib_frontend$(OBJ_EXT) ../src/shell/smtlib_frontend.cpp
shell/dimacs_frontend$(OBJ_EXT): ../src/shell/dimacs_frontend.cpp util/timeout.h.node util/rlimit.h.node util/gparams.h.node sat/dimacs.h.node sat/sat_params.hpp.node sat/sat_solver.h.node sat/tactic/goal2sat.h.node ast/reg_decl_plugins.h.node tactic/tactic.h.node tactic/fd_solver/fd_solver.h.node
	@echo src/shell/dimacs_frontend.cpp
	@$(CXX) $(CXXFLAGS) $(includes_66) $(CXX_OUT_FLAG)shell/dimacs_frontend$(OBJ_EXT) ../src/shell/dimacs_frontend.cpp
z3$(EXE_EXT): shell/opt_frontend$(OBJ_EXT) shell/datalog_frontend$(OBJ_EXT) shell/install_tactic$(OBJ_EXT) shell/drat_frontend$(OBJ_EXT) shell/gparams_register_modules$(OBJ_EXT) shell/mem_initializer$(OBJ_EXT) shell/z3_log_frontend$(OBJ_EXT) shell/lp_frontend$(OBJ_EXT) shell/main$(OBJ_EXT) shell/smtlib_frontend$(OBJ_EXT) shell/dimacs_frontend$(OBJ_EXT) cmd_context/extra_cmds/extra_cmds$(LIB_EXT) api/api$(LIB_EXT) opt/opt$(LIB_EXT) tactic/portfolio/portfolio$(LIB_EXT) tactic/fpa/fpa_tactics$(LIB_EXT) tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/ufbv/ufbv_tactic$(LIB_EXT) muz/fp/fp$(LIB_EXT) muz/bmc/bmc$(LIB_EXT) muz/ddnf/ddnf$(LIB_EXT) muz/tab/tab$(LIB_EXT) muz/clp/clp$(LIB_EXT) muz/spacer/spacer$(LIB_EXT) muz/rel/rel$(LIB_EXT) muz/transforms/transforms$(LIB_EXT) muz/dataflow/dataflow$(LIB_EXT) muz/base/muz$(LIB_EXT) tactic/fd_solver/fd_solver$(LIB_EXT) sat/sat_solver/sat_solver$(LIB_EXT) qe/qe$(LIB_EXT) tactic/sls/sls_tactic$(LIB_EXT) smt/tactic/smt_tactic$(LIB_EXT) tactic/bv/bv_tactics$(LIB_EXT) smt/smt$(LIB_EXT) smt/proto_model/proto_model$(LIB_EXT) math/subpaving/tactic/subpaving_tactic$(LIB_EXT) nlsat/tactic/nlsat_tactic$(LIB_EXT) sat/tactic/sat_tactic$(LIB_EXT) sat/smt/sat_smt$(LIB_EXT) qe/mbp/mbp$(LIB_EXT) tactic/arith/arith_tactics$(LIB_EXT) tactic/core/core_tactics$(LIB_EXT) smt/params/smt_params$(LIB_EXT) ast/pattern/pattern$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/fpa/fpa$(LIB_EXT) ackermannization/ackermannization$(LIB_EXT) tactic/aig/aig_tactic$(LIB_EXT) parsers/smt2/smt2parser$(LIB_EXT) cmd_context/cmd_context$(LIB_EXT) solver/solver$(LIB_EXT) ast/proofs/proofs$(LIB_EXT) parsers/util/parser_util$(LIB_EXT) ast/substitution/substitution$(LIB_EXT) tactic/tactic$(LIB_EXT) model/model$(LIB_EXT) ast/normal_forms/normal_forms$(LIB_EXT) ast/macros/macros$(LIB_EXT) ast/rewriter/rewriter$(LIB_EXT) math/lp/lp$(LIB_EXT) nlsat/nlsat$(LIB_EXT) sat/sat$(LIB_EXT) math/grobner/grobner$(LIB_EXT) ast/euf/euf$(LIB_EXT) params/params$(LIB_EXT) ast/ast$(LIB_EXT) math/subpaving/subpaving$(LIB_EXT) math/realclosure/realclosure$(LIB_EXT) math/automata/automata$(LIB_EXT) math/hilbert/hilbert$(LIB_EXT) math/simplex/simplex$(LIB_EXT) math/dd/dd$(LIB_EXT) math/interval/interval$(LIB_EXT) math/polynomial/polynomial$(LIB_EXT) util/util$(LIB_EXT)
	$(LINK)  $(LINK_OUT_FLAG)z3$(EXE_EXT) $(LINK_FLAGS) shell/opt_frontend$(OBJ_EXT) shell/datalog_frontend$(OBJ_EXT) shell/install_tactic$(OBJ_EXT) shell/drat_frontend$(OBJ_EXT) shell/gparams_register_modules$(OBJ_EXT) shell/mem_initializer$(OBJ_EXT) shell/z3_log_frontend$(OBJ_EXT) shell/lp_frontend$(OBJ_EXT) shell/main$(OBJ_EXT) shell/smtlib_frontend$(OBJ_EXT) shell/dimacs_frontend$(OBJ_EXT) cmd_context/extra_cmds/extra_cmds$(LIB_EXT) api/api$(LIB_EXT) opt/opt$(LIB_EXT) tactic/portfolio/portfolio$(LIB_EXT) tactic/fpa/fpa_tactics$(LIB_EXT) tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/ufbv/ufbv_tactic$(LIB_EXT) muz/fp/fp$(LIB_EXT) muz/bmc/bmc$(LIB_EXT) muz/ddnf/ddnf$(LIB_EXT) muz/tab/tab$(LIB_EXT) muz/clp/clp$(LIB_EXT) muz/spacer/spacer$(LIB_EXT) muz/rel/rel$(LIB_EXT) muz/transforms/transforms$(LIB_EXT) muz/dataflow/dataflow$(LIB_EXT) muz/base/muz$(LIB_EXT) tactic/fd_solver/fd_solver$(LIB_EXT) sat/sat_solver/sat_solver$(LIB_EXT) qe/qe$(LIB_EXT) tactic/sls/sls_tactic$(LIB_EXT) smt/tactic/smt_tactic$(LIB_EXT) tactic/bv/bv_tactics$(LIB_EXT) smt/smt$(LIB_EXT) smt/proto_model/proto_model$(LIB_EXT) math/subpaving/tactic/subpaving_tactic$(LIB_EXT) nlsat/tactic/nlsat_tactic$(LIB_EXT) sat/tactic/sat_tactic$(LIB_EXT) sat/smt/sat_smt$(LIB_EXT) qe/mbp/mbp$(LIB_EXT) tactic/arith/arith_tactics$(LIB_EXT) tactic/core/core_tactics$(LIB_EXT) smt/params/smt_params$(LIB_EXT) ast/pattern/pattern$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/fpa/fpa$(LIB_EXT) ackermannization/ackermannization$(LIB_EXT) tactic/aig/aig_tactic$(LIB_EXT) parsers/smt2/smt2parser$(LIB_EXT) cmd_context/cmd_context$(LIB_EXT) solver/solver$(LIB_EXT) ast/proofs/proofs$(LIB_EXT) parsers/util/parser_util$(LIB_EXT) ast/substitution/substitution$(LIB_EXT) tactic/tactic$(LIB_EXT) model/model$(LIB_EXT) ast/normal_forms/normal_forms$(LIB_EXT) ast/macros/macros$(LIB_EXT) ast/rewriter/rewriter$(LIB_EXT) math/lp/lp$(LIB_EXT) nlsat/nlsat$(LIB_EXT) sat/sat$(LIB_EXT) math/grobner/grobner$(LIB_EXT) ast/euf/euf$(LIB_EXT) params/params$(LIB_EXT) ast/ast$(LIB_EXT) math/subpaving/subpaving$(LIB_EXT) math/realclosure/realclosure$(LIB_EXT) math/automata/automata$(LIB_EXT) math/hilbert/hilbert$(LIB_EXT) math/simplex/simplex$(LIB_EXT) math/dd/dd$(LIB_EXT) math/interval/interval$(LIB_EXT) math/polynomial/polynomial$(LIB_EXT) util/util$(LIB_EXT) $(LINK_EXTRA_FLAGS)
shell: z3$(EXE_EXT)

includes_67 = -I../src/api -I../src/tactic/portfolio -I../src/tactic/smtlogics -I../src/ackermannization -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src/sat/sat_solver -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/nlsat/tactic -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src/tactic/ufbv -I../src/tactic/fpa -I../src/tactic/sls -I../src/math/subpaving/tactic -I../src/math/subpaving -I../src/math/realclosure -I../src/opt -I../src/test/fuzzing -I../src
test/egraph$(OBJ_EXT): ../src/test/egraph.cpp util/util.h.node util/timer.h.node ast/euf/euf_egraph.h.node ast/reg_decl_plugins.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node
	@echo src/test/egraph.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/egraph$(OBJ_EXT) ../src/test/egraph.cpp
test/datalog_parser$(OBJ_EXT): ../src/test/datalog_parser.cpp muz/fp/datalog_parser.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node smt/params/smt_params.h.node ast/reg_decl_plugins.h.node
	@echo src/test/datalog_parser.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/datalog_parser$(OBJ_EXT) ../src/test/datalog_parser.cpp
test/string_buffer$(OBJ_EXT): ../src/test/string_buffer.cpp util/debug.h.node util/string_buffer.h.node util/trace.h.node
	@echo src/test/string_buffer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/string_buffer$(OBJ_EXT) ../src/test/string_buffer.cpp
test/heap_trie$(OBJ_EXT): ../src/test/heap_trie.cpp math/hilbert/heap_trie.h.node
	@echo src/test/heap_trie.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/heap_trie$(OBJ_EXT) ../src/test/heap_trie.cpp
test/proof_checker$(OBJ_EXT): ../src/test/proof_checker.cpp ast/proofs/proof_checker.h.node ast/ast_ll_pp.h.node
	@echo src/test/proof_checker.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/proof_checker$(OBJ_EXT) ../src/test/proof_checker.cpp
test/mpbq$(OBJ_EXT): ../src/test/mpbq.cpp util/mpbq.h.node
	@echo src/test/mpbq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mpbq$(OBJ_EXT) ../src/test/mpbq.cpp
test/bit_vector$(OBJ_EXT): ../src/test/bit_vector.cpp util/bit_vector.h.node util/vector.h.node
	@echo src/test/bit_vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/bit_vector$(OBJ_EXT) ../src/test/bit_vector.cpp
test/ext_numeral$(OBJ_EXT): ../src/test/ext_numeral.cpp util/mpq.h.node util/ext_numeral.h.node
	@echo src/test/ext_numeral.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/ext_numeral$(OBJ_EXT) ../src/test/ext_numeral.cpp
test/quant_elim$(OBJ_EXT): ../src/test/quant_elim.cpp ast/ast.h.node smt/params/smt_params.h.node qe/qe.h.node ast/ast_pp.h.node util/lbool.h.node ast/reg_decl_plugins.h.node
	@echo src/test/quant_elim.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/quant_elim$(OBJ_EXT) ../src/test/quant_elim.cpp
test/map$(OBJ_EXT): ../src/test/map.cpp util/map.h.node util/str_hashtable.h.node
	@echo src/test/map.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/map$(OBJ_EXT) ../src/test/map.cpp
api/z3_private.h.node: ../src/api/z3_private.h util/rational.h.node api/z3_macros.h.node
	@echo done > api/z3_private.h.node
test/memory$(OBJ_EXT): ../src/test/memory.cpp api/z3.h.node api/z3_private.h.node util/util.h.node util/trace.h.node
	@echo src/test/memory.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/memory$(OBJ_EXT) ../src/test/memory.cpp
test/dl_product_relation$(OBJ_EXT): ../src/test/dl_product_relation.cpp ast/reg_decl_plugins.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node muz/rel/dl_finite_product_relation.h.node muz/rel/dl_sparse_table.h.node muz/rel/rel_context.h.node
	@echo src/test/dl_product_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/dl_product_relation$(OBJ_EXT) ../src/test/dl_product_relation.cpp
test/mpq$(OBJ_EXT): ../src/test/mpq.cpp util/mpq.h.node util/rational.h.node util/timeit.h.node
	@echo src/test/mpq.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mpq$(OBJ_EXT) ../src/test/mpq.cpp
test/model_retrieval$(OBJ_EXT): ../src/test/model_retrieval.cpp ast/ast.h.node smt/params/smt_params.h.node smt/smt_context.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node model/model_v2_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/model_retrieval.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/model_retrieval$(OBJ_EXT) ../src/test/model_retrieval.cpp
test/finder$(OBJ_EXT): ../src/test/finder.cpp sat/sat_solver.h.node sat/sat_npn3_finder.h.node
	@echo src/test/finder.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/finder$(OBJ_EXT) ../src/test/finder.cpp
test/dl_util$(OBJ_EXT): ../src/test/dl_util.cpp muz/base/dl_util.h.node
	@echo src/test/dl_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/dl_util$(OBJ_EXT) ../src/test/dl_util.cpp
test/stack$(OBJ_EXT): ../src/test/stack.cpp util/stack.h.node util/vector.h.node
	@echo src/test/stack.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/stack$(OBJ_EXT) ../src/test/stack.cpp
test/install_tactic$(OBJ_EXT): ../src/test/install_tactic.cpp tactic/tactic.h.node cmd_context/tactic_cmds.h.node cmd_context/cmd_context.h.node ackermannization/ackermannize_bv_tactic.h.node ackermannization/ackr_bound_probe.h.node math/subpaving/tactic/subpaving_tactic.h.node muz/fp/horn_tactic.h.node nlsat/tactic/nlsat_tactic.h.node nlsat/tactic/qfnra_nlsat_tactic.h.node qe/nlqsat.h.node qe/qe_lite.h.node qe/qe_tactic.h.node qe/qsat.h.node sat/sat_solver/inc_sat_solver.h.node sat/tactic/sat_tactic.h.node smt/tactic/ctx_solver_simplify_tactic.h.node smt/tactic/smt_tactic.h.node smt/tactic/unit_subsumption_tactic.h.node tactic/aig/aig_tactic.h.node tactic/arith/add_bounds_tactic.h.node tactic/arith/card2bv_tactic.h.node tactic/arith/degree_shift_tactic.h.node tactic/arith/diff_neq_tactic.h.node tactic/arith/eq2bv_tactic.h.node tactic/arith/factor_tactic.h.node tactic/arith/fix_dl_var_tactic.h.node tactic/arith/fm_tactic.h.node tactic/arith/lia2card_tactic.h.node tactic/arith/lia2pb_tactic.h.node tactic/arith/nla2bv_tactic.h.node tactic/arith/normalize_bounds_tactic.h.node tactic/arith/pb2bv_tactic.h.node tactic/arith/probe_arith.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/arith/purify_arith_tactic.h.node tactic/arith/recover_01_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/bv1_blaster_tactic.h.node tactic/bv/bv_bound_chk_tactic.h.node tactic/bv/bv_bounds_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/bv/bvarray2uf_tactic.h.node tactic/bv/dt2bv_tactic.h.node tactic/bv/elim_small_bv_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/core/blast_term_ite_tactic.h.node tactic/core/cofactor_term_ite_tactic.h.node tactic/core/collect_statistics_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/core/der_tactic.h.node tactic/core/distribute_forall_tactic.h.node tactic/core/dom_simplify_tactic.h.node tactic/core/elim_term_ite_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/core/injectivity_tactic.h.node tactic/core/nnf_tactic.h.node tactic/core/occf_tactic.h.node tactic/core/pb_preprocess_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/reduce_args_tactic.h.node tactic/core/reduce_invertible_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/special_relations_tactic.h.node tactic/core/split_clause_tactic.h.node tactic/core/symmetry_reduce_tactic.h.node tactic/core/tseitin_cnf_tactic.h.node tactic/fd_solver/fd_solver.h.node tactic/fd_solver/smtfd_solver.h.node tactic/fpa/fpa2bv_tactic.h.node tactic/fpa/qffp_tactic.h.node tactic/fpa/qffplra_tactic.h.node tactic/portfolio/default_tactic.h.node tactic/probe.h.node tactic/sine_filter.h.node tactic/sls/sls_tactic.h.node tactic/smtlogics/nra_tactic.h.node tactic/smtlogics/qfaufbv_tactic.h.node tactic/smtlogics/qfauflia_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/qfidl_tactic.h.node tactic/smtlogics/qflia_tactic.h.node tactic/smtlogics/qflra_tactic.h.node tactic/smtlogics/qfnia_tactic.h.node tactic/smtlogics/qfnra_tactic.h.node tactic/smtlogics/qfuf_tactic.h.node tactic/smtlogics/qfufbv_tactic.h.node tactic/smtlogics/quant_tactics.h.node tactic/tactic.h.node tactic/ufbv/macro_finder_tactic.h.node tactic/ufbv/quasi_macros_tactic.h.node tactic/ufbv/ufbv_rewriter_tactic.h.node tactic/ufbv/ufbv_tactic.h.node
	@echo src/test/install_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/install_tactic$(OBJ_EXT) ../src/test/install_tactic.cpp
test/simple_parser$(OBJ_EXT): ../src/test/simple_parser.cpp parsers/util/cost_parser.h.node smt/cost_evaluator.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node ast/well_sorted.h.node util/warning.h.node ast/reg_decl_plugins.h.node
	@echo src/test/simple_parser.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/simple_parser$(OBJ_EXT) ../src/test/simple_parser.cpp
test/mpz$(OBJ_EXT): ../src/test/mpz.cpp util/mpz.h.node util/rational.h.node util/timeit.h.node util/scoped_numeral.h.node
	@echo src/test/mpz.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mpz$(OBJ_EXT) ../src/test/mpz.cpp
test/pdd_solver$(OBJ_EXT): ../src/test/pdd_solver.cpp util/rlimit.h.node math/grobner/pdd_solver.h.node ast/bv_decl_plugin.h.node ast/ast_util.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node ast/reg_decl_plugins.h.node tactic/goal.h.node tactic/tactic.h.node tactic/bv/bit_blaster_tactic.h.node
	@echo src/test/pdd_solver.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/pdd_solver$(OBJ_EXT) ../src/test/pdd_solver.cpp
test/rcf$(OBJ_EXT): ../src/test/rcf.cpp math/realclosure/realclosure.h.node math/realclosure/mpz_matrix.h.node util/rlimit.h.node
	@echo src/test/rcf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/rcf$(OBJ_EXT) ../src/test/rcf.cpp
test/model_evaluator$(OBJ_EXT): ../src/test/model_evaluator.cpp model/model.h.node model/model_evaluator.h.node model/model_pp.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node ast/ast_pp.h.node
	@echo src/test/model_evaluator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/model_evaluator$(OBJ_EXT) ../src/test/model_evaluator.cpp
test/list$(OBJ_EXT): ../src/test/list.cpp util/trace.h.node util/util.h.node util/region.h.node util/list.h.node
	@echo src/test/list.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/list$(OBJ_EXT) ../src/test/list.cpp
test/sorting_network$(OBJ_EXT): ../src/test/sorting_network.cpp util/trace.h.node util/vector.h.node util/sorting_network.h.node ast/ast.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node ast/ast_util.h.node model/model_smt2_pp.h.node smt/smt_kernel.h.node smt/params/smt_params.h.node
	@echo src/test/sorting_network.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/sorting_network$(OBJ_EXT) ../src/test/sorting_network.cpp
test/vector$(OBJ_EXT): ../src/test/vector.cpp util/vector.h.node
	@echo src/test/vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/vector$(OBJ_EXT) ../src/test/vector.cpp
test/api$(OBJ_EXT): ../src/test/api.cpp api/z3.h.node api/z3_private.h.node util/util.h.node util/trace.h.node util/trace.h.node
	@echo src/test/api.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/api$(OBJ_EXT) ../src/test/api.cpp
test/qe_arith$(OBJ_EXT): ../src/test/qe_arith.cpp qe/mbp/mbp_arith.h.node qe/qe.h.node ast/rewriter/th_rewriter.h.node parsers/smt2/smt2parser.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node ast/rewriter/arith_rewriter.h.node ast/ast_pp.h.node smt/smt_context.h.node ast/expr_abstract.h.node ast/rewriter/expr_safe_replace.h.node
	@echo src/test/qe_arith.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/qe_arith$(OBJ_EXT) ../src/test/qe_arith.cpp
test/pdd$(OBJ_EXT): ../src/test/pdd.cpp math/dd/dd_pdd.h.node
	@echo src/test/pdd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/pdd$(OBJ_EXT) ../src/test/pdd.cpp
test/parray$(OBJ_EXT): ../src/test/parray.cpp util/parray.h.node util/small_object_allocator.h.node ast/ast.h.node
	@echo src/test/parray.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/parray$(OBJ_EXT) ../src/test/parray.cpp
test/uint_set$(OBJ_EXT): ../src/test/uint_set.cpp util/uint_set.h.node util/vector.h.node util/map.h.node
	@echo src/test/uint_set.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/uint_set$(OBJ_EXT) ../src/test/uint_set.cpp
test/dl_query$(OBJ_EXT): ../src/test/dl_query.cpp muz/fp/datalog_parser.h.node ast/ast_pp.h.node muz/rel/dl_table_relation.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node smt/params/smt_params.h.node util/stopwatch.h.node ast/reg_decl_plugins.h.node muz/rel/dl_relation_manager.h.node
	@echo src/test/dl_query.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/dl_query$(OBJ_EXT) ../src/test/dl_query.cpp
test/heap$(OBJ_EXT): ../src/test/heap.cpp util/util.h.node util/heap.h.node util/trace.h.node util/uint_set.h.node
	@echo src/test/heap.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/heap$(OBJ_EXT) ../src/test/heap.cpp
test/horn_subsume_model_converter$(OBJ_EXT): ../src/test/horn_subsume_model_converter.cpp tactic/horn_subsume_model_converter.h.node ast/arith_decl_plugin.h.node model/model_smt2_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/horn_subsume_model_converter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/horn_subsume_model_converter$(OBJ_EXT) ../src/test/horn_subsume_model_converter.cpp
test/value_generator$(OBJ_EXT): ../src/test/value_generator.cpp ast/value_generator.h.node ast/reg_decl_plugins.h.node ast/ast_pp.h.node ast/datatype_decl_plugin.h.node ast/seq_decl_plugin.h.node ast/array_decl_plugin.h.node
	@echo src/test/value_generator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/value_generator$(OBJ_EXT) ../src/test/value_generator.cpp
test/mpf$(OBJ_EXT): ../src/test/mpf.cpp util/mpf.h.node util/f2n.h.node
	@echo src/test/mpf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mpf$(OBJ_EXT) ../src/test/mpf.cpp
test/prime_generator$(OBJ_EXT): ../src/test/prime_generator.cpp util/mpz.h.node util/prime_generator.h.node
	@echo src/test/prime_generator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/prime_generator$(OBJ_EXT) ../src/test/prime_generator.cpp
test/timeout$(OBJ_EXT): ../src/test/timeout.cpp util/timeout.h.node util/trace.h.node
	@echo src/test/timeout.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/timeout$(OBJ_EXT) ../src/test/timeout.cpp
test/model2expr$(OBJ_EXT): ../src/test/model2expr.cpp model/model2expr.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node model/model_smt2_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/model2expr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/model2expr$(OBJ_EXT) ../src/test/model2expr.cpp
test/small_object_allocator$(OBJ_EXT): ../src/test/small_object_allocator.cpp util/util.h.node util/trace.h.node util/small_object_allocator.h.node
	@echo src/test/small_object_allocator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/small_object_allocator$(OBJ_EXT) ../src/test/small_object_allocator.cpp
test/no_overflow$(OBJ_EXT): ../src/test/no_overflow.cpp api/z3.h.node util/trace.h.node util/rational.h.node
	@echo src/test/no_overflow.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/no_overflow$(OBJ_EXT) ../src/test/no_overflow.cpp
test/get_implied_equalities$(OBJ_EXT): ../src/test/get_implied_equalities.cpp api/z3.h.node util/trace.h.node util/debug.h.node
	@echo src/test/get_implied_equalities.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/get_implied_equalities$(OBJ_EXT) ../src/test/get_implied_equalities.cpp
test/f2n$(OBJ_EXT): ../src/test/f2n.cpp util/f2n.h.node util/hwf.h.node util/mpf.h.node
	@echo src/test/f2n.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/f2n$(OBJ_EXT) ../src/test/f2n.cpp
test/bit_blaster$(OBJ_EXT): ../src/test/bit_blaster.cpp ast/rewriter/bit_blaster/bit_blaster.h.node ast/ast_pp.h.node ast/ast_ll_pp.h.node
	@echo src/test/bit_blaster.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/bit_blaster$(OBJ_EXT) ../src/test/bit_blaster.cpp
test/substitution$(OBJ_EXT): ../src/test/substitution.cpp ast/expr_substitution.h.node smt/params/smt_params.h.node ast/substitution/substitution.h.node ast/substitution/unifier.h.node ast/bv_decl_plugin.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node
	@echo src/test/substitution.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/substitution$(OBJ_EXT) ../src/test/substitution.cpp
test/mpfx$(OBJ_EXT): ../src/test/mpfx.cpp util/mpfx.h.node
	@echo src/test/mpfx.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mpfx$(OBJ_EXT) ../src/test/mpfx.cpp
test/hashtable$(OBJ_EXT): ../src/test/hashtable.cpp util/hashtable.h.node
	@echo src/test/hashtable.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/hashtable$(OBJ_EXT) ../src/test/hashtable.cpp
test/mpff$(OBJ_EXT): ../src/test/mpff.cpp util/mpff.h.node util/mpz.h.node util/mpq.h.node
	@echo src/test/mpff.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mpff$(OBJ_EXT) ../src/test/mpff.cpp
test/nlarith_util$(OBJ_EXT): ../src/test/nlarith_util.cpp qe/nlarith_util.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/nlarith_util.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/nlarith_util$(OBJ_EXT) ../src/test/nlarith_util.cpp
test/nlsat$(OBJ_EXT): ../src/test/nlsat.cpp nlsat/nlsat_assignment.h.node nlsat/nlsat_interval_set.h.node nlsat/nlsat_evaluator.h.node nlsat/nlsat_solver.h.node util/util.h.node nlsat/nlsat_explain.h.node math/polynomial/polynomial_cache.h.node util/rlimit.h.node
	@echo src/test/nlsat.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/nlsat$(OBJ_EXT) ../src/test/nlsat.cpp
test/simplex$(OBJ_EXT): ../src/test/simplex.cpp math/simplex/sparse_matrix_def.h.node math/simplex/simplex.h.node math/simplex/simplex_def.h.node util/mpq_inf.h.node util/vector.h.node util/rational.h.node util/rlimit.h.node
	@echo src/test/simplex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/simplex$(OBJ_EXT) ../src/test/simplex.cpp
test/var_subst$(OBJ_EXT): ../src/test/var_subst.cpp ast/rewriter/var_subst.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/for_each_expr.h.node ast/reg_decl_plugins.h.node
	@echo src/test/var_subst.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/var_subst$(OBJ_EXT) ../src/test/var_subst.cpp
test/arith_rewriter$(OBJ_EXT): ../src/test/arith_rewriter.cpp ast/rewriter/arith_rewriter.h.node ast/bv_decl_plugin.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node ast/rewriter/th_rewriter.h.node model/model.h.node parsers/smt2/smt2parser.h.node
	@echo src/test/arith_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/arith_rewriter$(OBJ_EXT) ../src/test/arith_rewriter.cpp
test/api_bug$(OBJ_EXT): ../src/test/api_bug.cpp api/z3.h.node
	@echo src/test/api_bug.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/api_bug$(OBJ_EXT) ../src/test/api_bug.cpp
test/for_each_file.h.node: ../src/test/for_each_file.h
	@echo done > test/for_each_file.h.node
test/for_each_file$(OBJ_EXT): ../src/test/for_each_file.cpp util/util.h.node test/for_each_file.h.node
	@echo src/test/for_each_file.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/for_each_file$(OBJ_EXT) ../src/test/for_each_file.cpp
test/symbol$(OBJ_EXT): ../src/test/symbol.cpp util/symbol.h.node util/debug.h.node
	@echo src/test/symbol.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/symbol$(OBJ_EXT) ../src/test/symbol.cpp
test/ast$(OBJ_EXT): ../src/test/ast.cpp ast/ast.h.node
	@echo src/test/ast.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/ast$(OBJ_EXT) ../src/test/ast.cpp
test/ex$(OBJ_EXT): ../src/test/ex.cpp util/z3_exception.h.node
	@echo src/test/ex.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/ex$(OBJ_EXT) ../src/test/ex.cpp
test/smt_context$(OBJ_EXT): ../src/test/smt_context.cpp smt/smt_context.h.node ast/reg_decl_plugins.h.node
	@echo src/test/smt_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/smt_context$(OBJ_EXT) ../src/test/smt_context.cpp
test/model_based_opt$(OBJ_EXT): ../src/test/model_based_opt.cpp math/simplex/model_based_opt.h.node util/util.h.node util/uint_set.h.node
	@echo src/test/model_based_opt.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/model_based_opt$(OBJ_EXT) ../src/test/model_based_opt.cpp
util/ptr_scoped_buffer.h.node: ../src/util/ptr_scoped_buffer.h util/util.h.node util/debug.h.node util/buffer.h.node
	@echo done > util/ptr_scoped_buffer.h.node
test/buffer$(OBJ_EXT): ../src/test/buffer.cpp util/ptr_scoped_buffer.h.node
	@echo src/test/buffer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/buffer$(OBJ_EXT) ../src/test/buffer.cpp
test/dl_context$(OBJ_EXT): ../src/test/dl_context.cpp ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node muz/fp/datalog_parser.h.node muz/base/dl_context.h.node smt/params/smt_params.h.node muz/fp/dl_register_engine.h.node
	@echo src/test/dl_context.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/dl_context$(OBJ_EXT) ../src/test/dl_context.cpp
test/quant_solve$(OBJ_EXT): ../src/test/quant_solve.cpp ast/ast.h.node smt/params/smt_params.h.node qe/qe.h.node ast/arith_decl_plugin.h.node ast/ast_pp.h.node util/lbool.h.node ast/rewriter/expr_replacer.h.node smt/smt_kernel.h.node ast/reg_decl_plugins.h.node ast/expr_abstract.h.node model/model_smt2_pp.h.node parsers/smt2/smt2parser.h.node ast/rewriter/var_subst.h.node
	@echo src/test/quant_solve.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/quant_solve$(OBJ_EXT) ../src/test/quant_solve.cpp
test/polynomial$(OBJ_EXT): ../src/test/polynomial.cpp math/polynomial/polynomial.h.node math/polynomial/polynomial_var2value.h.node math/polynomial/polynomial_cache.h.node math/polynomial/linear_eq_solver.h.node util/rlimit.h.node
	@echo src/test/polynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/polynomial$(OBJ_EXT) ../src/test/polynomial.cpp
test/fixed_bit_vector$(OBJ_EXT): ../src/test/fixed_bit_vector.cpp util/fixed_bit_vector.h.node util/vector.h.node
	@echo src/test/fixed_bit_vector.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/fixed_bit_vector$(OBJ_EXT) ../src/test/fixed_bit_vector.cpp
test/gparams_register_modules$(OBJ_EXT): ../src/test/gparams_register_modules.cpp util/gparams.h.node ackermannization/ackermannization_params.hpp.node ackermannization/ackermannize_bv_tactic_params.hpp.node ast/normal_forms/nnf.h.node ast/normal_forms/nnf_params.hpp.node ast/pp_params.hpp.node math/polynomial/algebraic_params.hpp.node math/realclosure/rcf_params.hpp.node model/model_evaluator_params.hpp.node model/model_params.hpp.node muz/base/fp_params.hpp.node nlsat/nlsat_params.hpp.node opt/opt_params.hpp.node params/arith_rewriter_params.hpp.node params/array_rewriter_params.hpp.node params/bool_rewriter_params.hpp.node params/bv_rewriter_params.hpp.node params/context_params.h.node params/fpa2bv_rewriter_params.hpp.node params/fpa_rewriter_params.hpp.node params/pattern_inference_params_helper.hpp.node params/poly_rewriter_params.hpp.node params/rewriter_params.hpp.node params/seq_rewriter_params.hpp.node parsers/util/parser_params.hpp.node sat/sat_asymm_branch_params.hpp.node sat/sat_params.hpp.node sat/sat_scc_params.hpp.node sat/sat_simplifier_params.hpp.node smt/params/smt_params_helper.hpp.node solver/combined_solver_params.hpp.node solver/parallel_params.hpp.node solver/solver_params.hpp.node tactic/sls/sls_params.hpp.node tactic/smtlogics/qfufbv_tactic_params.hpp.node tactic/tactic_params.hpp.node util/env_params.h.node
	@echo src/test/gparams_register_modules.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/gparams_register_modules$(OBJ_EXT) ../src/test/gparams_register_modules.cpp
test/inf_rational$(OBJ_EXT): ../src/test/inf_rational.cpp util/inf_rational.h.node
	@echo src/test/inf_rational.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/inf_rational$(OBJ_EXT) ../src/test/inf_rational.cpp
test/check_assumptions$(OBJ_EXT): ../src/test/check_assumptions.cpp util/memory_manager.h.node smt/params/smt_params.h.node ast/ast.h.node ast/arith_decl_plugin.h.node ast/bv_decl_plugin.h.node smt/smt_context.h.node ast/reg_decl_plugins.h.node
	@echo src/test/check_assumptions.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/check_assumptions$(OBJ_EXT) ../src/test/check_assumptions.cpp
test/symbol_table$(OBJ_EXT): ../src/test/symbol_table.cpp util/symbol_table.h.node
	@echo src/test/symbol_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/symbol_table$(OBJ_EXT) ../src/test/symbol_table.cpp
test/expr_substitution$(OBJ_EXT): ../src/test/expr_substitution.cpp ast/expr_substitution.h.node smt/params/smt_params.h.node ast/substitution/substitution.h.node ast/substitution/unifier.h.node ast/bv_decl_plugin.h.node ast/ast_pp.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node ast/rewriter/th_rewriter.h.node
	@echo src/test/expr_substitution.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/expr_substitution$(OBJ_EXT) ../src/test/expr_substitution.cpp
util/object_allocator.h.node: ../src/util/object_allocator.h util/util.h.node util/vector.h.node
	@echo done > util/object_allocator.h.node
test/object_allocator$(OBJ_EXT): ../src/test/object_allocator.cpp util/rational.h.node util/object_allocator.h.node
	@echo src/test/object_allocator.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/object_allocator$(OBJ_EXT) ../src/test/object_allocator.cpp
test/get_consequences$(OBJ_EXT): ../src/test/get_consequences.cpp sat/sat_solver/inc_sat_solver.h.node ast/bv_decl_plugin.h.node ast/datatype_decl_plugin.h.node ast/reg_decl_plugins.h.node ast/ast_pp.h.node tactic/bv/dt2bv_tactic.h.node tactic/tactic.h.node model/model_smt2_pp.h.node tactic/fd_solver/fd_solver.h.node
	@echo src/test/get_consequences.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/get_consequences$(OBJ_EXT) ../src/test/get_consequences.cpp
test/mem_initializer$(OBJ_EXT): ../src/test/mem_initializer.cpp util/debug.h.node util/gparams.h.node util/prime_generator.h.node util/rational.h.node util/rlimit.h.node util/scoped_timer.h.node util/symbol.h.node util/trace.h.node
	@echo src/test/mem_initializer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/mem_initializer$(OBJ_EXT) ../src/test/mem_initializer.cpp
test/theory_pb$(OBJ_EXT): ../src/test/theory_pb.cpp smt/smt_context.h.node ast/ast_pp.h.node model/model_v2_pp.h.node ast/reg_decl_plugins.h.node smt/theory_pb.h.node ast/rewriter/th_rewriter.h.node
	@echo src/test/theory_pb.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/theory_pb$(OBJ_EXT) ../src/test/theory_pb.cpp
test/arith_simplifier_plugin$(OBJ_EXT): ../src/test/arith_simplifier_plugin.cpp smt/arith_eq_solver.h.node smt/params/smt_params.h.node
	@echo src/test/arith_simplifier_plugin.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/arith_simplifier_plugin$(OBJ_EXT) ../src/test/arith_simplifier_plugin.cpp
test/im_float_config.h.node: ../src/test/im_float_config.h util/f2n.h.node util/mpf.h.node util/hwf.h.node
	@echo done > test/im_float_config.h.node
test/trigo$(OBJ_EXT): ../src/test/trigo.cpp math/interval/interval_def.h.node util/dependency.h.node util/mpq.h.node ast/ast.h.node util/debug.h.node test/im_float_config.h.node util/rlimit.h.node
	@echo src/test/trigo.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/trigo$(OBJ_EXT) ../src/test/trigo.cpp
test/smt2print_parse$(OBJ_EXT): ../src/test/smt2print_parse.cpp api/z3.h.node
	@echo src/test/smt2print_parse.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/smt2print_parse$(OBJ_EXT) ../src/test/smt2print_parse.cpp
test/value_sweep$(OBJ_EXT): ../src/test/value_sweep.cpp ast/rewriter/value_sweep.h.node ast/reg_decl_plugins.h.node ast/ast_pp.h.node ast/seq_decl_plugin.h.node ast/array_decl_plugin.h.node
	@echo src/test/value_sweep.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/value_sweep$(OBJ_EXT) ../src/test/value_sweep.cpp
test/cnf_backbones$(OBJ_EXT): ../src/test/cnf_backbones.cpp util/timeout.h.node util/rlimit.h.node sat/dimacs.h.node sat/sat_solver.h.node util/gparams.h.node
	@echo src/test/cnf_backbones.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/cnf_backbones$(OBJ_EXT) ../src/test/cnf_backbones.cpp
test/interval$(OBJ_EXT): ../src/test/interval.cpp math/interval/interval_def.h.node util/dependency.h.node util/mpq.h.node ast/ast.h.node util/debug.h.node util/rlimit.h.node test/im_float_config.h.node
	@echo src/test/interval.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/interval$(OBJ_EXT) ../src/test/interval.cpp
test/simplifier$(OBJ_EXT): ../src/test/simplifier.cpp api/z3.h.node api/z3_private.h.node util/util.h.node util/trace.h.node
	@echo src/test/simplifier.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/simplifier$(OBJ_EXT) ../src/test/simplifier.cpp
test/sat_local_search$(OBJ_EXT): ../src/test/sat_local_search.cpp sat/sat_local_search.h.node sat/sat_solver.h.node util/cancel_eh.h.node util/scoped_ctrl_c.h.node util/scoped_timer.h.node
	@echo src/test/sat_local_search.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/sat_local_search$(OBJ_EXT) ../src/test/sat_local_search.cpp
test/theory_dl$(OBJ_EXT): ../src/test/theory_dl.cpp ast/dl_decl_plugin.h.node ast/reg_decl_plugins.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node smt/smt_context.h.node model/model_v2_pp.h.node
	@echo src/test/theory_dl.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/theory_dl$(OBJ_EXT) ../src/test/theory_dl.cpp
test/main$(OBJ_EXT): ../src/test/main.cpp util/util.h.node util/trace.h.node util/debug.h.node util/timeit.h.node util/warning.h.node util/memory_manager.h.node util/gparams.h.node
	@echo src/test/main.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/main$(OBJ_EXT) ../src/test/main.cpp
test/bits$(OBJ_EXT): ../src/test/bits.cpp util/util.h.node util/debug.h.node util/vector.h.node util/mpz.h.node util/bit_util.h.node
	@echo src/test/bits.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/bits$(OBJ_EXT) ../src/test/bits.cpp
test/tbv$(OBJ_EXT): ../src/test/tbv.cpp muz/rel/tbv.h.node
	@echo src/test/tbv.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/tbv$(OBJ_EXT) ../src/test/tbv.cpp
test/dl_relation$(OBJ_EXT): ../src/test/dl_relation.cpp ast/reg_decl_plugins.h.node muz/base/dl_context.h.node muz/fp/dl_register_engine.h.node muz/rel/dl_relation_manager.h.node muz/rel/dl_interval_relation.h.node muz/rel/dl_bound_relation.h.node muz/rel/dl_product_relation.h.node util/util.h.node
	@echo src/test/dl_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/dl_relation$(OBJ_EXT) ../src/test/dl_relation.cpp
test/sat_user_scope$(OBJ_EXT): ../src/test/sat_user_scope.cpp sat/sat_solver.h.node util/util.h.node
	@echo src/test/sat_user_scope.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/sat_user_scope$(OBJ_EXT) ../src/test/sat_user_scope.cpp
test/doc$(OBJ_EXT): ../src/test/doc.cpp muz/rel/doc.h.node util/trace.h.node util/vector.h.node ast/ast.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node util/sorting_network.h.node smt/smt_kernel.h.node model/model_smt2_pp.h.node smt/params/smt_params.h.node ast/ast_util.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/th_rewriter.h.node
	@echo src/test/doc.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/doc$(OBJ_EXT) ../src/test/doc.cpp
test/ddnf$(OBJ_EXT): ../src/test/ddnf.cpp muz/ddnf/ddnf.h.node muz/rel/tbv.h.node
	@echo src/test/ddnf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/ddnf$(OBJ_EXT) ../src/test/ddnf.cpp
test/bdd$(OBJ_EXT): ../src/test/bdd.cpp math/dd/dd_bdd.h.node
	@echo src/test/bdd.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/bdd$(OBJ_EXT) ../src/test/bdd.cpp
test/dl_table$(OBJ_EXT): ../src/test/dl_table.cpp ast/reg_decl_plugins.h.node muz/base/dl_context.h.node muz/rel/dl_table.h.node muz/fp/dl_register_engine.h.node muz/rel/dl_relation_manager.h.node
	@echo src/test/dl_table.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/dl_table$(OBJ_EXT) ../src/test/dl_table.cpp
test/cube_clause$(OBJ_EXT): ../src/test/cube_clause.cpp ast/reg_decl_plugins.h.node solver/solver_pool.h.node smt/smt_solver.h.node
	@echo src/test/cube_clause.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/cube_clause$(OBJ_EXT) ../src/test/cube_clause.cpp
test/solver_pool$(OBJ_EXT): ../src/test/solver_pool.cpp ast/reg_decl_plugins.h.node solver/solver_pool.h.node smt/smt_solver.h.node
	@echo src/test/solver_pool.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/solver_pool$(OBJ_EXT) ../src/test/solver_pool.cpp
test/random$(OBJ_EXT): ../src/test/random.cpp util/util.h.node util/trace.h.node
	@echo src/test/random.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/random$(OBJ_EXT) ../src/test/random.cpp
test/permutation$(OBJ_EXT): ../src/test/permutation.cpp util/permutation.h.node util/util.h.node util/vector.h.node
	@echo src/test/permutation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/permutation$(OBJ_EXT) ../src/test/permutation.cpp
test/escaped$(OBJ_EXT): ../src/test/escaped.cpp util/util.h.node
	@echo src/test/escaped.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/escaped$(OBJ_EXT) ../src/test/escaped.cpp
test/hwf$(OBJ_EXT): ../src/test/hwf.cpp util/hwf.h.node util/f2n.h.node util/rational.h.node
	@echo src/test/hwf.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/hwf$(OBJ_EXT) ../src/test/hwf.cpp
test/algebraic$(OBJ_EXT): ../src/test/algebraic.cpp math/polynomial/algebraic_numbers.h.node math/polynomial/polynomial_var2value.h.node util/mpbq.h.node util/rlimit.h.node
	@echo src/test/algebraic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/algebraic$(OBJ_EXT) ../src/test/algebraic.cpp
test/pb2bv$(OBJ_EXT): ../src/test/pb2bv.cpp util/trace.h.node util/vector.h.node ast/ast.h.node ast/ast_pp.h.node util/statistics.h.node ast/reg_decl_plugins.h.node ast/rewriter/pb2bv_rewriter.h.node smt/smt_kernel.h.node model/model_smt2_pp.h.node smt/params/smt_params.h.node ast/ast_util.h.node ast/pb_decl_plugin.h.node ast/rewriter/th_rewriter.h.node tactic/fd_solver/fd_solver.h.node solver/solver.h.node ast/arith_decl_plugin.h.node
	@echo src/test/pb2bv.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/pb2bv$(OBJ_EXT) ../src/test/pb2bv.cpp
test/udoc_relation$(OBJ_EXT): ../src/test/udoc_relation.cpp muz/rel/udoc_relation.h.node util/trace.h.node util/vector.h.node ast/ast.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node util/sorting_network.h.node smt/smt_kernel.h.node model/model_smt2_pp.h.node smt/params/smt_params.h.node ast/ast_util.h.node ast/rewriter/expr_safe_replace.h.node ast/rewriter/th_rewriter.h.node muz/rel/dl_relation_manager.h.node muz/fp/dl_register_engine.h.node muz/rel/rel_context.h.node ast/bv_decl_plugin.h.node muz/rel/check_relation.h.node
	@echo src/test/udoc_relation.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/udoc_relation$(OBJ_EXT) ../src/test/udoc_relation.cpp
test/rational$(OBJ_EXT): ../src/test/rational.cpp util/vector.h.node util/rational.h.node util/trace.h.node util/ext_gcd.h.node util/timeit.h.node
	@echo src/test/rational.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/rational$(OBJ_EXT) ../src/test/rational.cpp
test/expr_rand$(OBJ_EXT): ../src/test/expr_rand.cpp test/fuzzing/expr_rand.h.node ast/ast_pp.h.node ast/bv_decl_plugin.h.node ast/array_decl_plugin.h.node ast/arith_decl_plugin.h.node ast/ast_smt_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/expr_rand.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/expr_rand$(OBJ_EXT) ../src/test/expr_rand.cpp
test/region$(OBJ_EXT): ../src/test/region.cpp util/region.h.node
	@echo src/test/region.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/region$(OBJ_EXT) ../src/test/region.cpp
test/factor_rewriter$(OBJ_EXT): ../src/test/factor_rewriter.cpp ast/rewriter/factor_rewriter.h.node ast/bv_decl_plugin.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/factor_rewriter.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/factor_rewriter$(OBJ_EXT) ../src/test/factor_rewriter.cpp
test/sat_lookahead$(OBJ_EXT): ../src/test/sat_lookahead.cpp sat/sat_solver.h.node sat/sat_watched.h.node util/statistics.h.node sat/sat_lookahead.h.node sat/dimacs.h.node
	@echo src/test/sat_lookahead.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/sat_lookahead$(OBJ_EXT) ../src/test/sat_lookahead.cpp
test/diff_logic$(OBJ_EXT): ../src/test/diff_logic.cpp util/rational.h.node smt/diff_logic.h.node smt/smt_literal.h.node util/util.h.node util/debug.h.node
	@echo src/test/diff_logic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/diff_logic$(OBJ_EXT) ../src/test/diff_logic.cpp
test/chashtable$(OBJ_EXT): ../src/test/chashtable.cpp util/chashtable.h.node util/hashtable.h.node util/hash.h.node util/util.h.node
	@echo src/test/chashtable.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/chashtable$(OBJ_EXT) ../src/test/chashtable.cpp
util/uint_map.h.node: ../src/util/uint_map.h util/vector.h.node
	@echo done > util/uint_map.h.node
util/total_order.h.node: ../src/util/total_order.h util/util.h.node util/small_object_allocator.h.node util/map.h.node util/uint_map.h.node util/trace.h.node
	@echo done > util/total_order.h.node
test/total_order$(OBJ_EXT): ../src/test/total_order.cpp util/total_order.h.node util/timeit.h.node
	@echo src/test/total_order.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/total_order$(OBJ_EXT) ../src/test/total_order.cpp
test/old_interval$(OBJ_EXT): ../src/test/old_interval.cpp smt/old_interval.h.node util/basic_interval.h.node util/mpz.h.node util/scoped_numeral.h.node
	@echo src/test/old_interval.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/old_interval$(OBJ_EXT) ../src/test/old_interval.cpp
test/karr$(OBJ_EXT): ../src/test/karr.cpp util/rlimit.h.node math/hilbert/hilbert_basis.h.node
	@echo src/test/karr.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/karr$(OBJ_EXT) ../src/test/karr.cpp
test/polynorm$(OBJ_EXT): ../src/test/polynorm.cpp ast/rewriter/th_rewriter.h.node parsers/smt2/smt2parser.h.node ast/arith_decl_plugin.h.node ast/reg_decl_plugins.h.node ast/rewriter/arith_rewriter.h.node ast/ast_pp.h.node
	@echo src/test/polynorm.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/polynorm$(OBJ_EXT) ../src/test/polynorm.cpp
test/upolynomial$(OBJ_EXT): ../src/test/upolynomial.cpp math/polynomial/upolynomial.h.node util/timeit.h.node util/rlimit.h.node
	@echo src/test/upolynomial.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/upolynomial$(OBJ_EXT) ../src/test/upolynomial.cpp
test/hilbert_basis$(OBJ_EXT): ../src/test/hilbert_basis.cpp math/hilbert/hilbert_basis.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node ast/arith_decl_plugin.h.node tactic/smtlogics/quant_tactics.h.node tactic/tactic.h.node solver/tactic2solver.h.node solver/solver.h.node util/rlimit.h.node
	@echo src/test/hilbert_basis.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/hilbert_basis$(OBJ_EXT) ../src/test/hilbert_basis.cpp
test/optional$(OBJ_EXT): ../src/test/optional.cpp util/trace.h.node util/debug.h.node util/memory_manager.h.node util/optional.h.node
	@echo src/test/optional.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/optional$(OBJ_EXT) ../src/test/optional.cpp
test/matcher$(OBJ_EXT): ../src/test/matcher.cpp ast/substitution/matcher.h.node ast/ast_pp.h.node ast/reg_decl_plugins.h.node
	@echo src/test/matcher.cpp
	@$(CXX) $(CXXFLAGS) $(includes_67) $(CXX_OUT_FLAG)test/matcher$(OBJ_EXT) ../src/test/matcher.cpp
test-z3$(EXE_EXT): test/egraph$(OBJ_EXT) test/datalog_parser$(OBJ_EXT) test/string_buffer$(OBJ_EXT) test/heap_trie$(OBJ_EXT) test/proof_checker$(OBJ_EXT) test/mpbq$(OBJ_EXT) test/bit_vector$(OBJ_EXT) test/ext_numeral$(OBJ_EXT) test/quant_elim$(OBJ_EXT) test/map$(OBJ_EXT) test/memory$(OBJ_EXT) test/dl_product_relation$(OBJ_EXT) test/mpq$(OBJ_EXT) test/model_retrieval$(OBJ_EXT) test/finder$(OBJ_EXT) test/dl_util$(OBJ_EXT) test/stack$(OBJ_EXT) test/install_tactic$(OBJ_EXT) test/simple_parser$(OBJ_EXT) test/mpz$(OBJ_EXT) test/pdd_solver$(OBJ_EXT) test/rcf$(OBJ_EXT) test/model_evaluator$(OBJ_EXT) test/list$(OBJ_EXT) test/sorting_network$(OBJ_EXT) test/vector$(OBJ_EXT) test/api$(OBJ_EXT) test/qe_arith$(OBJ_EXT) test/pdd$(OBJ_EXT) test/parray$(OBJ_EXT) test/uint_set$(OBJ_EXT) test/dl_query$(OBJ_EXT) test/heap$(OBJ_EXT) test/horn_subsume_model_converter$(OBJ_EXT) test/value_generator$(OBJ_EXT) test/mpf$(OBJ_EXT) test/prime_generator$(OBJ_EXT) test/timeout$(OBJ_EXT) test/model2expr$(OBJ_EXT) test/small_object_allocator$(OBJ_EXT) test/no_overflow$(OBJ_EXT) test/get_implied_equalities$(OBJ_EXT) test/f2n$(OBJ_EXT) test/bit_blaster$(OBJ_EXT) test/substitution$(OBJ_EXT) test/mpfx$(OBJ_EXT) test/hashtable$(OBJ_EXT) test/mpff$(OBJ_EXT) test/nlarith_util$(OBJ_EXT) test/nlsat$(OBJ_EXT) test/simplex$(OBJ_EXT) test/var_subst$(OBJ_EXT) test/arith_rewriter$(OBJ_EXT) test/api_bug$(OBJ_EXT) test/for_each_file$(OBJ_EXT) test/symbol$(OBJ_EXT) test/ast$(OBJ_EXT) test/ex$(OBJ_EXT) test/smt_context$(OBJ_EXT) test/model_based_opt$(OBJ_EXT) test/buffer$(OBJ_EXT) test/dl_context$(OBJ_EXT) test/quant_solve$(OBJ_EXT) test/polynomial$(OBJ_EXT) test/fixed_bit_vector$(OBJ_EXT) test/gparams_register_modules$(OBJ_EXT) test/inf_rational$(OBJ_EXT) test/check_assumptions$(OBJ_EXT) test/symbol_table$(OBJ_EXT) test/expr_substitution$(OBJ_EXT) test/object_allocator$(OBJ_EXT) test/get_consequences$(OBJ_EXT) test/mem_initializer$(OBJ_EXT) test/theory_pb$(OBJ_EXT) test/arith_simplifier_plugin$(OBJ_EXT) test/trigo$(OBJ_EXT) test/smt2print_parse$(OBJ_EXT) test/value_sweep$(OBJ_EXT) test/cnf_backbones$(OBJ_EXT) test/interval$(OBJ_EXT) test/simplifier$(OBJ_EXT) test/sat_local_search$(OBJ_EXT) test/theory_dl$(OBJ_EXT) test/main$(OBJ_EXT) test/bits$(OBJ_EXT) test/tbv$(OBJ_EXT) test/dl_relation$(OBJ_EXT) test/sat_user_scope$(OBJ_EXT) test/doc$(OBJ_EXT) test/ddnf$(OBJ_EXT) test/bdd$(OBJ_EXT) test/dl_table$(OBJ_EXT) test/cube_clause$(OBJ_EXT) test/solver_pool$(OBJ_EXT) test/random$(OBJ_EXT) test/permutation$(OBJ_EXT) test/escaped$(OBJ_EXT) test/hwf$(OBJ_EXT) test/algebraic$(OBJ_EXT) test/pb2bv$(OBJ_EXT) test/udoc_relation$(OBJ_EXT) test/rational$(OBJ_EXT) test/expr_rand$(OBJ_EXT) test/region$(OBJ_EXT) test/factor_rewriter$(OBJ_EXT) test/sat_lookahead$(OBJ_EXT) test/diff_logic$(OBJ_EXT) test/chashtable$(OBJ_EXT) test/total_order$(OBJ_EXT) test/old_interval$(OBJ_EXT) test/karr$(OBJ_EXT) test/polynorm$(OBJ_EXT) test/upolynomial$(OBJ_EXT) test/hilbert_basis$(OBJ_EXT) test/optional$(OBJ_EXT) test/matcher$(OBJ_EXT) api/api$(LIB_EXT) opt/opt$(LIB_EXT) tactic/portfolio/portfolio$(LIB_EXT) tactic/fpa/fpa_tactics$(LIB_EXT) tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/ufbv/ufbv_tactic$(LIB_EXT) muz/fp/fp$(LIB_EXT) muz/bmc/bmc$(LIB_EXT) muz/ddnf/ddnf$(LIB_EXT) muz/tab/tab$(LIB_EXT) muz/clp/clp$(LIB_EXT) muz/spacer/spacer$(LIB_EXT) muz/rel/rel$(LIB_EXT) muz/transforms/transforms$(LIB_EXT) muz/dataflow/dataflow$(LIB_EXT) muz/base/muz$(LIB_EXT) tactic/fd_solver/fd_solver$(LIB_EXT) sat/sat_solver/sat_solver$(LIB_EXT) qe/qe$(LIB_EXT) tactic/sls/sls_tactic$(LIB_EXT) smt/tactic/smt_tactic$(LIB_EXT) test/fuzzing/fuzzing$(LIB_EXT) tactic/bv/bv_tactics$(LIB_EXT) smt/smt$(LIB_EXT) smt/proto_model/proto_model$(LIB_EXT) math/subpaving/tactic/subpaving_tactic$(LIB_EXT) nlsat/tactic/nlsat_tactic$(LIB_EXT) sat/tactic/sat_tactic$(LIB_EXT) sat/smt/sat_smt$(LIB_EXT) qe/mbp/mbp$(LIB_EXT) tactic/arith/arith_tactics$(LIB_EXT) tactic/core/core_tactics$(LIB_EXT) smt/params/smt_params$(LIB_EXT) ast/pattern/pattern$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/fpa/fpa$(LIB_EXT) ackermannization/ackermannization$(LIB_EXT) tactic/aig/aig_tactic$(LIB_EXT) parsers/smt2/smt2parser$(LIB_EXT) cmd_context/cmd_context$(LIB_EXT) solver/solver$(LIB_EXT) ast/proofs/proofs$(LIB_EXT) parsers/util/parser_util$(LIB_EXT) ast/substitution/substitution$(LIB_EXT) tactic/tactic$(LIB_EXT) model/model$(LIB_EXT) ast/normal_forms/normal_forms$(LIB_EXT) ast/macros/macros$(LIB_EXT) ast/rewriter/rewriter$(LIB_EXT) math/lp/lp$(LIB_EXT) nlsat/nlsat$(LIB_EXT) sat/sat$(LIB_EXT) math/grobner/grobner$(LIB_EXT) ast/euf/euf$(LIB_EXT) params/params$(LIB_EXT) ast/ast$(LIB_EXT) math/subpaving/subpaving$(LIB_EXT) math/realclosure/realclosure$(LIB_EXT) math/automata/automata$(LIB_EXT) math/hilbert/hilbert$(LIB_EXT) math/simplex/simplex$(LIB_EXT) math/dd/dd$(LIB_EXT) math/interval/interval$(LIB_EXT) math/polynomial/polynomial$(LIB_EXT) util/util$(LIB_EXT)
	$(LINK)  $(LINK_OUT_FLAG)test-z3$(EXE_EXT) $(LINK_FLAGS) test/egraph$(OBJ_EXT) test/datalog_parser$(OBJ_EXT) test/string_buffer$(OBJ_EXT) test/heap_trie$(OBJ_EXT) test/proof_checker$(OBJ_EXT) test/mpbq$(OBJ_EXT) test/bit_vector$(OBJ_EXT) test/ext_numeral$(OBJ_EXT) test/quant_elim$(OBJ_EXT) test/map$(OBJ_EXT) test/memory$(OBJ_EXT) test/dl_product_relation$(OBJ_EXT) test/mpq$(OBJ_EXT) test/model_retrieval$(OBJ_EXT) test/finder$(OBJ_EXT) test/dl_util$(OBJ_EXT) test/stack$(OBJ_EXT) test/install_tactic$(OBJ_EXT) test/simple_parser$(OBJ_EXT) test/mpz$(OBJ_EXT) test/pdd_solver$(OBJ_EXT) test/rcf$(OBJ_EXT) test/model_evaluator$(OBJ_EXT) test/list$(OBJ_EXT) test/sorting_network$(OBJ_EXT) test/vector$(OBJ_EXT) test/api$(OBJ_EXT) test/qe_arith$(OBJ_EXT) test/pdd$(OBJ_EXT) test/parray$(OBJ_EXT) test/uint_set$(OBJ_EXT) test/dl_query$(OBJ_EXT) test/heap$(OBJ_EXT) test/horn_subsume_model_converter$(OBJ_EXT) test/value_generator$(OBJ_EXT) test/mpf$(OBJ_EXT) test/prime_generator$(OBJ_EXT) test/timeout$(OBJ_EXT) test/model2expr$(OBJ_EXT) test/small_object_allocator$(OBJ_EXT) test/no_overflow$(OBJ_EXT) test/get_implied_equalities$(OBJ_EXT) test/f2n$(OBJ_EXT) test/bit_blaster$(OBJ_EXT) test/substitution$(OBJ_EXT) test/mpfx$(OBJ_EXT) test/hashtable$(OBJ_EXT) test/mpff$(OBJ_EXT) test/nlarith_util$(OBJ_EXT) test/nlsat$(OBJ_EXT) test/simplex$(OBJ_EXT) test/var_subst$(OBJ_EXT) test/arith_rewriter$(OBJ_EXT) test/api_bug$(OBJ_EXT) test/for_each_file$(OBJ_EXT) test/symbol$(OBJ_EXT) test/ast$(OBJ_EXT) test/ex$(OBJ_EXT) test/smt_context$(OBJ_EXT) test/model_based_opt$(OBJ_EXT) test/buffer$(OBJ_EXT) test/dl_context$(OBJ_EXT) test/quant_solve$(OBJ_EXT) test/polynomial$(OBJ_EXT) test/fixed_bit_vector$(OBJ_EXT) test/gparams_register_modules$(OBJ_EXT) test/inf_rational$(OBJ_EXT) test/check_assumptions$(OBJ_EXT) test/symbol_table$(OBJ_EXT) test/expr_substitution$(OBJ_EXT) test/object_allocator$(OBJ_EXT) test/get_consequences$(OBJ_EXT) test/mem_initializer$(OBJ_EXT) test/theory_pb$(OBJ_EXT) test/arith_simplifier_plugin$(OBJ_EXT) test/trigo$(OBJ_EXT) test/smt2print_parse$(OBJ_EXT) test/value_sweep$(OBJ_EXT) test/cnf_backbones$(OBJ_EXT) test/interval$(OBJ_EXT) test/simplifier$(OBJ_EXT) test/sat_local_search$(OBJ_EXT) test/theory_dl$(OBJ_EXT) test/main$(OBJ_EXT) test/bits$(OBJ_EXT) test/tbv$(OBJ_EXT) test/dl_relation$(OBJ_EXT) test/sat_user_scope$(OBJ_EXT) test/doc$(OBJ_EXT) test/ddnf$(OBJ_EXT) test/bdd$(OBJ_EXT) test/dl_table$(OBJ_EXT) test/cube_clause$(OBJ_EXT) test/solver_pool$(OBJ_EXT) test/random$(OBJ_EXT) test/permutation$(OBJ_EXT) test/escaped$(OBJ_EXT) test/hwf$(OBJ_EXT) test/algebraic$(OBJ_EXT) test/pb2bv$(OBJ_EXT) test/udoc_relation$(OBJ_EXT) test/rational$(OBJ_EXT) test/expr_rand$(OBJ_EXT) test/region$(OBJ_EXT) test/factor_rewriter$(OBJ_EXT) test/sat_lookahead$(OBJ_EXT) test/diff_logic$(OBJ_EXT) test/chashtable$(OBJ_EXT) test/total_order$(OBJ_EXT) test/old_interval$(OBJ_EXT) test/karr$(OBJ_EXT) test/polynorm$(OBJ_EXT) test/upolynomial$(OBJ_EXT) test/hilbert_basis$(OBJ_EXT) test/optional$(OBJ_EXT) test/matcher$(OBJ_EXT) api/api$(LIB_EXT) opt/opt$(LIB_EXT) tactic/portfolio/portfolio$(LIB_EXT) tactic/fpa/fpa_tactics$(LIB_EXT) tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/ufbv/ufbv_tactic$(LIB_EXT) muz/fp/fp$(LIB_EXT) muz/bmc/bmc$(LIB_EXT) muz/ddnf/ddnf$(LIB_EXT) muz/tab/tab$(LIB_EXT) muz/clp/clp$(LIB_EXT) muz/spacer/spacer$(LIB_EXT) muz/rel/rel$(LIB_EXT) muz/transforms/transforms$(LIB_EXT) muz/dataflow/dataflow$(LIB_EXT) muz/base/muz$(LIB_EXT) tactic/fd_solver/fd_solver$(LIB_EXT) sat/sat_solver/sat_solver$(LIB_EXT) qe/qe$(LIB_EXT) tactic/sls/sls_tactic$(LIB_EXT) smt/tactic/smt_tactic$(LIB_EXT) test/fuzzing/fuzzing$(LIB_EXT) tactic/bv/bv_tactics$(LIB_EXT) smt/smt$(LIB_EXT) smt/proto_model/proto_model$(LIB_EXT) math/subpaving/tactic/subpaving_tactic$(LIB_EXT) nlsat/tactic/nlsat_tactic$(LIB_EXT) sat/tactic/sat_tactic$(LIB_EXT) sat/smt/sat_smt$(LIB_EXT) qe/mbp/mbp$(LIB_EXT) tactic/arith/arith_tactics$(LIB_EXT) tactic/core/core_tactics$(LIB_EXT) smt/params/smt_params$(LIB_EXT) ast/pattern/pattern$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/fpa/fpa$(LIB_EXT) ackermannization/ackermannization$(LIB_EXT) tactic/aig/aig_tactic$(LIB_EXT) parsers/smt2/smt2parser$(LIB_EXT) cmd_context/cmd_context$(LIB_EXT) solver/solver$(LIB_EXT) ast/proofs/proofs$(LIB_EXT) parsers/util/parser_util$(LIB_EXT) ast/substitution/substitution$(LIB_EXT) tactic/tactic$(LIB_EXT) model/model$(LIB_EXT) ast/normal_forms/normal_forms$(LIB_EXT) ast/macros/macros$(LIB_EXT) ast/rewriter/rewriter$(LIB_EXT) math/lp/lp$(LIB_EXT) nlsat/nlsat$(LIB_EXT) sat/sat$(LIB_EXT) math/grobner/grobner$(LIB_EXT) ast/euf/euf$(LIB_EXT) params/params$(LIB_EXT) ast/ast$(LIB_EXT) math/subpaving/subpaving$(LIB_EXT) math/realclosure/realclosure$(LIB_EXT) math/automata/automata$(LIB_EXT) math/hilbert/hilbert$(LIB_EXT) math/simplex/simplex$(LIB_EXT) math/dd/dd$(LIB_EXT) math/interval/interval$(LIB_EXT) math/polynomial/polynomial$(LIB_EXT) util/util$(LIB_EXT) $(LINK_EXTRA_FLAGS)
test: test-z3$(EXE_EXT)

includes_68 = -I../src/api -I../src/tactic/portfolio -I../src/tactic/smtlogics -I../src/ackermannization -I../src/model -I../src/ast/rewriter -I../src/ast -I../src/util -I../src/math/polynomial -I../src/math/automata -I../src/params -I../src/ast/macros -I../src/solver -I../src/tactic -I../src/ast/proofs -I../src/sat/sat_solver -I../src/tactic/core -I../src/ast/normal_forms -I../src/ast/pattern -I../src/parsers/smt2 -I../src/cmd_context -I../src/parsers/util -I../src/tactic/aig -I../src/tactic/bv -I../src/ast/rewriter/bit_blaster -I../src/tactic/arith -I../src/sat -I../src/math/dd -I../src/math/interval -I../src/math/grobner -I../src/math/simplex -I../src/sat/tactic -I../src/sat/smt -I../src/ast/euf -I../src/smt/params -I../src/ast/fpa -I../src/qe/mbp -I../src/math/lp -I../src/nlsat -I../src/nlsat/tactic -I../src/smt/tactic -I../src/smt -I../src/smt/proto_model -I../src/ast/substitution -I../src/muz/fp -I../src/muz/base -I../src/qe -I../src/muz/clp -I../src/muz/transforms -I../src/math/hilbert -I../src/muz/dataflow -I../src/muz/tab -I../src/muz/rel -I../src/muz/bmc -I../src/tactic/fd_solver -I../src/muz/ddnf -I../src/muz/spacer -I../src/tactic/ufbv -I../src/tactic/fpa -I../src/tactic/sls -I../src/math/subpaving/tactic -I../src/math/subpaving -I../src/math/realclosure -I../src/opt -I../src/cmd_context/extra_cmds -I../src
api/dll/install_tactic$(OBJ_EXT): ../src/api/dll/install_tactic.cpp tactic/tactic.h.node cmd_context/tactic_cmds.h.node cmd_context/cmd_context.h.node ackermannization/ackermannize_bv_tactic.h.node ackermannization/ackr_bound_probe.h.node math/subpaving/tactic/subpaving_tactic.h.node muz/fp/horn_tactic.h.node nlsat/tactic/nlsat_tactic.h.node nlsat/tactic/qfnra_nlsat_tactic.h.node qe/nlqsat.h.node qe/qe_lite.h.node qe/qe_tactic.h.node qe/qsat.h.node sat/sat_solver/inc_sat_solver.h.node sat/tactic/sat_tactic.h.node smt/tactic/ctx_solver_simplify_tactic.h.node smt/tactic/smt_tactic.h.node smt/tactic/unit_subsumption_tactic.h.node tactic/aig/aig_tactic.h.node tactic/arith/add_bounds_tactic.h.node tactic/arith/card2bv_tactic.h.node tactic/arith/degree_shift_tactic.h.node tactic/arith/diff_neq_tactic.h.node tactic/arith/eq2bv_tactic.h.node tactic/arith/factor_tactic.h.node tactic/arith/fix_dl_var_tactic.h.node tactic/arith/fm_tactic.h.node tactic/arith/lia2card_tactic.h.node tactic/arith/lia2pb_tactic.h.node tactic/arith/nla2bv_tactic.h.node tactic/arith/normalize_bounds_tactic.h.node tactic/arith/pb2bv_tactic.h.node tactic/arith/probe_arith.h.node tactic/arith/propagate_ineqs_tactic.h.node tactic/arith/purify_arith_tactic.h.node tactic/arith/recover_01_tactic.h.node tactic/bv/bit_blaster_tactic.h.node tactic/bv/bv1_blaster_tactic.h.node tactic/bv/bv_bound_chk_tactic.h.node tactic/bv/bv_bounds_tactic.h.node tactic/bv/bv_size_reduction_tactic.h.node tactic/bv/bvarray2uf_tactic.h.node tactic/bv/dt2bv_tactic.h.node tactic/bv/elim_small_bv_tactic.h.node tactic/bv/max_bv_sharing_tactic.h.node tactic/core/blast_term_ite_tactic.h.node tactic/core/cofactor_term_ite_tactic.h.node tactic/core/collect_statistics_tactic.h.node tactic/core/ctx_simplify_tactic.h.node tactic/core/der_tactic.h.node tactic/core/distribute_forall_tactic.h.node tactic/core/dom_simplify_tactic.h.node tactic/core/elim_term_ite_tactic.h.node tactic/core/elim_uncnstr_tactic.h.node tactic/core/injectivity_tactic.h.node tactic/core/nnf_tactic.h.node tactic/core/occf_tactic.h.node tactic/core/pb_preprocess_tactic.h.node tactic/core/propagate_values_tactic.h.node tactic/core/reduce_args_tactic.h.node tactic/core/reduce_invertible_tactic.h.node tactic/core/simplify_tactic.h.node tactic/core/solve_eqs_tactic.h.node tactic/core/special_relations_tactic.h.node tactic/core/split_clause_tactic.h.node tactic/core/symmetry_reduce_tactic.h.node tactic/core/tseitin_cnf_tactic.h.node tactic/fd_solver/fd_solver.h.node tactic/fd_solver/smtfd_solver.h.node tactic/fpa/fpa2bv_tactic.h.node tactic/fpa/qffp_tactic.h.node tactic/fpa/qffplra_tactic.h.node tactic/portfolio/default_tactic.h.node tactic/probe.h.node tactic/sine_filter.h.node tactic/sls/sls_tactic.h.node tactic/smtlogics/nra_tactic.h.node tactic/smtlogics/qfaufbv_tactic.h.node tactic/smtlogics/qfauflia_tactic.h.node tactic/smtlogics/qfbv_tactic.h.node tactic/smtlogics/qfidl_tactic.h.node tactic/smtlogics/qflia_tactic.h.node tactic/smtlogics/qflra_tactic.h.node tactic/smtlogics/qfnia_tactic.h.node tactic/smtlogics/qfnra_tactic.h.node tactic/smtlogics/qfuf_tactic.h.node tactic/smtlogics/qfufbv_tactic.h.node tactic/smtlogics/quant_tactics.h.node tactic/tactic.h.node tactic/ufbv/macro_finder_tactic.h.node tactic/ufbv/quasi_macros_tactic.h.node tactic/ufbv/ufbv_rewriter_tactic.h.node tactic/ufbv/ufbv_tactic.h.node
	@echo src/api/dll/install_tactic.cpp
	@$(CXX) $(CXXFLAGS) $(includes_68) $(CXX_OUT_FLAG)api/dll/install_tactic$(OBJ_EXT) ../src/api/dll/install_tactic.cpp
api/dll/dll$(OBJ_EXT): ../src/api/dll/dll.cpp
	@echo src/api/dll/dll.cpp
	@$(CXX) $(CXXFLAGS) $(includes_68) $(CXX_OUT_FLAG)api/dll/dll$(OBJ_EXT) ../src/api/dll/dll.cpp
api/dll/gparams_register_modules$(OBJ_EXT): ../src/api/dll/gparams_register_modules.cpp util/gparams.h.node ackermannization/ackermannization_params.hpp.node ackermannization/ackermannize_bv_tactic_params.hpp.node ast/normal_forms/nnf.h.node ast/normal_forms/nnf_params.hpp.node ast/pp_params.hpp.node math/polynomial/algebraic_params.hpp.node math/realclosure/rcf_params.hpp.node model/model_evaluator_params.hpp.node model/model_params.hpp.node muz/base/fp_params.hpp.node nlsat/nlsat_params.hpp.node opt/opt_params.hpp.node params/arith_rewriter_params.hpp.node params/array_rewriter_params.hpp.node params/bool_rewriter_params.hpp.node params/bv_rewriter_params.hpp.node params/context_params.h.node params/fpa2bv_rewriter_params.hpp.node params/fpa_rewriter_params.hpp.node params/pattern_inference_params_helper.hpp.node params/poly_rewriter_params.hpp.node params/rewriter_params.hpp.node params/seq_rewriter_params.hpp.node parsers/util/parser_params.hpp.node sat/sat_asymm_branch_params.hpp.node sat/sat_params.hpp.node sat/sat_scc_params.hpp.node sat/sat_simplifier_params.hpp.node smt/params/smt_params_helper.hpp.node solver/combined_solver_params.hpp.node solver/parallel_params.hpp.node solver/solver_params.hpp.node tactic/sls/sls_params.hpp.node tactic/smtlogics/qfufbv_tactic_params.hpp.node tactic/tactic_params.hpp.node util/env_params.h.node
	@echo src/api/dll/gparams_register_modules.cpp
	@$(CXX) $(CXXFLAGS) $(includes_68) $(CXX_OUT_FLAG)api/dll/gparams_register_modules$(OBJ_EXT) ../src/api/dll/gparams_register_modules.cpp
api/dll/mem_initializer$(OBJ_EXT): ../src/api/dll/mem_initializer.cpp util/debug.h.node util/gparams.h.node util/prime_generator.h.node util/rational.h.node util/rlimit.h.node util/scoped_timer.h.node util/symbol.h.node util/trace.h.node
	@echo src/api/dll/mem_initializer.cpp
	@$(CXX) $(CXXFLAGS) $(includes_68) $(CXX_OUT_FLAG)api/dll/mem_initializer$(OBJ_EXT) ../src/api/dll/mem_initializer.cpp
libz3$(SO_EXT): api/dll/install_tactic$(OBJ_EXT) api/dll/dll$(OBJ_EXT) api/dll/gparams_register_modules$(OBJ_EXT) api/dll/mem_initializer$(OBJ_EXT) api/z3_replayer$(OBJ_EXT) api/api_ast_vector$(OBJ_EXT) api/api_log_macros$(OBJ_EXT) api/api_context$(OBJ_EXT) api/api_stats$(OBJ_EXT) api/api_rcf$(OBJ_EXT) api/api_tactic$(OBJ_EXT) api/api_special_relations$(OBJ_EXT) api/api_commands$(OBJ_EXT) api/api_ast$(OBJ_EXT) api/api_pb$(OBJ_EXT) api/api_ast_map$(OBJ_EXT) api/api_config_params$(OBJ_EXT) api/api_log$(OBJ_EXT) api/api_quant$(OBJ_EXT) api/api_numeral$(OBJ_EXT) api/api_datalog$(OBJ_EXT) api/api_array$(OBJ_EXT) api/api_solver$(OBJ_EXT) api/api_qe$(OBJ_EXT) api/api_parsers$(OBJ_EXT) api/api_seq$(OBJ_EXT) api/api_goal$(OBJ_EXT) api/api_opt$(OBJ_EXT) api/api_algebraic$(OBJ_EXT) api/api_arith$(OBJ_EXT) api/api_datatype$(OBJ_EXT) api/api_model$(OBJ_EXT) api/api_params$(OBJ_EXT) api/api_fpa$(OBJ_EXT) api/api_polynomial$(OBJ_EXT) api/api_bv$(OBJ_EXT) cmd_context/extra_cmds/extra_cmds$(LIB_EXT) opt/opt$(LIB_EXT) tactic/portfolio/portfolio$(LIB_EXT) tactic/fpa/fpa_tactics$(LIB_EXT) tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/ufbv/ufbv_tactic$(LIB_EXT) muz/fp/fp$(LIB_EXT) muz/bmc/bmc$(LIB_EXT) muz/ddnf/ddnf$(LIB_EXT) muz/tab/tab$(LIB_EXT) muz/clp/clp$(LIB_EXT) muz/spacer/spacer$(LIB_EXT) muz/rel/rel$(LIB_EXT) muz/transforms/transforms$(LIB_EXT) muz/dataflow/dataflow$(LIB_EXT) muz/base/muz$(LIB_EXT) tactic/fd_solver/fd_solver$(LIB_EXT) sat/sat_solver/sat_solver$(LIB_EXT) qe/qe$(LIB_EXT) tactic/sls/sls_tactic$(LIB_EXT) smt/tactic/smt_tactic$(LIB_EXT) tactic/bv/bv_tactics$(LIB_EXT) smt/smt$(LIB_EXT) smt/proto_model/proto_model$(LIB_EXT) math/subpaving/tactic/subpaving_tactic$(LIB_EXT) nlsat/tactic/nlsat_tactic$(LIB_EXT) sat/tactic/sat_tactic$(LIB_EXT) sat/smt/sat_smt$(LIB_EXT) qe/mbp/mbp$(LIB_EXT) tactic/arith/arith_tactics$(LIB_EXT) tactic/core/core_tactics$(LIB_EXT) smt/params/smt_params$(LIB_EXT) ast/pattern/pattern$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/fpa/fpa$(LIB_EXT) ackermannization/ackermannization$(LIB_EXT) tactic/aig/aig_tactic$(LIB_EXT) parsers/smt2/smt2parser$(LIB_EXT) cmd_context/cmd_context$(LIB_EXT) solver/solver$(LIB_EXT) ast/proofs/proofs$(LIB_EXT) parsers/util/parser_util$(LIB_EXT) ast/substitution/substitution$(LIB_EXT) tactic/tactic$(LIB_EXT) model/model$(LIB_EXT) ast/normal_forms/normal_forms$(LIB_EXT) ast/macros/macros$(LIB_EXT) ast/rewriter/rewriter$(LIB_EXT) math/lp/lp$(LIB_EXT) nlsat/nlsat$(LIB_EXT) sat/sat$(LIB_EXT) math/grobner/grobner$(LIB_EXT) ast/euf/euf$(LIB_EXT) params/params$(LIB_EXT) ast/ast$(LIB_EXT) math/subpaving/subpaving$(LIB_EXT) math/realclosure/realclosure$(LIB_EXT) math/automata/automata$(LIB_EXT) math/hilbert/hilbert$(LIB_EXT) math/simplex/simplex$(LIB_EXT) math/dd/dd$(LIB_EXT) math/interval/interval$(LIB_EXT) math/polynomial/polynomial$(LIB_EXT) util/util$(LIB_EXT)
	$(LINK) $(SLINK_OUT_FLAG)libz3$(SO_EXT) $(SLINK_FLAGS) api/dll/install_tactic$(OBJ_EXT) api/dll/dll$(OBJ_EXT) api/dll/gparams_register_modules$(OBJ_EXT) api/dll/mem_initializer$(OBJ_EXT) api/z3_replayer$(OBJ_EXT) api/api_ast_vector$(OBJ_EXT) api/api_log_macros$(OBJ_EXT) api/api_context$(OBJ_EXT) api/api_stats$(OBJ_EXT) api/api_rcf$(OBJ_EXT) api/api_tactic$(OBJ_EXT) api/api_special_relations$(OBJ_EXT) api/api_commands$(OBJ_EXT) api/api_ast$(OBJ_EXT) api/api_pb$(OBJ_EXT) api/api_ast_map$(OBJ_EXT) api/api_config_params$(OBJ_EXT) api/api_log$(OBJ_EXT) api/api_quant$(OBJ_EXT) api/api_numeral$(OBJ_EXT) api/api_datalog$(OBJ_EXT) api/api_array$(OBJ_EXT) api/api_solver$(OBJ_EXT) api/api_qe$(OBJ_EXT) api/api_parsers$(OBJ_EXT) api/api_seq$(OBJ_EXT) api/api_goal$(OBJ_EXT) api/api_opt$(OBJ_EXT) api/api_algebraic$(OBJ_EXT) api/api_arith$(OBJ_EXT) api/api_datatype$(OBJ_EXT) api/api_model$(OBJ_EXT) api/api_params$(OBJ_EXT) api/api_fpa$(OBJ_EXT) api/api_polynomial$(OBJ_EXT) api/api_bv$(OBJ_EXT) cmd_context/extra_cmds/extra_cmds$(LIB_EXT) opt/opt$(LIB_EXT) tactic/portfolio/portfolio$(LIB_EXT) tactic/fpa/fpa_tactics$(LIB_EXT) tactic/smtlogics/smtlogic_tactics$(LIB_EXT) tactic/ufbv/ufbv_tactic$(LIB_EXT) muz/fp/fp$(LIB_EXT) muz/bmc/bmc$(LIB_EXT) muz/ddnf/ddnf$(LIB_EXT) muz/tab/tab$(LIB_EXT) muz/clp/clp$(LIB_EXT) muz/spacer/spacer$(LIB_EXT) muz/rel/rel$(LIB_EXT) muz/transforms/transforms$(LIB_EXT) muz/dataflow/dataflow$(LIB_EXT) muz/base/muz$(LIB_EXT) tactic/fd_solver/fd_solver$(LIB_EXT) sat/sat_solver/sat_solver$(LIB_EXT) qe/qe$(LIB_EXT) tactic/sls/sls_tactic$(LIB_EXT) smt/tactic/smt_tactic$(LIB_EXT) tactic/bv/bv_tactics$(LIB_EXT) smt/smt$(LIB_EXT) smt/proto_model/proto_model$(LIB_EXT) math/subpaving/tactic/subpaving_tactic$(LIB_EXT) nlsat/tactic/nlsat_tactic$(LIB_EXT) sat/tactic/sat_tactic$(LIB_EXT) sat/smt/sat_smt$(LIB_EXT) qe/mbp/mbp$(LIB_EXT) tactic/arith/arith_tactics$(LIB_EXT) tactic/core/core_tactics$(LIB_EXT) smt/params/smt_params$(LIB_EXT) ast/pattern/pattern$(LIB_EXT) ast/rewriter/bit_blaster/bit_blaster$(LIB_EXT) ast/fpa/fpa$(LIB_EXT) ackermannization/ackermannization$(LIB_EXT) tactic/aig/aig_tactic$(LIB_EXT) parsers/smt2/smt2parser$(LIB_EXT) cmd_context/cmd_context$(LIB_EXT) solver/solver$(LIB_EXT) ast/proofs/proofs$(LIB_EXT) parsers/util/parser_util$(LIB_EXT) ast/substitution/substitution$(LIB_EXT) tactic/tactic$(LIB_EXT) model/model$(LIB_EXT) ast/normal_forms/normal_forms$(LIB_EXT) ast/macros/macros$(LIB_EXT) ast/rewriter/rewriter$(LIB_EXT) math/lp/lp$(LIB_EXT) nlsat/nlsat$(LIB_EXT) sat/sat$(LIB_EXT) math/grobner/grobner$(LIB_EXT) ast/euf/euf$(LIB_EXT) params/params$(LIB_EXT) ast/ast$(LIB_EXT) math/subpaving/subpaving$(LIB_EXT) math/realclosure/realclosure$(LIB_EXT) math/automata/automata$(LIB_EXT) math/hilbert/hilbert$(LIB_EXT) math/simplex/simplex$(LIB_EXT) math/dd/dd$(LIB_EXT) math/interval/interval$(LIB_EXT) math/polynomial/polynomial$(LIB_EXT) util/util$(LIB_EXT) $(SLINK_EXTRA_FLAGS)
	ln -f -s ../libz3$(SO_EXT) python
api_dll: libz3$(SO_EXT)

example$(OBJ_EXT): ../examples/c++/example.cpp
	$(CXX) $(CXXFLAGS) $(OS_DEFINES) $(EXAMP_DEBUG_FLAG) $(CXX_OUT_FLAG)example$(OBJ_EXT) $(LINK_FLAGS) -I../src/api -I../src/api/c++ ../examples/c++/example.cpp
cpp_example$(EXE_EXT): libz3$(SO_EXT) example$(OBJ_EXT) 
	$(LINK) $(LINK_OUT_FLAG)cpp_example$(EXE_EXT) $(LINK_FLAGS) example$(OBJ_EXT)  libz3$(SO_EXT) $(LINK_EXTRA_FLAGS)
_ex_cpp_example: cpp_example$(EXE_EXT)

tptp5$(OBJ_EXT): ../examples/tptp/tptp5.cpp
	$(CXX) $(CXXFLAGS) $(OS_DEFINES) $(EXAMP_DEBUG_FLAG) $(CXX_OUT_FLAG)tptp5$(OBJ_EXT) $(LINK_FLAGS) -I../src/api -I../src/api/c++ ../examples/tptp/tptp5.cpp
tptp5.lex$(OBJ_EXT): ../examples/tptp/tptp5.lex.cpp
	$(CXX) $(CXXFLAGS) $(OS_DEFINES) $(EXAMP_DEBUG_FLAG) $(CXX_OUT_FLAG)tptp5.lex$(OBJ_EXT) $(LINK_FLAGS) -I../src/api -I../src/api/c++ ../examples/tptp/tptp5.lex.cpp
z3_tptp$(EXE_EXT): libz3$(SO_EXT) tptp5$(OBJ_EXT) tptp5.lex$(OBJ_EXT) 
	$(LINK) $(LINK_OUT_FLAG)z3_tptp$(EXE_EXT) $(LINK_FLAGS) tptp5$(OBJ_EXT) tptp5.lex$(OBJ_EXT)  libz3$(SO_EXT) $(LINK_EXTRA_FLAGS)
_ex_z3_tptp: z3_tptp$(EXE_EXT)

test_capi$(OBJ_EXT): ../examples/c/test_capi.c
	$(CC) $(CFLAGS) $(OS_DEFINES) $(EXAMP_DEBUG_FLAG) $(C_OUT_FLAG)test_capi$(OBJ_EXT) $(LINK_FLAGS) -I../src/api ../examples/c/test_capi.c
c_example$(EXE_EXT): libz3$(SO_EXT) test_capi$(OBJ_EXT) 
	$(LINK) $(LINK_OUT_FLAG)c_example$(EXE_EXT) $(LINK_FLAGS) test_capi$(OBJ_EXT)  libz3$(SO_EXT) $(LINK_EXTRA_FLAGS)
_ex_c_example: c_example$(EXE_EXT)

maxsat$(OBJ_EXT): ../examples/maxsat/maxsat.c
	$(CC) $(CFLAGS) $(OS_DEFINES) $(EXAMP_DEBUG_FLAG) $(C_OUT_FLAG)maxsat$(OBJ_EXT) $(LINK_FLAGS) -I../src/api ../examples/maxsat/maxsat.c
maxsat$(EXE_EXT): libz3$(SO_EXT) maxsat$(OBJ_EXT) 
	$(LINK) $(LINK_OUT_FLAG)maxsat$(EXE_EXT) $(LINK_FLAGS) maxsat$(OBJ_EXT)  libz3$(SO_EXT) $(LINK_EXTRA_FLAGS)
_ex_maxsat: maxsat$(EXE_EXT)

_ex_py_example: 

install:                                                                   z3$(EXE_EXT)  libz3$(SO_EXT)                
	@mkdir -p $(DESTDIR)$(PREFIX)/bin
	@mkdir -p $(DESTDIR)$(PREFIX)/include
	@mkdir -p $(DESTDIR)$(PREFIX)/lib
	@cp ../src/util/z3_version.h $(DESTDIR)$(PREFIX)/include/z3_version.h
	@cp ../src/api/z3.h $(DESTDIR)$(PREFIX)/include/z3.h
	@cp ../src/api/z3_v1.h $(DESTDIR)$(PREFIX)/include/z3_v1.h
	@cp ../src/api/z3_macros.h $(DESTDIR)$(PREFIX)/include/z3_macros.h
	@cp ../src/api/z3_api.h $(DESTDIR)$(PREFIX)/include/z3_api.h
	@cp ../src/api/z3_ast_containers.h $(DESTDIR)$(PREFIX)/include/z3_ast_containers.h
	@cp ../src/api/z3_algebraic.h $(DESTDIR)$(PREFIX)/include/z3_algebraic.h
	@cp ../src/api/z3_polynomial.h $(DESTDIR)$(PREFIX)/include/z3_polynomial.h
	@cp ../src/api/z3_rcf.h $(DESTDIR)$(PREFIX)/include/z3_rcf.h
	@cp ../src/api/z3_fixedpoint.h $(DESTDIR)$(PREFIX)/include/z3_fixedpoint.h
	@cp ../src/api/z3_optimization.h $(DESTDIR)$(PREFIX)/include/z3_optimization.h
	@cp ../src/api/z3_fpa.h $(DESTDIR)$(PREFIX)/include/z3_fpa.h
	@cp ../src/api/z3_spacer.h $(DESTDIR)$(PREFIX)/include/z3_spacer.h
	@cp z3$(EXE_EXT) $(DESTDIR)$(PREFIX)/bin/z3$(EXE_EXT)
	@cp libz3$(SO_EXT) $(DESTDIR)$(PREFIX)/lib/libz3$(SO_EXT)
	@cp ../src/api/c++/z3++.h $(DESTDIR)$(PREFIX)/include/z3++.h
	@echo Z3 was successfully installed.

uninstall:
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_version.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_v1.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_macros.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_api.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_ast_containers.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_algebraic.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_polynomial.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_rcf.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_fixedpoint.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_optimization.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_fpa.h
	@rm -f $(DESTDIR)$(PREFIX)/include/z3_spacer.h
	@rm -f $(DESTDIR)$(PREFIX)/bin/z3$(EXE_EXT)
	@rm -f $(DESTDIR)$(PREFIX)/lib/libz3$(SO_EXT)
	@rm -f $(DESTDIR)$(PREFIX)/lib/libz3$(LIB_EXT)
	@rm -f $(DESTDIR)$(PREFIX)/include/z3++.h
	@echo Z3 was successfully uninstalled.

