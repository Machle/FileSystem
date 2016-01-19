type File struct{

}

type Directory struct{

}

type FileSystem interface{

    //Returns a file
    Get_file(path string) (File, error)
    //Returns a directory
    Get_dir(path string) (Directory, error)
    
    //Creates new file
    Create(name string) ()
    //Creates new directory
    Mkdir(name string) ()
    //Pwd ... 
    Pwd() (Directory)
    //Removes file
    Remove(f File) ()
    //Removes directory
    Remove(d Driectory) ()
    //Moves a file from source to destination
    Move(source Directory, destination Directory) ()

    // Creates hard or symbolic link
    Link(source string, destination string, symbolic bool) ()
}
// Creates a File_system with size
Create_FS (size int) (FileSystem)
// Mounts other file system to ours
Mount (a FileSystem, path string) (FileSystem, error)
//Unmounts a directory from our file system
Unmount (path string) (FileSystem, error)