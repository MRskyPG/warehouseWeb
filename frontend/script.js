<script type="text/javascript">
	$(document).ready(function() {
	$('.buttonGive').on('click', function() {
		var id = $(this).val();
		$.ajax({
			data:{'id':id},
			url: "http://localhost/giveItem",
			method: "GET",
		});
		alert("Normalin");
	});
});
</script>
<script type="text/javascript">
	$(document).ready(function() {
	$(".buttonChangePlacement").on('click', function() {
		var id = $(this).val();
		$.ajax({
			data:{'id':id},
			url: "http://localhost/changeItemPlacement",
			method: "GET",
			success: function(data) {
				$("#response").html(data);
				id = data.position;
			}
		});
		alert(id);
	});
});
</script>