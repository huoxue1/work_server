package com.gjs.utils;

import com.gjs.controller.FilesController;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.PropertySource;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.zip.ZipEntry;
import java.util.zip.ZipOutputStream;


public class Util {

    private static final org.slf4j.Logger logger = LoggerFactory.getLogger(Util.class);


    public static void compressSingleFile(String filePath, String outPut) {
        try {
            File file = new File(filePath);
            String zipFileName = file.getName().concat(".zip");
            System.out.println("zipFileName:" + zipFileName);

            // if you want change the menu of output ,just fix here
            // FileOutputStream fos = new FileOutputStream(zipFileName);
            FileOutputStream fos = new FileOutputStream(outPut + File.separator + zipFileName);

            ZipOutputStream zos = new ZipOutputStream(fos);

            zos.putNextEntry(new ZipEntry(file.getName()));

            byte[] bytes = Files.readAllBytes(Paths.get(filePath));
            zos.write(bytes, 0, bytes.length);
            zos.closeEntry();
            zos.close();

        } catch (FileNotFoundException ex) {
            logger.error("The file "+filePath+" does not exist");
        } catch (IOException ex) {
           logger.error("I/O error: " + ex);
        }
    }


    public static void compressMultipleFiles(String outPath,String... filePaths) {
        try {



            FileOutputStream fos = new FileOutputStream(outPath);
            ZipOutputStream zos = new ZipOutputStream(fos);

            for (String aFile : filePaths) {
                zos.putNextEntry(new ZipEntry(new File(aFile).getName()));

                byte[] bytes = Files.readAllBytes(Paths.get(aFile));
                zos.write(bytes, 0, bytes.length);
                zos.closeEntry();
            }

            zos.close();

        } catch (FileNotFoundException ex) {
            logger.error("A file does not exist: " + ex);
        } catch (IOException ex) {
            logger.error("I/O error: " + ex);
        }
    }
}
