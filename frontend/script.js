<script type="text/javascript">
	$(document).ready(function() {
	$('.buttonGive').on('click', function() {
		var id = $(this).val();
		$.ajax({
			data:{'id':id},
			url: "http://localhost/giveItem",
			method: "GET",
		});
		alert("Товар выдан.");
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
		var text = "smth";
		if (id == -1) {
			text = "Нет свободного места для перемещения товара.";
		} else {
			text = "Товар перемещен.";
		}
		alert(text);
	});
});
</script>