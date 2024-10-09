import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SecretService } from 'src/app/services/secret.service';

@Component({
  standalone: true,
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  selector: 'asm-secret-list',
  templateUrl: './secret-list.component.html',
  styleUrl: './secret-list.component.css'
})
export class SecretListComponent implements OnInit{
  constructor(
    public secretService: SecretService,
  ) {}
  searchForm = new FormGroup({
    keyword: new FormControl(''),
    page: new FormControl(0),
    pageSize: new FormControl(1)
  })

  
  ngOnInit(): void {
  }

  submitSearch() {
    console.log('searchForm.value', this.searchForm.value);
    this.secretService.search$.next({
      keyword: this.searchForm.value.keyword || '',
      page: this.searchForm.value.page || 0,
      page_size: this.searchForm.value.pageSize || 2,
    });
  }


}
