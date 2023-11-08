 <html>
    <script type="text/javascript">
			$(document).ready(function() {
				$("#buttonRemove").on('click', function() {
					$.ajax({
						url: "http://localhost/removeItem",
						method: "GET",
						success: function(data) {
							$("#response").html(data);
						},
					});
				});
			});
</script>
</html>