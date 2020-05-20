
const registerButton = document.getElementById(`register`);
registerButton.onclick = function(){
    const form = {};
    form.schedule = document.schedule_data.schedule.value;
    form.priority = document.schedule_data.priority.value;
    form.timeLimit = document.schedule_data.time_limit.value;
    const JsonData = JSON.stringify(form)
    fetch('/register',{
        method:'post',
        body:JsonData
    })
    .then()
    /*
    const table = document.getElementById(`todo_list`);
    let row =  table.insertRow(-1);
    let cell = row.insertCell(-1);
    let text = document.createTextNode(scheduleName);
    cell.appendChild(text);
    cell = row.insertCell(-1);
    text = document.createTextNode(priorityValue);
    cell.appendChild(text);
    cell = row.insertCell(-1);
    text = document.createTextNode(timeLimitName);
    cell.appendChild(text);
    cell = row.insertCell(-1);
    let button = document.createElement('button');
    button.type = 'button';

    button.innerText = "削除";
    
    cell.appendChild(button);
    button.addEventListener("click", removeSchedule, false);
    */
    refreshIndex();
}

function removeSchedule(event){
    const val = event.target.value;
    const table = document.getElementById(`todo_list`);
    table.deleteRow(val);
    refreshIndex();
}

function refreshIndex(){
    const table = document.getElementById(`todo_list`);
    let length = table.rows.length;
    for(let i = 0; i < length; i++){
        table.rows[i].id = 'schedule[' + i +']'
        let btn = table.rows[i].cells[3].getElementsByTagName('button');
        btn[0].value = i;
    }
}
