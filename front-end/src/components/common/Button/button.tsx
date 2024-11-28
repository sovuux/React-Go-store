import { ButtonPropsType } from '../../../types/button/buttonPropsType'
import styles from './button.module.scss'

export function CommonButton({
	text,
	onClick,
	disabled,
	type,
}: ButtonPropsType) {
	return (
		<>
			<button
				type={type}
				className={`${styles.button} ${disabled ? styles.disabled : ''}`}
				onClick={onClick}
				disabled={disabled}
			>
				{text}
			</button>
		</>
	)
}
