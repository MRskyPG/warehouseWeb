<html>
    <script type="text/javascript">
			$(document).ready(function() {
				$("#buttonChangePlacement").on('click', function() {
					$.ajax({
						url: "http://localhost/changeItemPlacement",
						method: "GET",
						success: function(data) {
							$("#response").html(data);
						},
					});
				});
			});
</script>
</html>