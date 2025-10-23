<script lang="ts">
	let { route }: { route: string } = $props();
	let login = true;

	$inspect(route);
</script>

<nav
	class="bg-jfgray-900/75 fixed z-50 flex h-16 w-full select-none justify-center backdrop-blur-sm"
>
	<div
		class="w-5xl hover:text-ctp-lavender-50/75 flex items-center justify-between transition-colors"
	>
		<!-- left nav -->
		<div class="flex h-full">
			{@render NavPage('')}
			{@render NavPage('ladder')}
			{@render NavPage('formats')}
			{@render NavPage('help')}
		</div>
		<!-- right nav -->
		<div class="flex h-16 flex-row-reverse items-center">
			{#if !login}
				<img class="cursor-pointer" src="src/lib/assets/static/sits_small.png" alt="" />
			{:else}
				<a class="group px-4" href="/players/1">
					<img
						class="group-hover:border-ctp-lavender-50 border-jfgray-800 size-14 rounded-full border-2 transition-colors"
						src="https://avatars.akamai.steamstatic.com/55f57bf80f543b1e7261f9e0a17b8e12b992de28_full.jpg"
						alt=""
					/>
				</a>
				<a
					class="text-ctp-lavender-50 flex items-center px-2 opacity-75 transition-opacity hover:opacity-100"
					href="/support"
					title="support"
				>
					<span class="icon-[ri--heart-line] size-6"></span>
				</a>
				<a
					class="text-ctp-lavender-50 flex items-center px-2 opacity-75 transition-opacity hover:opacity-100"
					href="https://discord.gg/tusBc64wnv"
					title="discord"
				>
					<span class="icon-[ri--discord-fill] size-6"></span>
				</a>
			{/if}
		</div>
	</div>
</nav>

<hr class="h-14 opacity-0" />

{#snippet NavPage(name: string)}
	<a
		href="/{name === '' ? '' : name}"
		class="hover:text-ctp-lavender-50 flex h-full items-center px-4 transition-colors"
		><div class="relative">
			<span>{name === '' ? 'home' : name}</span>
			{#if (route.includes(name) && name !== '') || (route === '' && name === '')}
				<hr
					class="text-ctp-lavender border-1 absolute left-0 right-0 m-auto flex w-11/12 rounded-full"
				/>
			{/if}
		</div></a
	>
{/snippet}
