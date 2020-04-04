import CryptoJS from 'crypto-js';

// sample to use CryptoJS module to encrypt and decrypt

const key = CryptoJS.enc.Utf8.parse('a4e0bdkkksa9a59s');
const iv = CryptoJS.enc.Utf8.parse('141deb09efc61e9b');
const encrypted = CryptoJS.AES.encrypt('java123JDK', key, {
    iv,
    mode: CryptoJS.mode.CBC,
});
const encryptPwd = encrypted.toString();
console.log(encryptPwd);
if (this.form.account) {
    const decrypt = CryptoJS.AES.decrypt('rDqcSr0sRNRMcgJIeRvPaQ==', key, {
        iv,
        mode: CryptoJS.mode.CBC,
    });
    const decryptedStr = decrypt.toString(CryptoJS.enc.Utf8);
    console.log(decryptedStr);
}