<script lang="ts">
  import { fade } from 'svelte/transition';

  type Props = {
    label: string;
    placeholder?: string;
    initialMessage?: string;
    submitInput: (val: string) => Promise<InputError>;
  };

  let { label, placeholder = '', initialMessage = '', submitInput }: Props = $props();
  let value: string = $state('');

  // response is a promise so we can use #await syntax to handle pending state
  let response: Promise<InputError> = $derived(
    Promise.resolve({
      error: false,
      message: initialMessage
    })
  );
</script>

<!-- todo: response validation / rejection feedback -->
<div class="flex h-12 items-center gap-2">
  <label
    for={label}
    class="border-jfgray-700 focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50 relative mt-2 w-80 text-nowrap border-2 transition-colors"
  >
    <span class="bg-jfgray-800 leading-1 absolute -top-0.5 left-2 px-1 text-base">{label} </span>
    <div class="flex h-10">
      <input
        bind:value
        {placeholder}
        onkeypress={async (event: KeyboardEvent) => {
          if (event.key === 'Enter') {
            const error = submitInput(value);
            if (error) response = error;
          }
        }}
        class="bg-jfgray-800 text-ctp-lavender focus:bg-jfgray-900 group peer size-full bg-clip-padding px-1 transition-colors"
        id={label}
        type="text"
      />
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button
        onclick={async () => {
          response = submitInput(value);
        }}
        class="bg-jfgray-800 peer-focus:bg-jfgray-800 flex h-full w-12 cursor-pointer items-center justify-center"
      >
        {#await response}
          <span in:fade class="icon-[ri--loader-3-line] text-ctp-lavender animate-spin"></span>
        {:then response}
          <span
            in:fade
            class="icon-[ri--send-plane-line] {response.error === true
              ? 'text-ctp-red'
              : 'text-ctp-lavender'}"
          ></span>
        {/await}
      </button>
    </div>
  </label>
  {#await response then response}
    <span
      in:fade
      class="text-base {response.error === true
        ? 'text-ctp-red'
        : 'text-ctp-lavender'} transition-colors">{response.message}</span
    >
  {/await}
</div>
