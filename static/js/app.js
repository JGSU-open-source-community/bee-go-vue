(function(Vue){
	"use strict";

	new Vue({
		// A Dom element to mount our view model.
		el:  'body',
        
        // This ia the model (javascript object)
        // Define properties and give them initial values, here is empey values.
		data:{
			tasks: [],
			newTask: {}
		},

		created: function() {
            this.$http.get('/tasks').then(function(res) {
                this.tasks = res.data.iterms ? res.data.iterms : [];
            });
        },
        
        // Functions we will be using
		methods: {
			createTask: function() {
			    if (!$.trim(this.newTask.name)) {
			    	this.newTask = {};
			    	return;
			    };

			    this.newTask.done = false;

                // The http client for vue.js
				this.$http.post('/task',this.newTask).success(function(res) {
					this.newTask.id = res.created;
					this.tasks.push(this.newTask);

					this.newTask = {};
				}).error(function(err) {
					console.log(err);
				});		    
			},

			deleteTask : function(index)	{
				this.$http.delete('/task/'+index).success(function() {
					this.$http.get('/task').then(function(res) {
						this.tasks = res.data.items ? res.data.items : [];
					});
				}).error(function(err) {
					console.log(err)
				});
			},	

			updateTask: function(task, completed) {
				if (completed) {
					task.done = true;
				}

				this.$http.put('/task', task).success(function(res) {
					this.$http.get('/tasks').then(function(res) {
						this.tasks = res.data.items ? res.data.items : [];
					});
				}).error(function(err) {
					console.log(err)
				});
			}
		}
	});
})(Vue);