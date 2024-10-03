import { Route } from "@angular/router";
import { SecretListComponent } from "./components/secret-list/secret-list.component";

export const APP_ROUTES: Route[] = [
    {path: '', pathMatch: 'full', redirectTo: 'list'},
    {path: 'list', component: SecretListComponent},
]