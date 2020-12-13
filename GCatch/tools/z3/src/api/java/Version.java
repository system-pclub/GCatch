/**
Copyright (c) 2012-2014 Microsoft Corporation
   
Module Name:

    Version.java

Abstract:

Author:

    @author Christoph Wintersteiger (cwinter) 2012-03-15

Notes:
    
**/ 

package com.microsoft.z3;

/**
 * Version information. 
     * Remarks: Note that this class is static.
 **/
public class Version
{
    /**
     * The major version
     **/
    public static int getMajor()
    {
        Native.IntPtr major = new Native.IntPtr(), minor = new Native.IntPtr(), build = new Native.IntPtr(), revision = new Native.IntPtr();
        Native.getVersion(major, minor, build, revision);
        return major.value;
    }

    /**
     * The minor version
     **/
    public static int getMinor()
    {
        Native.IntPtr major = new Native.IntPtr(), minor = new Native.IntPtr(), build = new Native.IntPtr(), revision = new Native.IntPtr();
        Native.getVersion(major, minor, build, revision);
        return minor.value;
    }

    /**
     * The build version
     **/
    public static int getBuild()
    {
        Native.IntPtr major = new Native.IntPtr(), minor = new Native.IntPtr(), build = new Native.IntPtr(), revision = new Native.IntPtr();
        Native.getVersion(major, minor, build, revision);
        return build.value;
    }

    /**
     * The revision
     **/
    public static int getRevision()
    {
        Native.IntPtr major = new Native.IntPtr(), minor = new Native.IntPtr(), build = new Native.IntPtr(), revision = new Native.IntPtr();
        Native.getVersion(major, minor, build, revision);
        return revision.value;
    }

    /**
     * A full version string
     **/
    public static String getFullVersion()
    {
        return Native.getFullVersion();
    }

    /**
     * A string representation of the version information.
     **/
    public static String getString()
    {
        Native.IntPtr major = new Native.IntPtr(), minor = new Native.IntPtr(), build = new Native.IntPtr(), revision = new Native.IntPtr();
        Native.getVersion(major, minor, build, revision);
        return Integer.toString(major.value) + "." + Integer.toString(minor.value) + "."
                + Integer.toString(build.value) + "." + Integer.toString(revision.value);
    }
}
