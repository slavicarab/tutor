
const tableUsers = document.getElementById('usersTable');
const result = document.getElementById('result');

const apiUrl = "http://localhost:8080/users";

//funkcija koja pravi tabelu u html
async function makeTable(){
    const response = await fetch(apiUrl);
    const users = await response.json();
    console.log(users)
    let html="";
        for(let user of users){
               let status = '';
               (user.userStatus) ? status = 'active' : status = 'inactive'
                html = `
                <tr>
                    <td>${user.userName}</td>
                    <td>${user.userClass}</td>
                    <td>${user.userAddress}</td>
                    <td>${user.userNumber}</td>
                    <td>${user.userEmail}</td>
                    <td>${user.userCourse}</td>
                    <td>${status}</td>
                    <td>
                        <button class="change" onclick='deleteUser(${user.id})'>Delete</button>
                        <button class="change" onclick='editUser(${user.id})'>Edit</button>
                    </td>
                </tr>`
               
             result.innerHTML+=html;
 
        }
};

//funkcija koja brise podatke iz naseg niza i iz lokal storage
async function deleteUser(id){
    deleteResult();
    await fetch(`http://localhost:8080/users/${id}`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({id: id}),
    });

    makeTable();
   
};

//funkcija koja brise podatke iz niza
function deleteResult (){
    while(result.firstChild) result.removeChild(result.firstChild);
};
document.addEventListener("DOMContentLoaded", function () {
    makeTable();
});

//funkcija koja nas vodi na stranu edituser
function editUser(id){
    window.location.href = `edituser.html?id=${id}`;
}

//funkcija za hamburger menu
function myFunction() {
    var x = document.getElementById("myLinks");
    if (x.style.display === "block") {
      x.style.display = "none";
    } else {
      x.style.display = "block";
    }
};
