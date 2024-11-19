interface IPageHeader {
    pageTitle: string;
}

interface IBreadcrumbPageHeader extends IPageHeader {
    breadcrumbItems: string[];
}

interface ISubPageHeader extends IPageHeader {
    chevronLink: string;
    children?: React.ReactNode;
}