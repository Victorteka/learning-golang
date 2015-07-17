// You can look for the file size in the FormFile's Header (which is a MIMEHEader).
// For a single file
w.Header.Get("Content-Length")

// read it into a buffer first to get the size
var buff bytes.Buffer
fileSize, err := buff.ReadFrom(f)

/*

   A multipart File will be an *io.SectionReader if it's in memory, or an *os.File if it was written to a temp file

*/
switch f := f.(type) {
case *io.SectionReader:
    fileSize = r.Size()
case *os.File:
    if s, err := f.Stat(); err == nil {
        fileSize = s.Size()
    }
}