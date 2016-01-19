# FileSystem
Имплементация на файлова система написана на Go.
##Работа с файлове и директории:

#Get_file
Get_file(path string) (File, error)
Функцията Get_file връща файл по зададен път от потребителя, ако такъв не съществува връща error.

#Get_dir
Get_dir(path string) (File, error)
Функцията Get_dir връща директория от файловата система по зададен път от потребителя, ако такaвa не съществува връща error.

#Create
Create(name string) ()
Функцията Create създава нов файл в текущата директория, в която се намираме.

#Mkdir
Mkdir(name string) ()
Функцията Mkdir създава нова директория в текущата.

#Pwd
Pwd() (Directory)
Функцията Pwd връща текущата директория.

#Remove
Remove(path string, dir bool) ()
Функцията Remove премахва файл или директория от файловата система.

#Move
Move(source Directory, destination Directory) ()
Функцията Move мести файл от директория source в destination.

#Link
Link(source string, destination string, symbolic bool) ()
Функцията Link създава твърда или символна връзка.

##Функции за работа с файловата система:

#Create_FS
Create_FS (size int) (FileSystem)
Създава нова файлова система с размер size.Щом бъде създадена тя разполага със size байта и при надскачане на това ограничение ще бъдат извиквани error-и.

#Mount
Mount (a FileSystem, path string) (FileSystem, error)
Mонтира друга файлова система, към текущата в path. Ако не може хвърля error.

#Unmount
Unmount (path string) (FileSystem, error)
Премахва дадена директория от файловата система,заедно с всички нейни "деца". Ако няма такава, връща error.
