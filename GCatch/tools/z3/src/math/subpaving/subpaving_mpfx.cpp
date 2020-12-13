/*++
Copyright (c) 2012 Microsoft Corporation

Module Name:

    subpaving_mpfx.cpp

Abstract:

    Subpaving for non-linear arithmetic using mpfx numerals.

Author:

    Leonardo de Moura (leonardo) 2012-09-18.

Revision History:

--*/
#include "math/subpaving/subpaving_mpfx.h"
#include "math/subpaving/subpaving_t_def.h"

// force template instantiation
template class subpaving::context_t<subpaving::config_mpfx>;
