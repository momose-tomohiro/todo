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

const registerForm = new Vue({
  el:'#schedule_data',
  methods:{
    register:function(){
      const form = {
        schedule : this.schedule,
        priority : this.priority,
        timeLimit : this.time_limit
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
          this.schedule = "";
          this.time_limit = "";
          getTodoList();
        }
      }).catch((err) => {
        console.log(err);
      })
    }
  }
})

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