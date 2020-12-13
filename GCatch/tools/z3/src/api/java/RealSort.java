/**
Copyright (c) 2012-2014 Microsoft Corporation
   
Module Name:

    RealSort.java

Abstract:

Author:

    @author Christoph Wintersteiger (cwinter) 2012-03-15

Notes:
    
**/ 

package com.microsoft.z3;

/**
 * A real sort
 **/
public class RealSort extends ArithSort
{
    RealSort(Context ctx, long obj)
    {
        super(ctx, obj);
    }

    RealSort(Context ctx)
    {
        super(ctx, Native.mkRealSort(ctx.nCtx()));
    }
}
