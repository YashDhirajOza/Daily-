import java.io.FileInputStream;
import java.io.BufferedInputStream;
import java.io.IOException;

public class BinaryReadExample {

    public static void main(String[] args) {
        String filename = "sample.txt";

        // Wrap FileInputStream in BufferedInputStream for better performance
        try (FileInputStream fis = new FileInputStream(filename);
             BufferedInputStream bis = new BufferedInputStream(fis)) {

            // We'll read chunks of up to 4 KB
            byte[] buffer = new byte[4096];
            int bytesRead;

            // StringBuilder to accumulate the text
            StringBuilder sb = new StringBuilder();

            // Read until EOF (read() returns -1)
            while ((bytesRead = bis.read(buffer)) != -1) {
                // Convert exactly the bytes we read into a String
                String chunk = new String(buffer, 0, bytesRead, "UTF-8");
                sb.append(chunk);
            }

            // Print out the whole file
            System.out.println("=== File Contents ===");
            System.out.println(sb.toString());

        } catch (IOException e) {
            System.err.println("I/O error while reading " + filename + ": " + e.getMessage());
            e.printStackTrace();
        }
    }
}
