<script>
  import overlay from '$lib/assets/static/logo/anim/small/overlay.png';
  import mask from '$lib/assets/static/logo/anim/small/mask.png';
  import bg from '$lib/assets/static/logo/anim/small/bg.png';
</script>

<div class="size-16 animate-[spin_120s_linear_infinite]">
  <img class="absolute z-10 select-none" src={overlay} alt="" draggable="false" />
  <div style={`mask-image: url(${mask})`} class="mask-alpha mask-center mask-no-repeat">
    <img class="animate-[reverse-spin_120s_linear_infinite]" src={bg} alt="" draggable="false" />
  </div>
</div>
