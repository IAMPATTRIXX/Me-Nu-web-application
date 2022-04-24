
// import { getStorage, ref } from "firebase/storage";
// const httpsReference = ref(storage, 'https://firebasestorage.googleapis.com/b/bucket/o/images%20stars.jpg');
// return httpsReference

const storage = firebase.storage();

storage.ref('image.jpg').getDownloadURL()
  .then((url) => {
    // Do something with the URL ...
  })