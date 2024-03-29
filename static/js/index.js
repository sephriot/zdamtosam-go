// Import the functions you need from the SDKs you need
import {initializeApp} from "https://www.gstatic.com/firebasejs/9.9.4/firebase-app.js";
import {getAnalytics} from "https://www.gstatic.com/firebasejs/9.9.4/firebase-analytics.js";
import {
    FacebookAuthProvider,
    getAuth,
    GoogleAuthProvider,
    signInWithPopup
} from 'https://www.gstatic.com/firebasejs/9.9.4/firebase-auth.js'

const firebaseConfig = {
    apiKey: "AIzaSyBT_BvlXPW3aoKBnnjVZqzfCypMGlh2pUU",
    authDomain: "zdamtosam.pl",
    projectId: "zdamtosam-312622",
    storageBucket: "zdamtosam-312622.appspot.com",
    messagingSenderId: "300966227960",
    appId: "1:300966227960:web:1927fcaa1325514427300d",
    measurementId: "G-Z8D6DPFNYQ"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
const auth = getAuth(app)
auth.useDeviceLanguage();

auth.onAuthStateChanged(async (user) => {
    if(user) {
        setCookie("__session", await user.getIdToken())
    } else {
        setCookie("__session", "")
    }
})

function setCookie(cname, cvalue) {
    const d = new Date();
    d.setTime(d.getTime() + (60 * 60 * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

const federatedLogin = (provider) => {
    signInWithPopup(auth, provider)
        .then(() => {
            location.href = "/"
        })
        .catch((error) => {
            if (error.code === 'auth/account-exists-with-different-credential') {
                alert("Konto o podanym adresie email już istnieje.\nSpróbuj zalogować się z użyciem konta z innego portalu.")
            } else {
                alert("Logowanie zakończone niepowodzeniem. Aby uzyskać pomoc skontaktuj się z administratorem strony.")
            }
        });
}

window.logout = () => {
    auth.signOut()
        .then(function() {
            location.href = "/"
        })
        .catch(function(error) {
        });
}

window.googleLogin = () => {
    const provider = new GoogleAuthProvider();
    provider.addScope("email");
    federatedLogin(provider)
}

window.facebookLogin = () => {
    const provider = new FacebookAuthProvider();
    provider.addScope("email");
    provider.addScope("public_profile");
    provider.setCustomParameters({
        'display': 'popup'
    });
    federatedLogin(provider)
}
