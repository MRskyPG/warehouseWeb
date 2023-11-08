<html>
<script type="text/javascript">
			$(document).ready(function() {
				$("#buttonGive").on('click', function() {
					$.ajax({
						url: "http://localhost/giveItem",
						method: "GET",
						success: function(data) {
							$("#response").html(data);
						},
					});
				});
			});
</script>
</html>