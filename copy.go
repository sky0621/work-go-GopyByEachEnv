package GopyByEachEnv

//
// // [MEMO] とりあえず作った版にしても、雑さが・・・
// func copyToEachDir() {
// 	fromDir := config.m["fromDir"]
// 	fromFile := config.m["fromFile"]
// 	from := fromDir + fromFile
// 	toFiles := config.m["to"]
// 	toArray := strings.Split(toFiles, "$")
//
// 	for _, to := range toArray {
// 		src, err := os.Open(from)
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 		defer src.Close()
//
// 		toSplit := strings.Split(to, "\t")
// 		dst, err := os.Create(toSplit[0])
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 		defer dst.Close()
//
// 		fmt.Println("[COPY] " + from + " => " + toSplit[0])
// 		_, err = io.Copy(dst, src)
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 	}
// 	fmt.Println("")
// }
//
// // [MEMO] とりあえず作った版にしても、雑さが・・・
// func replaceEachToFile() {
// 	var fp *os.File
// 	var err error
// 	var fp2 *os.File
// 	var err2 error
// 	toFiles := config.m["to"]
// 	toArray := strings.Split(toFiles, "$")
//
// 	for _, to := range toArray {
// 		toSplit := strings.Split(to, "\t")
// 		fp, err = os.OpenFile(toSplit[0], os.O_APPEND, 0777)
// 		fp2, err2 = os.Create(toSplit[0] + ".tmp")
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 		if err2 != nil {
// 			log.Println(err2)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 		defer fp.Close()
// 		defer fp2.Close()
// 		scanner := bufio.NewScanner(fp)
// 		writer := bufio.NewWriter(fp2)
// 		for scanner.Scan() {
// 			line := scanner.Text()
// 			line = strings.Replace(line, toSplit[1], toSplit[2], -1)
// 			writer.WriteString(line + "\r\n")
// 		}
// 		writer.Flush()
// 		err = scanner.Err()
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 	}
// }
//
// // [MEMO] とりあえず作った版にしても、雑さが・・・
// func renameTmp() {
// 	var err error
// 	toFiles := config.m["to"]
// 	toArray := strings.Split(toFiles, "$")
//
// 	for _, to := range toArray {
// 		toSplit := strings.Split(to, "\t")
// 		err = os.Rename(toSplit[0]+".tmp", toSplit[0])
// 		if err != nil {
// 			log.Println(err)
// 			ExitCode = ExitCodeCopyError
// 			return
// 		}
// 	}
// }
