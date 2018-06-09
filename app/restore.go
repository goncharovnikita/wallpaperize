package main

// Restore set wallpaper to its original file
func (a application) Restore() {
	fpath := a.rec.getRecoverFilepath()
	a.master.SetFromFile(fpath)
}
