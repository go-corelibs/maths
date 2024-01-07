#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += MATHS
MATHS_MK_SUMMARY := go-corelibs/maths
MATHS_MK_VERSION := v1.0.1

include CoreLibs.mk
