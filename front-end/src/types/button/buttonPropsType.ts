export type ButtonPropsType = {
	text: string
	onClick: () => void
	disabled?: boolean
	type?: 'button' | 'submit' | 'reset'
}
