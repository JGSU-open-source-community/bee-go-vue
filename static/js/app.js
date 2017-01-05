(function(Vue){
	"use strict"

	new Vue({
		// A Dom element to mount our view model.
		el:  'body',
        
        // This ia the model (javascript object)
        // Define properties and give them initial values, here is empey values.
		data:{
			tasks: [],
			newTask: {}
		},
        
        // Functions we will be using
		methods: {
			createTask: function() {
			    if (!$.trim(this.newTask.name)) {
			    	this.newTask = {};
			    	return
			    };

			    this.newTask.done = false;
                
                // The http client for vue.js
				this.$http.post('/task',this.newTask).success(function(res) {
					this.newTask.id = res.created;
					this.tasks.push(this.newTask);

					this.newTask = {};
				}).error(function(err) {
					console.log(err)
				});		    
			},

			
		}
	});
});