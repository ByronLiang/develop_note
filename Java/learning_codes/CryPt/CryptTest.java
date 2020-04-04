import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.io.UnsupportedEncodingException;
import java.security.InvalidAlgorithmParameterException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;

public class CryptTest {

    public static void main(String[] args) {

        String msg = "java123JDK";
        String key = "a4e0bdkkksa9a59s";
        String iv = "141deb09efc61e9b";

        String encrypt = AESCrypt.encryptCBC(msg, key, iv);
        System.out.println(encrypt);
        String decryptStr = AESCrypt.decryptCBC(encrypt, key, iv);
        System.out.println(decryptStr);
        // raise error and catch it
        String sec = AESCrypt.decryptCBC("rDqcSr0sRNRMcgJIeRvPaQ==", key, iv);
        System.out.println(sec);
    }
}


class AESCrypt {

    public static String encryptCBC(String message, String key, String iv) {
        final String cipherMode = "AES/CBC/PKCS5Padding";
        final String charsetName = "UTF-8";
        try {
            byte[] content = new byte[0];
            content = message.getBytes(charsetName);
            //
            byte[] keyByte = key.getBytes(charsetName);
            SecretKeySpec keySpec = new SecretKeySpec(keyByte, "AES");
            //
            byte[] ivByte = iv.getBytes(charsetName);
            IvParameterSpec ivSpec = new IvParameterSpec(ivByte);

            Cipher cipher = Cipher.getInstance(cipherMode);
            cipher.init(Cipher.ENCRYPT_MODE, keySpec, ivSpec);
            byte[] data = cipher.doFinal(content);
            final Base64.Encoder encoder = Base64.getEncoder();
            final String result = encoder.encodeToString(data);
            return result;
        } catch (UnsupportedEncodingException e) {
            e.printStackTrace();
        }
        catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (NoSuchPaddingException e) {
            e.printStackTrace();
        } catch (InvalidKeyException e) {
            e.printStackTrace();
        } catch (InvalidAlgorithmParameterException e) {
            e.printStackTrace();
        } catch (IllegalBlockSizeException e) {
            e.printStackTrace();
        } catch (BadPaddingException e) {
            e.printStackTrace();
        }

        return null;
    }

    public static String decryptCBC(String messageBase64, String key, String iv) {
        final String cipherMode = "AES/CBC/PKCS5Padding";
        final String charsetName = "UTF-8";
        try {
            final Base64.Decoder decoder = Base64.getDecoder();
            byte[] messageByte = decoder.decode(messageBase64);

            //
            byte[] keyByte = key.getBytes(charsetName);
            SecretKeySpec keySpec = new SecretKeySpec(keyByte, "AES");
            //
            byte[] ivByte = iv.getBytes(charsetName);
            IvParameterSpec ivSpec = new IvParameterSpec(ivByte);

            Cipher cipher = Cipher.getInstance(cipherMode);
            cipher.init(Cipher.DECRYPT_MODE, keySpec, ivSpec);
            byte[] content = cipher.doFinal(messageByte);
            String result = new String(content, charsetName);
            return result;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }
}
