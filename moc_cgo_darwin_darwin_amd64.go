// +build !ios

package main

/*
#cgo CFLAGS: -pipe -O2 -arch x86_64 -isysroot /Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.12 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -arch x86_64 -isysroot /Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.12 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../syahrul12345 -I. -I/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib/QtGui.framework/Headers -I/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib/QtQml.framework/Headers -I/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib/QtNetwork.framework/Headers -I/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib/QtCore.framework/Headers -I. -I/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/AGL.framework/Headers -I/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/mkspecs/macx-clang -F/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib
#cgo LDFLAGS: -stdlib=libc++ -headerpad_max_install_names -arch x86_64 -Wl,-syslibroot,/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.12  -Wl,-rpath,/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib
#cgo LDFLAGS:  -F/Users/nizamsyahrul/Qt5.13.0/5.13.0/clang_64/lib -framework QtGui -framework QtQml -framework QtNetwork -framework QtCore -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
