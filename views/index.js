const todos = new Vue({
  el:'#todo_list',
  data:{
    todoList:""
  },
  methods:{
    remove:function(event){
      const id = event.target.value;
      fetch(`/todos?id=${id}`, {
        method: 'DELETE'
      }).then(response => {
        if (response.ok) {
          alert('削除しました');
          getTodoList();
        }
      }).catch((err) => {
        console.log(err);
      })
    }
  }
})

getTodoList();

const registerButton = document.getElementById(`register`);
registerButton.onclick = function() {
  const form = {
    schedule : document.schedule_data.schedule.value,
    priority : document.schedule_data.priority.value,
    timeLimit : document.schedule_data.time_limit.value
  };
  
  fetch('/todos', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(form)
  }).then(response => {
    if (response.ok) {
      alert('登録に成功しました');
      document.schedule_data.schedule.value = "";
      document.schedule_data.time_limit.value = "";
      getTodoList();
    }
  }).catch((err) => {
    console.log(err);
  })
};

function getTodoList() {
  fetch('/todos')
    .then(response => {
      if (response.ok) {
        return response.json();
      }
    })
    .then(todoList => {
      todos.todoList = todoList
    }).catch((err) => {
      console.log(err);
    })
}