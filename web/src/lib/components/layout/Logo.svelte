<script>
  import bg from '$lib/assets/logo/bg.png';
  import mask from '$lib/assets/logo/mask.png';
  import overlay from '$lib/assets/logo/overlay.png';
</script>

<div class="invisible size-128 animate-[spin_120s_linear_infinite] select-none xl:visible">
  <img class="absolute z-10 select-none" src={overlay} alt="" draggable="false" />
  <div
    class="mask-alpha mask-center mask-no-repeat select-none"
    style={`mask-image: url(${mask})`}
    draggable="false">
    <img
      class="animate-[reverse-spin_120s_linear_infinite] select-none"
      src={bg}
      alt=""
      draggable="false" />
  </div>
</div>
