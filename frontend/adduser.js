const userForm = document.getElementById('tutorForm');
const userName = document.getElementById('name');
const userClass = document.getElementById('class');
const userAddress = document.getElementById('address');
const userNumber = document.getElementById('phoneNumber');
const userEmail = document.getElementById('email');
const userCourse = document.getElementById('course');
const buttonSave = document.getElementById('save');

const apiUrl = "http://localhost:8080/users";

buttonSave.addEventListener('click', async function(event){
    event.preventDefault()
    const userNameVal = userName.value.trim();
    const userClassVal = userClass.value.trim();
    const userAddressVal = userAddress.value.trim();
    const userNumberVal = userNumber.value.trim();
    const userEmailVal = userEmail.value.trim();
    const userCourseVal = userCourse.value.trim();
    const userStatusValString = document.querySelector('input[name="status"]:checked').value; 
    const userStatusVal = userStatusValString === "true";

    await fetch(apiUrl, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({userName: userNameVal, userClass: userClassVal, userAddress: userAddressVal, userNumber: userNumberVal, userEmail: userEmailVal, userCourse: userCourseVal, userStatus: userStatusVal,}),
    });

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

