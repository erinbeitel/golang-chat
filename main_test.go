func checkPort (port int) error {
  
  if port > 65535 {
    return errors.New("The port cannot be greater than 65535.")
  } else if port < 0 {
    return errors.New("The web server port cannot be negative.")
  } else if port < 1024 {
    return errors.New("Ports have to be between 1024 and 65535")
  } else {
    return nil
  }
  
}

func checkDir (dir string) error {
  
  if (len(dir) == 0) {
    return errors.New("The length of the directory string was zero.")
  } else if dir == "/" {
    return errors.New("Can't run the webserver from the filesystem root, this was probably an accident.")
  } else {
    return nil
  }
  
}

func TestCheckPort (t *testing.T) {
  
  failMessage := "Failed testing of checkPort() function."
  
  if checkPort(-1) == nil {
    t.Errorf(failMessage)
  } else if checkPort(00000000) == nil {
    t.Errorf(failMessage)
  } else if checkPort(99999999) == nil {
    t.Errorf(failMessage)
  }
  
}


func TestCheckDir (t *testing.T) {
  
  failMessage := "Failed testing of checkDir() function."
  
  if checkDir("/") == nil {
    t.Errorf(failMessage)
  } else if checkDir("usr/bin/test") == nil {
    t.Errorf(failMessage)
  }
  
}

