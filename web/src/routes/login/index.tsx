import { createFileRoute } from '@tanstack/react-router'
import LoginPage from '../../ui/pages/login/Login.page'

export const Route = createFileRoute('/login/')({
    component: RouteComponent,
})

function RouteComponent() {
    return <LoginPage />
}
