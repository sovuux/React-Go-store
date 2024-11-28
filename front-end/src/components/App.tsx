import './App.scss'
import { CommonButton } from './common/Button/button'

export function App() {
	return (
		<>
			<CommonButton
				text='Кнопка'
				type='button'
				onClick={() => console.log(1)}
			/>
		</>
	)
}
