const Login: React.FC = () => {
    const clientId: string = import.meta.env.VITE_GITHUB_CLIENT_ID as string

    return(
        <>
        <div>Khoury Classroom</div>
        <a href={`https://github.com/login/oauth/authorize?client_id=${clientId}&scope=repo,read:org,classroom&allow_signup=false`}>
            Sign In</a>
            </>
    )
} 

export default Login;