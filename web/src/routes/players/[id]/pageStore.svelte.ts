type pageStore = {
	ownProfile: boolean;
	preferredClass: {
		set: boolean;
		class: 'Soldier' | 'Demo';
	};
	setDefault: () => void;
};

export const pageStore: pageStore = $state({
	ownProfile: false,
	preferredClass: {
		set: true,
		class: 'Soldier'
	},
	setDefault: function () {
		pageStore.ownProfile = false;
		pageStore.preferredClass.set = false;
	}
});
