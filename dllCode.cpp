void ErasePEHeaderFromMemory(HANDLE handle) {
	DWORD OldProtect = 0;

	// Change memory protection
	VirtualProtect(handle, 4096, // Assume x86 page size
		PAGE_READWRITE, &OldProtect);

	// Erase the header
	ZeroMemory(handle, 4096);
}