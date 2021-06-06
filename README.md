### Visual Studio Anti-Debug

###### Minor exploit in order to disable process debugging with the Visual Studio IDE

This only works on VS, WinDbg and DBX are not affected by this, but you can always prevent those by fully disabling via process checks.

The dll is precompiled with the code in dllCode.cpp, not the full source of the Dll but that is the function code referenced if you have any doubts. Feel free to recompile your own if you are worried about viruses.

![Image of VS while applied](https://i.gyazo.com/394b8c6eab549e50cdd5faa73346fa52.png)
