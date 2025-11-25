import { PathsInternalPlayersPreferredclassClassPutParametersPathClass as Class } from '$lib/schema.d.ts';

export function classToEnum(playerClass: 'Soldier' | 'Demo') {
  if (playerClass === 'Soldier') {
    return Class.Soldier;
  }
  return Class.Demo;
}
