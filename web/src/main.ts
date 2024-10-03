import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { bootstrapApplication, BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';
import { APP_ROUTES } from './app/app.route';
import { PreloadAllModules, provideRouter, withDebugTracing, withPreloading } from '@angular/router';
import { importProvidersFrom } from '@angular/core';


// platformBrowserDynamic().bootstrapModule(AppModule)
//   .catch(err => console.error(err));


bootstrapApplication(AppComponent, {
  providers: [
    importProvidersFrom(BrowserModule),
    provideRouter(
      APP_ROUTES,
      withPreloading(PreloadAllModules),
      withDebugTracing(),
    ),
  ],
  
}).catch(e => console.log(e));