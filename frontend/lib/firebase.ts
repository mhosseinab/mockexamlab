// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyCqMNBiX-2KId0HEpULe8YAsknCrCVzikY",
  authDomain: "mock-exam-20921.firebaseapp.com",
  projectId: "mock-exam-20921",
  storageBucket: "mock-exam-20921.appspot.com",
  messagingSenderId: "898881542444",
  appId: "1:898881542444:web:948724b09d14bfcd86f489",
  measurementId: "G-EVCQDW3G7Q",
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
export default auth;
