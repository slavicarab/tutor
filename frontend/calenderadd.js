const calDate = document.getElementById('date');
const calTime = document.getElementById('time');
const calStudent = document.getElementById('student');
const buttonSave = document.getElementById('save');

let users = [];
let calendar = [];

function saveItemsToLocalStorage(){
    localStorage.setItem("calendar", JSON.stringify(calendar));

};

function loadFromLocalStorage() {
    const localList = localStorage.getItem('users')
    const initialUser = JSON.parse(localList);
    users.push(...initialUser);
    const localCalendar = localStorage.getItem('calendar');
    //console.log(localCalendar)
    const initialCalendar = JSON.parse(localCalendar);
    //console.log(initialCalendar)
    //console.log(...initialCalendar)
    calendar.push(...initialCalendar);
};


document.addEventListener("DOMContentLoaded", function () {
    loadFromLocalStorage();
    loadUserData();
});


buttonSave.addEventListener('click', function(event){
    event.preventDefault();
    const calDateVal = calDate.value;
    const calTimeVal = calTime.value;
    const calStudentVal = calStudent.value;
    const id = Date.now()
    calendar.push ({calDate: calDateVal, calTime: calTimeVal, calStudent: calStudentVal, id: id});
    saveItemsToLocalStorage();
    //redirect to another page
    window.location.href = 'appointments.html';
})


function loadUserData() {
    let html ='';
    for(let user of users){
        html =`
          <option value="${user.userName}">${user.userName}</option>`
          calStudent.innerHTML+=html;
    }
    
};

function myFunction() {
    var x = document.getElementById("myLinks");
    if (x.style.display === "block") {
      x.style.display = "none";
    } else {
      x.style.display = "block";
    }
  };
