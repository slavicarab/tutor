const userName = document.getElementById('name');
const userClass = document.getElementById('class');
const userAddress = document.getElementById('address');
const userNumber = document.getElementById('phoneNumber');
const userEmail = document.getElementById('email');
const userCourse = document.getElementById('course');
const buttonSave = document.getElementById('save');
let id = '';


document.addEventListener("DOMContentLoaded", function () {
    const url = new URL(window.location.href);
    id = url.searchParams.get("id");
    loadUserData(parseInt(id));
});

async function loadUserData(id){
  const response = await fetch(`http://localhost:8080/users/${id}`);
  const user = await response.json();
        
    userName.value = user.userName;
    userClass.value = user.userClass;
    userAddress.value = user.userAddress;
    userNumber.value = user.userNumber;
    userEmail.value = user.userEmail;
    userCourse.value = user.userCourse;
    if (user.userStatus === true) {
      document.getElementById("userStatusTrue").checked = true;
    } else if (user.userStatus === false) {
      document.getElementById("userStatusFalse").checked = true;
    }

};

buttonSave.addEventListener('click', async function(event){
    event.preventDefault()
    const userNameVal = userName.value;
    const userClassVal = userClass.value;
    const userAddressVal = userAddress.value;
    const userNumberVal = userNumber.value;
    const userEmailVal = userEmail.value;
    const userCourseVal = userCourse.value;
    const userStatusValString = document.querySelector('input[name="status"]:checked').value; 
    const userStatusVal = userStatusValString === "true";
    
    await fetch(`http://localhost:8080/users/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({userName: userNameVal, userClass: userClassVal, userAddress: userAddressVal, userNumber: userNumberVal, userEmail: userEmailVal, userCourse: userCourseVal, userStatus: userStatusVal}),
    })
    .then((response) => response.json())
    .then((data) => console.log("User updated:", data))
    .catch((error) => console.error("Error:", error));
    //redirect to another page
    window.location.href = 'listusers.html';
});

function myFunction() {
    var x = document.getElementById("myLinks");
    if (x.style.display === "block") {
      x.style.display = "none";
    } else {
      x.style.display = "block";
    }
  }