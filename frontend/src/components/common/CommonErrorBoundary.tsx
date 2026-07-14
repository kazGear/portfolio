import { Component, ErrorInfo, ReactNode } from 'react';
import { NavigateFunction } from 'react-router-dom';
import { KEYS } from '../../lib/Constants';
import CommonErrorView from './CommonErrorView';

interface ErrorBoundaryProps {
    children: ReactNode;
    navigate: NavigateFunction;
}

interface ErrorBoundaryState {
    hasError: boolean;
}
/**
 * 包括的なエラー処理クラス
 */
class CommonErrorBoundary extends Component<ErrorBoundaryProps, ErrorBoundaryState> {
    private eventHandler;

    constructor(props: ErrorBoundaryProps) {
        super(props);
        this.state = { hasError: false };
        this.eventHandler = this.updateError.bind(this);
    }

    static getDerivedStateFromError(error: Error): ErrorBoundaryState {
        return { hasError: true };
    }

    updateError() {
        this.setState({ hasError: true });
    }

    /**
     * エラー処理
     */
    componentDidCatch(error: Error, info: ErrorInfo) {
        console.error("ErrorBoundary caught an error", error, info);
        localStorage.removeItem(KEYS.TOKEN);
        localStorage.removeItem(KEYS.USER_ID);
        localStorage.removeItem(KEYS.USER_ROLE);
    }
    /**
     * 非同期のエラー対応
     */
    componentDidMount(): void {
        // window.addEventListener('unhandledrejection', () => window.location.href = "/IndexnPage")
    }
    // componentWillUnmount(): void {
    //     window.removeEventListener('unhandledrejection', this.eventHandler)
    // }

    render() {
        if (this.state.hasError) {
            return (
                <CommonErrorView/>
            )
        }
        return this.props.children;
    }
}

export default CommonErrorBoundary;