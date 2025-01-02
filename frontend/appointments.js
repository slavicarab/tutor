

const monthNames = [
    "January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
];

const calendarGrid = document.getElementById('calendar-grid');
const monthYearDisplay = document.getElementById('month-year');
const prevMonthButton = document.getElementById('prev-month');
const nextMonthButton = document.getElementById('next-month');

let currentDate = new Date();
let calendar = [];
let users = [];

function checkAppointments(date) {
    for (let i = 0; i < calendar.length; i++) {
        if (calendar[i].calDate === date) {
            console.log(`${calendar[i].calDate}`)
            const time = calendar[i].calTime;
            const user = calendar[i].calStudent
            const val = true
            return [val, time, user]
        }
    }
    return [false, "time", "user"]
}

function renderCalendar(date) {
    const year = date.getFullYear();
    const month = date.getMonth();
    const firstDayOfMonth = new Date(year, month, 1).getDay();
    const lastDateOfMonth = new Date(year, month + 1, 0).getDate();
    const lastDayOfPrevMonth = new Date(year, month, 0).getDate();

    // Update header
    monthYearDisplay.textContent = `${monthNames[month]} ${year}`;

    // Clear the calendar grid
    calendarGrid.innerHTML = '';

    // Add previous month's trailing days
    for (let i = firstDayOfMonth - 1; i >= 0; i--) {
        const day = lastDayOfPrevMonth - i;
        calendarGrid.innerHTML += `<div class="empty">${day}</div>`;
    }

    // Add current month's days
    for (let day = 1; day <= lastDateOfMonth; day++) {
        const today = new Date();
        const myMonth = today.getMonth();
        const myYear = today.getFullYear();
        const myCalendar = `${myYear}-${myMonth + 1}-${day}`;
        let [isA, isTime, isUser] = checkAppointments(myCalendar);
        const isCurrentDay =
            day === today.getDate() &&
            month === today.getMonth() &&
            year === today.getFullYear();
        if (isCurrentDay) {
            calendarGrid.innerHTML += `<div class='current-day day-format'>${day}
            <div class='day-content'></div></div>`
           
        } else if (isA) {
            calendarGrid.innerHTML += `<div class='a-day hover-container day-format'>${day}<div class="hover-text">${isTime} - ${isUser}</div></div>`;
        } else {
            calendarGrid.innerHTML += `<div class='day-format'>${day}</div>`;
        }
        
    }

    // Add next month's leading days
    const totalCells = calendarGrid.children.length;
    const remainingCells = 42 - totalCells; // 6 rows * 7 days
    for (let day = 1; day <= remainingCells; day++) {
        calendarGrid.innerHTML += `<div class="empty">${day}</div>`;
    }
}

// Event listeners for navigation
prevMonthButton.addEventListener('click', () => {
    currentDate.setMonth(currentDate.getMonth() - 1);
    renderCalendar(currentDate);
});

nextMonthButton.addEventListener('click', () => {
    currentDate.setMonth(currentDate.getMonth() + 1);
    renderCalendar(currentDate);
});

// Initial render
/*
function loadFromLocalStorage() {
    const localList = localStorage.getItem('users')
    const initialUser = JSON.parse(localList);
    users.push(...initialUser);
    const localCalendar = localStorage.getItem('calendar');
    //console.log(localCalendar)
    const initialCalendar = JSON.parse(localCalendar);
    //console.log(initialCalendar)
    //console.log(...initialCalendar)
    calendar.push({calDate: "2025-01-02", calTime: "20:57", calStudent: "1735403852646", id: 1735405028542});
};
*/
document.addEventListener("DOMContentLoaded", function () {
    //loadFromLocalStorage();
    renderCalendar(currentDate);
});

const localCalendar = localStorage.getItem('calendar');
const initialCalendar = JSON.parse(localCalendar);
calendar.push(...initialCalendar)
console.log(calendar)


function myFunction() {
    var x = document.getElementById("myLinks");
    if (x.style.display === "block") {
      x.style.display = "none";
    } else {
      x.style.display = "block";
    }
  };