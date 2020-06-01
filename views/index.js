const todos = new Vue({
  el: "#todoPage",
  vuetify: new Vuetify(),
  data: {
      header:[
        {
          text:'予定',
          align: 'center',
          value: 'schedule'
        },
        {
          text:'優先度',
          align: 'center',
          value:'priority'
        },
        {
          text:'期限',
          align: 'center',
          value:'time_limit'
        },
        {
          text:'削除',
          align: 'center',
          value:'action',
          sortable:false
        }
      ],
      todoList: [
        {
          id: "",
          schedule: "",
          priority: "",
          time_limit: ""
        },
      ]
    },
  mounted: function () {
    this.display();
  },
  methods: {
    display: function () {
      fetch("/todos")
        .then((response) => {
          if (response.ok) {
            return response.json();
          }
        })
        .then((result) => {
          todos.todoList = result;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    register: function () {
      const form = {
        schedule: this.todoList.schedule,
        priority: this.todoList.priority,
        time_limit: this.todoList.time_limit,
      };

      fetch("/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(form),
      })
        .then((response) => {
          if (response.ok) {
            alert("登録に成功しました");
            this.todoList.schedule = "";
            this.todoList.time_limit = "";
            this.display();
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    remove: function (item) {
      fetch(`/todos?id=${item.id}`, {
        method: "DELETE",
      })
        .then((response) => {
          if (response.ok) {
            alert("削除しました");
            this.display();
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
});
