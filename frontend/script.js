<script type="text/javascript">
			$(document).ready(function() {
				$("#buttonGive").on('click', function() {
					var id = $("#buttonGive").val();
					$.ajax({
						data:{'id':id},
						url: "http://localhost/giveItem",
						method: "GET",
						success: function(data) {
							$("#response").html(data);
						},
					});
				});
			});
</script>
<script type="text/javascript">
			$(document).ready(function() {
				$("#buttonRemove").on('click', function() {
					var id = $("#buttonRemove").val();
					$.ajax({
						data:{'id':id},
						url: "http://localhost/removeItem",
						method: "GET",
						success: function(data) {
							$("#response").html(data);
						},
					});
				});
			});
</script>
<script type="text/javascript">
			$(document).ready(function() {
				$("#buttonChangePlacement").on('click', function() {
					var id = $("#buttonChangePlacement").val();
					$.ajax({
						data:{'id':id},
						url: "http://localhost/changeItemPlacement",
						method: "GET",
						success: function(data) {
							$("#response").html(data);
						},
					});
				});
			});
</script>