export const divisions: string[] = [
  'Diamond',
  'Platinum',
  'Gold',
  'Silver',
  'Bronze',
  'Steel',
  'Wood'
];

// expects divisions to have first letter uppercase
export function compareDivisions(a: string, b: string): number {
  const ai = divisions.indexOf(a);
  const bi = divisions.indexOf(b);
  if (ai === bi) return 0;
  return ai > bi ? 1 : -1;
}
