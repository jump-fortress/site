<script lang="ts">
  type Props = {
    label: string;
    submitInput: (val: string) => InputError;
  };

  let { label, submitInput }: Props = $props();
  let value: string = $state('');

  let response: InputError = $state({
    error: false,
    message: ''
  });
</script>

<!-- todo: response validation / rejection feedback -->
<div class="flex items-center gap-2">
  <label
    for={label}
    class="relative mt-2 w-80 border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50"
  >
    <span class="absolute -top-0.5 left-2 bg-jfgray-800 px-1 text-base leading-1">{label} </span>
    <div class="flex h-10">
      <input
        bind:value
        onkeypress={(event: KeyboardEvent) => {
          if (event.key === 'Enter') {
            const error = submitInput(value);
            if (error) response = error;
          }
        }}
        class="group peer w-full bg-jfgray-800 bg-clip-padding p-1 text-ctp-lavender transition-colors focus:bg-jfgray-900"
        id={label}
        type="text"
      />
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button
        onclick={() => {
          response = submitInput(value);
        }}
        class="flex h-full w-12 cursor-pointer items-center justify-center bg-jfgray-800 transition-colors peer-focus:bg-jfgray-800"
      >
        <span
          class="icon-[ri--send-plane-line] {response.error === true
            ? 'text-ctp-red'
            : 'text-ctp-lavender'} transition-colors"
        ></span>
      </button>
    </div>
  </label>
  <span class="{response.error === true ? 'text-ctp-red' : 'text-ctp-lavender'} transition-colors"
    >{response.message}</span
  >
</div>

<!-- <style lang="postcss">
  @reference "../../app.css";

  input {
    @apply group peer w-full bg-jfgray-800 bg-clip-padding p-1 text-ctp-lavender transition-colors focus:bg-jfgray-900;
  }
</style> -->
